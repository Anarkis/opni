//go:build !noagentv1 && !nomanager

package commands

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-logr/zapr"
	upgraderesponder "github.com/longhorn/upgrade-responder/client"
	loggingv1beta1 "github.com/rancher/opni/apis/logging/v1beta1"
	"github.com/rancher/opni/controllers"
	agentv1 "github.com/rancher/opni/pkg/agent/v1"
	"github.com/rancher/opni/pkg/bootstrap"
	"github.com/rancher/opni/pkg/capabilities/wellknown"
	"github.com/rancher/opni/pkg/config"
	"github.com/rancher/opni/pkg/config/v1beta1"
	"github.com/rancher/opni/pkg/events"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/opni/common"
	"github.com/rancher/opni/pkg/pkp"
	"github.com/rancher/opni/pkg/test/testutil"
	"github.com/rancher/opni/pkg/tokens"
	"github.com/rancher/opni/pkg/tracing"
	"github.com/rancher/opni/pkg/trust"
	"github.com/rancher/opni/pkg/util"
	"github.com/rancher/opni/pkg/util/manager"
	"github.com/rancher/opni/pkg/util/waitctx"
	"github.com/rancher/opni/pkg/versions"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	ctrlzap "sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	configLocation       string
	agentLogLevel        string
	eventOutputEndpoint  string
	enableMetrics        bool
	enableLogging        bool
	enableEventCollector bool
)

func BuildAgentCmd() *cobra.Command {
	var disableUsage bool

	run := func(parentCtx context.Context) error {
		ctx := waitctx.FromContext(parentCtx)

		tracing.Configure("agent")
		agentlg := logger.New(logger.WithLogLevel(util.Must(zapcore.ParseLevel(agentLogLevel))))

		if os.Getenv("DO_NOT_TRACK") == "1" {
			disableUsage = true
		}

		var upgradeChecker *upgraderesponder.UpgradeChecker
		if !(disableUsage || common.DisableUsage) {
			upgradeRequester := manager.UpgradeRequester{
				Version:     versions.Version,
				InstallType: manager.InstallTypeAgent,
			}
			upgradeRequester.SetupLogger(zapr.NewLogger(agentlg.Desugar()))
			setupLog.Info("Usage tracking enabled", "current-version", versions.Version)
			upgradeChecker = upgraderesponder.NewUpgradeChecker(upgradeResponderAddress, &upgradeRequester)
			upgradeChecker.Start()
			defer upgradeChecker.Stop()
		}

		if enableMetrics {
			waitctx.Go(ctx, func() {
				runMonitoringAgent(ctx, agentlg)
			})
		}

		if enableLogging {
			waitctx.Go(ctx, func() {
				err := runLoggingControllers(ctx)
				if err != nil {
					agentlg.Fatalf("failed to start controllers: %v", err)
				}
			})
		}

		if enableEventCollector {
			waitctx.Go(ctx, func() {
				err := runEventsCollector(ctx)
				if err != nil {
					agentlg.Fatalf("failed to run event collector: %v", err)
				}
			})
		}

		waitctx.Wait(ctx)
		return nil
	}

	agentCmd := &cobra.Command{
		Use:   "agent",
		Short: "Run the Opni Monitoring Agent",
		Long: `The client component of the opni gateway, used to proxy the prometheus
agent remote-write requests to add dynamic authentication.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !enableMetrics && !enableLogging {
				return errors.New("at least one of [--metrics, --logging] must be set")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd.Context())
		},
	}

	agentCmd.Flags().StringVar(&configLocation, "config", "", "Absolute path to a config file")
	agentCmd.Flags().StringVar(&agentLogLevel, "log-level", "info", "log level (debug, info, warning, error)")
	agentCmd.Flags().BoolVar(&enableMetrics, "metrics", false, "enable metrics agent")
	agentCmd.Flags().BoolVar(&enableLogging, "logging", false, "enable logging controllers")
	agentCmd.Flags().BoolVar(&enableEventCollector, "events", false, "enable event collector")
	agentCmd.Flags().StringVar(&eventOutputEndpoint, "events-output", "http://opni-shipper:2021/log/ingest", "endpoint to post events to")
	return agentCmd
}

func configureBootstrap(conf *v1beta1.AgentConfig, agentlg logger.ExtendedSugaredLogger) (bootstrap.Bootstrapper, error) {
	var bootstrapper bootstrap.Bootstrapper
	var trustStrategy trust.Strategy
	if conf.Spec.Bootstrap == nil {
		return nil, errors.New("no bootstrap config provided")
	}
	if conf.Spec.Bootstrap.InClusterManagementAddress != nil {
		bootstrapper = &bootstrap.InClusterBootstrapper{
			Capability:         wellknown.CapabilityMetrics,
			GatewayEndpoint:    conf.Spec.GatewayAddress,
			ManagementEndpoint: *conf.Spec.Bootstrap.InClusterManagementAddress,
		}
	} else {
		agentlg.Info("loading bootstrap tokens from config file")
		tokenData := conf.Spec.Bootstrap.Token

		switch conf.Spec.TrustStrategy {
		case v1beta1.TrustStrategyPKP:
			var err error
			pins := conf.Spec.Bootstrap.Pins
			publicKeyPins := make([]*pkp.PublicKeyPin, len(pins))
			for i, pin := range pins {
				publicKeyPins[i], err = pkp.DecodePin(pin)
				if err != nil {
					agentlg.With(
						zap.Error(err),
						zap.String("pin", string(pin)),
					).Error("failed to parse pin")
					return nil, err
				}
			}
			conf := trust.StrategyConfig{
				PKP: &trust.PKPConfig{
					Pins: trust.NewPinSource(publicKeyPins),
				},
			}
			trustStrategy, err = conf.Build()
			if err != nil {
				agentlg.With(
					zap.Error(err),
				).Error("error configuring PKP trust strategy")
				return nil, err
			}
		case v1beta1.TrustStrategyCACerts:
			paths := conf.Spec.Bootstrap.CACerts
			certs := []*x509.Certificate{}
			for _, path := range paths {
				data, err := os.ReadFile(path)
				if err != nil {
					agentlg.With(
						zap.Error(err),
						zap.String("path", path),
					).Error("failed to read CA cert")
					return nil, err
				}
				cert, err := util.ParsePEMEncodedCert(data)
				if err != nil {
					agentlg.With(
						zap.Error(err),
						zap.String("path", path),
					).Error("failed to parse CA cert")
					return nil, err
				}
				certs = append(certs, cert)
			}
			conf := trust.StrategyConfig{
				CACerts: &trust.CACertsConfig{
					CACerts: trust.NewCACertsSource(certs),
				},
			}
			var err error
			trustStrategy, err = conf.Build()
			if err != nil {
				agentlg.With(
					zap.Error(err),
				).Error("error configuring CA Certs trust strategy")
				return nil, err
			}
		case v1beta1.TrustStrategyInsecure:
			agentlg.Warn(chalk.Bold.NewStyle().WithForeground(chalk.Yellow).Style(
				"*** Using insecure trust strategy. This is not recommended. ***",
			))
			conf := trust.StrategyConfig{
				Insecure: &trust.InsecureConfig{},
			}
			var err error
			trustStrategy, err = conf.Build()
			if err != nil {
				agentlg.With(
					zap.Error(err),
				).Error("error configuring insecure trust strategy")
				return nil, err
			}
		}

		token, err := tokens.ParseHex(tokenData)
		if err != nil {
			agentlg.With(
				zap.Error(err),
				zap.String("token", fmt.Sprintf("[redacted (len: %d)]", len(tokenData))),
			).Error("failed to parse token")
			return nil, err
		}
		bootstrapper = &bootstrap.ClientConfig{
			Capability:    wellknown.CapabilityMetrics,
			Token:         token,
			Endpoint:      conf.Spec.GatewayAddress,
			TrustStrategy: trustStrategy,
		}
	}

	return bootstrapper, nil
}

func runMonitoringAgent(ctx context.Context, agentlg logger.ExtendedSugaredLogger) {
	if configLocation == "" {
		// find config file
		path, err := config.FindConfig()
		if err != nil {
			if errors.Is(err, config.ErrConfigNotFound) {
				wd, _ := os.Getwd()
				agentlg.Fatalf(`could not find a config file in ["%s","/etc/opni"], and --config was not given`, wd)
			}
			agentlg.With(
				zap.Error(err),
			).Fatal("an error occurred while searching for a config file")
		}
		agentlg.With(
			"path", path,
		).Info("using config file")
		configLocation = path
	}

	objects, err := config.LoadObjectsFromFile(configLocation)
	if err != nil {
		agentlg.With(
			zap.Error(err),
		).Fatal("failed to load config")
	}
	var agentConfig *v1beta1.AgentConfig
	objects.Visit(func(config *v1beta1.AgentConfig) {
		agentConfig = config
	})

	bootstrapper, err := configureBootstrap(agentConfig, agentlg)
	if err != nil {
		agentlg.With(
			zap.Error(err),
		).Fatal("failed to configure bootstrap")
	}

	p, err := agentv1.New(ctx, agentConfig,
		agentv1.WithBootstrapper(bootstrapper),
	)
	if err != nil {
		agentlg.Error(err)
		return
	}
	if err := p.ListenAndServe(ctx); err != nil {
		agentlg.Error(err)
	}
}

func runLoggingControllers(ctx context.Context) error {
	ctrl.SetLogger(ctrlzap.New(
		ctrlzap.Level(util.Must(zapcore.ParseLevel(agentLogLevel))),
		ctrlzap.Encoder(zapcore.NewConsoleEncoder(testutil.EncoderConfig)),
	))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     "0",
		Port:                   9443,
		HealthProbeBindAddress: ":8081",
		LeaderElection:         false,
		LeaderElectionID:       "98e737d4.opni.io",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		return err
	}

	if err = (&controllers.LoggingReconciler{}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Logging")
		return err
	}

	if err = (&controllers.LoggingLogAdapterReconciler{}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Logging LogAdapter")
		return err
	}

	if err = (&controllers.LoggingDataPrepperReconciler{}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Logging DataPrepper")
		return err
	}

	if err = (&loggingv1beta1.LogAdapter{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Logging LogAdapter")
		return err
	}

	if !enableMetrics {
		if err := mgr.AddHealthzCheck("health", healthz.Ping); err != nil {
			setupLog.Error(err, "unable to set up health check")
			return err
		}
		if err := mgr.AddReadyzCheck("check", healthz.Ping); err != nil {
			setupLog.Error(err, "unable to set up ready check")
			return err
		}
	}

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	setupLog.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		return err
	}
	return nil
}

func runEventsCollector(ctx context.Context) error {
	collector := events.NewEventCollector(ctx, eventOutputEndpoint)
	return collector.Run(ctx.Done())
}

func init() {
	AddCommandsToGroup(OpniComponents, BuildAgentCmd())
}

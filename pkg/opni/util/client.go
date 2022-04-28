// package opni contains various utility and helper functions that are used
// by the Opnictl CLI.
package cliutil

import (
	"fmt"
	"os"

	"github.com/rancher/opni/apis"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClientOptions can be passed to some of the functions in this package when
// creating clients and/or client configurations.
type ClientOptions struct {
	overrides    *clientcmd.ConfigOverrides
	explicitPath string
}

type ClientOption func(*ClientOptions)

func (o *ClientOptions) Apply(opts ...ClientOption) {
	for _, op := range opts {
		op(o)
	}
}

// WithConfigOverrides allows overriding specific kubeconfig fields from the
// user's loaded kubeconfig.
func WithConfigOverrides(overrides *clientcmd.ConfigOverrides) ClientOption {
	return func(o *ClientOptions) {
		o.overrides = overrides
	}
}

func WithExplicitPath(path string) ClientOption {
	return func(o *ClientOptions) {
		o.explicitPath = path
	}
}

// CreateClientOrDie constructs a new controller-runtime client, or exit
// with a fatal error if an error occurs.
func CreateClientOrDie(opts ...ClientOption) (*api.Config, *rest.Config, client.Client) {
	scheme := CreateScheme()
	apiConfig, clientConfig := LoadClientConfig(opts...)

	cli, err := client.New(clientConfig, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	return apiConfig, clientConfig, cli
}

// LoadClientConfig loads the user's kubeconfig using the same logic as kubectl.
func LoadClientConfig(opts ...ClientOption) (*api.Config, *rest.Config) {
	options := ClientOptions{}
	options.Apply(opts...)

	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	// the loading rules check for empty string in the ExplicitPath, so it is
	// safe to always set this, it defaults to empty string.
	rules.ExplicitPath = options.explicitPath
	apiConfig, err := rules.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	clientConfig, err := clientcmd.NewDefaultClientConfig(
		*apiConfig, options.overrides).ClientConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	return apiConfig, clientConfig
}

// CreateScheme creates a new scheme with the types necessary for opnictl.
func CreateScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()
	apis.InitScheme(scheme)
	return scheme
}

module github.com/rancher/opni

go 1.16

require (
	github.com/go-logr/logr v0.3.0
	github.com/k3s-io/helm-controller v0.9.3
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/spf13/cobra v1.1.1
	github.com/ttacon/chalk v0.0.0-20160626202418-22c06c80ed31
	github.com/vbauerster/mpb/v7 v7.0.1
	go.uber.org/atomic v1.7.0
	go.uber.org/zap v1.17.0
	k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver v0.20.1
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009
	sigs.k8s.io/controller-runtime v0.8.3
)

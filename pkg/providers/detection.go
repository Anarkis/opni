// Package providers provides methods to detect different Kubernetes distros.
package providers

import (
	"context"
	"strings"

	"github.com/rancher/opni/apis/v1beta2"
	corev1 "k8s.io/api/core/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TODO: this probably already exists in a library somewhere

type Provider int

const (
	Unknown Provider = iota
	K3S
	RKE2
	RKE
)

func Detect(ctx context.Context, c client.Client) (Provider, error) {
	nodes := &corev1.NodeList{}
	if err := c.List(ctx, nodes); err != nil {
		return Unknown, err
	}
	for _, node := range nodes.Items {
		if strings.Contains(node.Spec.ProviderID, "k3s") {
			return K3S, nil
		} else if strings.Contains(node.Spec.ProviderID, "rke2") {
			return RKE2, nil
		} else if _, ok := node.ObjectMeta.Annotations["rke.cattle.io/internal-ip"]; ok {
			return RKE, nil
		}
	}
	return Unknown, nil
}

func (p Provider) ContainerRuntime() v1beta2.ContainerRuntime {
	switch p {
	case RKE:
		return v1beta2.Docker
	default:
		return v1beta2.Containerd
	}
}

func (p Provider) String() string {
	switch p {
	case K3S:
		return "k3s"
	case RKE2:
		return "rke2"
	case RKE:
		return "rke"
	default:
		return "unknown"
	}
}

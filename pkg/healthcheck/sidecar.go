package healthcheck

import (
	"strings"

	"github.com/linkerd/linkerd2/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
)

const (
	containerImageIstio       = "gcr.io/istio-release/proxyv2:"
	containerImageContour     = "gcr.io/heptio-images/contour:"
	containerImageEnvoy       = "docker.io/envoyproxy/envoy-alpine:"
	containerNameIstio        = "istio-proxy"
	containerNameContour      = "contour"
	containerNameEnvoy        = "envoy"
	initContainerImageIstio   = "gcr.io/istio-release/proxy_init:"
	initContainerImageContour = "gcr.io/heptio-images/contour:"
	initContainerNameIstio    = "istio-init"
	initContainerNameContour  = "envoy-initconfig"
)

// HasOtherProxies returns true if the pod spec already has either an istio,
// contour or envoy proxy container. Otherwise, it returns false.
func HasOtherProxies(podSpec *corev1.PodSpec) bool {
	for _, container := range podSpec.Containers {
		if strings.HasPrefix(container.Image, "gcr.io/linkerd-io/proxy:") ||
			strings.HasPrefix(container.Image, containerImageIstio) ||
			strings.HasPrefix(container.Image, containerImageContour) ||
			strings.HasPrefix(container.Image, containerImageEnvoy) ||
			container.Name == k8s.ProxyContainerName ||
			container.Name == containerNameIstio ||
			container.Name == containerNameContour ||
			container.Name == containerNameEnvoy {
			return true
		}
	}

	for _, ic := range podSpec.InitContainers {
		if strings.HasPrefix(ic.Image, "gcr.io/linkerd-io/proxy-init:") ||
			strings.HasPrefix(ic.Image, initContainerImageIstio) ||
			strings.HasPrefix(ic.Image, initContainerImageContour) ||
			ic.Name == "linkerd-init" ||
			ic.Name == initContainerNameIstio ||
			ic.Name == initContainerNameContour {
			return true
		}
	}

	return false
}

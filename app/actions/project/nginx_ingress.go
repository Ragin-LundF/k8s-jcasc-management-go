package project

import (
	"k8s-management-go/app/models"
)

// ----- Structures
// nginx : Model which describes the NginxIngressController and load balancer
type nginx struct {
	Ingress      ingress
	LoadBalancer loadBalancer
}

// ingress : Model which describes the Nginx Ingress controller
type ingress struct {
	AnnotationIngressClass       string
	DeploymentName               string
	ContainerImage               string
	ImagePullSecrets             string
	EnableControllerForNamespace bool
}

// loadBalancer : Model which describes the Nginx load balancer
type loadBalancer struct {
	Enabled     bool
	Annotations loadBalancerAnnotations
	Ports       loadBalancerPorts
}

// loadBalancerPorts : Model which describes the ports of loadBalancer
type loadBalancerPorts struct {
	HTTP  loadBalancerPortsHTTP
	HTTPS loadBalancerPortsHTTP
}

// loadBalancerPortsHTTP : Model which describes the concrete ports for loadBalancerPorts
type loadBalancerPortsHTTP struct {
	Port       uint64
	TargetPort uint64
}

// loadBalancerAnnotations : Model which describes the loadBalancer annotations
type loadBalancerAnnotations struct {
	Enabled        bool
	ExternalDnsTtl uint64
	namespace      string
}

// NewNginx : creates a new instance of Nginx
func NewNginx() *nginx {
	return &nginx{
		Ingress:      newDefaultIngress(),
		LoadBalancer: newDefaultLoadBalancer(),
	}
}

// ----- internal methods
// newDefaultMetadata : create a new default ingress structure
func newDefaultIngress() ingress {
	var configuration = models.GetConfiguration()
	return ingress{
		AnnotationIngressClass:       configuration.Nginx.Ingress.AnnotationClass,
		DeploymentName:               configuration.Nginx.Ingress.Controller.DeploymentName,
		ContainerImage:               configuration.Nginx.Ingress.Controller.Container.Name,
		ImagePullSecrets:             configuration.Nginx.Ingress.Controller.Container.PullSecret,
		EnableControllerForNamespace: configuration.Nginx.Ingress.Controller.Container.Namespace,
	}
}

// newDefaultLoadBalancer : create a new default loadBalancer structure
func newDefaultLoadBalancer() loadBalancer {
	var configuration = models.GetConfiguration()
	return loadBalancer{
		Enabled:     configuration.LoadBalancer.Enabled,
		Annotations: newDefaultLoadBalancerAnnotations(),
		Ports:       newDefaultLoadBalancerPorts(),
	}
}

// newDefaultLoadBalancerAnnotations : create a new default loadBalancerAnnotations structure
func newDefaultLoadBalancerAnnotations() loadBalancerAnnotations {
	var configuration = models.GetConfiguration()

	return loadBalancerAnnotations{
		Enabled:        configuration.LoadBalancer.Annotations.Enabled,
		ExternalDnsTtl: configuration.LoadBalancer.Annotations.ExtDNS.Ttl,
	}
}

// newDefaultLoadBalancerPorts : create a new default loadBalancerPorts structure
func newDefaultLoadBalancerPorts() loadBalancerPorts {
	return loadBalancerPorts{
		HTTP:  newDefaultLoadBalancerPortsHTTP(),
		HTTPS: newDefaultLoadBalancerPortsHTTPS(),
	}
}

// newDefaultLoadBalancerPortsHTTP : create a new default loadBalancerPortsHTTP structure for HTTP
func newDefaultLoadBalancerPortsHTTP() loadBalancerPortsHTTP {
	var configuration = models.GetConfiguration()
	return loadBalancerPortsHTTP{
		Port:       configuration.LoadBalancer.Port.HTTP,
		TargetPort: configuration.LoadBalancer.Port.HTTPTarget,
	}
}

// newDefaultLoadBalancerPortsHTTPS : create a new default loadBalancerPortsHTTP structure for HTTPS
func newDefaultLoadBalancerPortsHTTPS() loadBalancerPortsHTTP {
	var configuration = models.GetConfiguration()
	return loadBalancerPortsHTTP{
		Port:       configuration.LoadBalancer.Port.HTTPS,
		TargetPort: configuration.LoadBalancer.Port.HTTPSTarget,
	}
}

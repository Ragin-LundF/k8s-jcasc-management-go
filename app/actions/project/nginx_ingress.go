package project

import (
	"k8s-management-go/app/configuration"
)

// ----- Structures
// nginx : Model which describes the NginxIngressController and load balancer
type nginx struct {
	Ingress      ingress      `yaml:"ingress,omitempty"`
	LoadBalancer loadBalancer `yaml:"loadBalancer,omitempty"`
}

// ingress : Model which describes the Nginx Ingress controller
type ingress struct {
	AnnotationIngressClass       string `yaml:"annotationIngressClass,omitempty"`
	DeploymentName               string `yaml:"deploymentName,omitempty"`
	ContainerImage               string `yaml:"containerImage,omitempty"`
	ImagePullSecrets             string `yaml:"imagePullSecrets,omitempty"`
	EnableControllerForNamespace bool   `yaml:"enableControllerForNamespace,omitempty"`
}

// loadBalancer : Model which describes the Nginx load balancer
type loadBalancer struct {
	Enabled     bool
	Annotations loadBalancerAnnotations `yaml:"annotations,omitempty"`
	Ports       loadBalancerPorts       `yaml:"ports,omitempty"`
}

// loadBalancerPorts : Model which describes the ports of loadBalancer
type loadBalancerPorts struct {
	HTTP  loadBalancerPortsHTTP `yaml:"http,omitempty"`
	HTTPS loadBalancerPortsHTTP `yaml:"https,omitempty"`
}

// loadBalancerPortsHTTP : Model which describes the concrete ports for loadBalancerPorts
type loadBalancerPortsHTTP struct {
	Port       uint64 `yaml:"port,omitempty"`
	TargetPort uint64 `yaml:"targetPort,omitempty"`
}

// loadBalancerAnnotations : Model which describes the loadBalancer annotations
type loadBalancerAnnotations struct {
	Enabled             bool   `yaml:"enabled,omitempty"`
	ExternalDnsHostname string `yaml:"externalDNSHostName,omitempty"`
	ExternalDnsTtl      uint64 `yaml:"externalDNSTtl,omitempty"`
}

// newNginx : creates a new instance of Nginx
func newNginx() *nginx {
	return &nginx{
		Ingress:      newDefaultIngress(),
		LoadBalancer: newDefaultLoadBalancer(),
	}
}

// ----- internal methods
// newDefaultMetadata : create a new default ingress structure
func newDefaultIngress() ingress {
	return ingress{
		AnnotationIngressClass:       configuration.GetConfiguration().Nginx.Ingress.Annotationclass,
		DeploymentName:               configuration.GetConfiguration().Nginx.Ingress.Deployment.DeploymentName,
		ContainerImage:               configuration.GetConfiguration().Nginx.Ingress.Container.Image,
		ImagePullSecrets:             configuration.GetConfiguration().Nginx.Ingress.Container.PullSecret,
		EnableControllerForNamespace: configuration.GetConfiguration().Nginx.Ingress.Deployment.ForEachNamespace,
	}
}

// newDefaultLoadBalancer : create a new default loadBalancer structure
func newDefaultLoadBalancer() loadBalancer {
	return loadBalancer{
		Enabled:     configuration.GetConfiguration().Nginx.Loadbalancer.Enabled,
		Annotations: newDefaultLoadBalancerAnnotations(),
		Ports:       newDefaultLoadBalancerPorts(),
	}
}

// newDefaultLoadBalancerAnnotations : create a new default loadBalancerAnnotations structure
func newDefaultLoadBalancerAnnotations() loadBalancerAnnotations {
	return loadBalancerAnnotations{
		Enabled:             configuration.GetConfiguration().Nginx.Loadbalancer.Annotations.Enabled,
		ExternalDnsHostname: configuration.GetConfiguration().Nginx.Loadbalancer.ExternalDNS.HostName,
		ExternalDnsTtl:      configuration.GetConfiguration().Nginx.Loadbalancer.ExternalDNS.TTL,
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
	return loadBalancerPortsHTTP{
		Port:       configuration.GetConfiguration().Nginx.Loadbalancer.Ports.HTTP,
		TargetPort: configuration.GetConfiguration().Nginx.Loadbalancer.Ports.HTTPTarget,
	}
}

// newDefaultLoadBalancerPortsHTTPS : create a new default loadBalancerPortsHTTP structure for HTTPS
func newDefaultLoadBalancerPortsHTTPS() loadBalancerPortsHTTP {
	return loadBalancerPortsHTTP{
		Port:       configuration.GetConfiguration().Nginx.Loadbalancer.Ports.HTTPS,
		TargetPort: configuration.GetConfiguration().Nginx.Loadbalancer.Ports.HTTPSTarget,
	}
}

package project

import (
	"fmt"
	"k8s-management-go/app/models"
)

// ----- Structures
// nginx : Model which describes the nginx ingress controller and load balancer
type nginx struct {
	Ingress      ingress
	LoadBalancer loadBalancer
}

// ingress : Model which describes the Nginx Ingress controller
type ingress struct {
	AnnotationIngressClass       string
	LoadBalancerIP               string
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
	Enabled             bool
	ExternalDnsHostname string
	ExternalDnsTtl      uint64
}

// NewNginx : creates a new instance of Nginx
func NewNginx(namespace string, loadBalancerIP *string, annotationExtDnsName *string) *nginx {
	var assignedLoadBalancerIP string

	// set default values
	if loadBalancerIP != nil {
		assignedLoadBalancerIP = *loadBalancerIP
	} else {
		assignedLoadBalancerIP = ""
	}

	return &nginx{
		Ingress:      newDefaultIngress(assignedLoadBalancerIP),
		LoadBalancer: newDefaultLoadBalancer(namespace, annotationExtDnsName),
	}
}

// ----- Setter to manipulate the default object
// SetIngressLoadBalancerIPAddress : Set load balancer IP address to ingress controller
func (nginx *nginx) SetIngressLoadBalancerIPAddress(ipAddress string) {
	nginx.Ingress.LoadBalancerIP = ipAddress
}

// ----- internal methods
// newDefaultMetadata : create a new default ingress structure
func newDefaultIngress(loadBalancerIP string) ingress {
	var configuration = models.GetConfiguration()
	return ingress{
		LoadBalancerIP:               loadBalancerIP,
		AnnotationIngressClass:       configuration.Nginx.Ingress.AnnotationClass,
		DeploymentName:               configuration.Nginx.Ingress.Controller.DeploymentName,
		ContainerImage:               configuration.Nginx.Ingress.Controller.Container.Name,
		ImagePullSecrets:             configuration.Nginx.Ingress.Controller.Container.PullSecret,
		EnableControllerForNamespace: configuration.Nginx.Ingress.Controller.Container.Namespace,
	}
}

// newDefaultLoadBalancer : create a new default loadBalancer structure
func newDefaultLoadBalancer(namespace string, annotationExtDnsName *string) loadBalancer {
	var configuration = models.GetConfiguration()
	return loadBalancer{
		Enabled:     configuration.LoadBalancer.Enabled,
		Annotations: newDefaultLoadBalancerAnnotations(namespace, annotationExtDnsName),
		Ports:       newDefaultLoadBalancerPorts(),
	}
}

// newDefaultLoadBalancerAnnotations : create a new default loadBalancerAnnotations structure
func newDefaultLoadBalancerAnnotations(namespace string, annotationExtDnsName *string) loadBalancerAnnotations {
	var configuration = models.GetConfiguration()

	var externalDnsHost string
	if annotationExtDnsName == nil {
		externalDnsHost = fmt.Sprintf("%v.%v", namespace, configuration.LoadBalancer.Annotations.ExtDNS.Hostname)
	} else {
		externalDnsHost = *annotationExtDnsName
	}

	return loadBalancerAnnotations{
		Enabled:             configuration.LoadBalancer.Annotations.Enabled,
		ExternalDnsHostname: externalDnsHost,
		ExternalDnsTtl:      configuration.LoadBalancer.Annotations.ExtDNS.Ttl,
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

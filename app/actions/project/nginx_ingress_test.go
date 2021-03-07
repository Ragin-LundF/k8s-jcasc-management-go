package project

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNginx(t *testing.T) {
	testDefaultProjectConfiguration(t, false)

	var nginx = NewNginx()
	assertDefaultNginxConfiguration(nginx, t)
}

func assertDefaultNginxConfiguration(nginx *nginx, t *testing.T) {
	assert.NotNil(t, nginx.Ingress)
	assert.Equal(t, testNginxIngressAnnotationClass, nginx.Ingress.AnnotationIngressClass)

	assert.Equal(t, testNginxIngressDeploymentName, nginx.Ingress.DeploymentName)
	assert.Equal(t, testNginxIngressControllerContainerImage, nginx.Ingress.ContainerImage)
	assert.Equal(t, testNginxIngressControllerContainerPullSecrets, nginx.Ingress.ImagePullSecrets)
	assert.Equal(t, testNginxIngressControllerForNamespace, nginx.Ingress.EnableControllerForNamespace)

	assert.NotNil(t, nginx.LoadBalancer)
	assert.Equal(t, testNginxLoadBalancerEnabled, nginx.LoadBalancer.Enabled)
	assert.Equal(t, testNginxLoadBalancerHttpPort, nginx.LoadBalancer.Ports.HTTP.Port)
	assert.Equal(t, testNginxLoadBalancerHttpTargetPort, nginx.LoadBalancer.Ports.HTTP.TargetPort)
	assert.Equal(t, testNginxLoadBalancerHttpsPort, nginx.LoadBalancer.Ports.HTTPS.Port)
	assert.Equal(t, testNginxLoadBalancerHttpsTargetPort, nginx.LoadBalancer.Ports.HTTPS.TargetPort)
	assert.Equal(t, testNginxLoadBalancerAnnotationsEnabled, nginx.LoadBalancer.Annotations.Enabled)
	assert.Equal(t, testNginxLoadBalancerAnnotationsExtDnsHostname, nginx.LoadBalancer.Annotations.ExternalDnsHostname)
	assert.Equal(t, testNginxLoadBalancerAnnotationsExtDnsTtl, nginx.LoadBalancer.Annotations.ExternalDnsTtl)
}

package project

import (
	"github.com/stretchr/testify/assert"
	"strconv"
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
	expectedEnableControllerForNamespace, _ := strconv.ParseBool(testNginxIngressControllerForNamespace)
	assert.Equal(t, expectedEnableControllerForNamespace, nginx.Ingress.EnableControllerForNamespace)

	assert.NotNil(t, nginx.LoadBalancer)
	expectedEnableLoadBalancer, _ := strconv.ParseBool(testNginxLoadBalancerEnabled)
	assert.Equal(t, expectedEnableLoadBalancer, nginx.LoadBalancer.Enabled)
	expectedLoadBalancerHTTPPort, _ := strconv.ParseUint(testNginxLoadBalancerHttpPort, 10, 16)
	assert.Equal(t, expectedLoadBalancerHTTPPort, nginx.LoadBalancer.Ports.HTTP.Port)
	expectedLoadBalancerHTTPTargetPort, _ := strconv.ParseUint(testNginxLoadBalancerHttpTargetPort, 10, 16)
	assert.Equal(t, expectedLoadBalancerHTTPTargetPort, nginx.LoadBalancer.Ports.HTTP.TargetPort)
	expectedLoadBalancerHTTPSPort, _ := strconv.ParseUint(testNginxLoadBalancerHttpsPort, 10, 16)
	assert.Equal(t, expectedLoadBalancerHTTPSPort, nginx.LoadBalancer.Ports.HTTPS.Port)
	expectedLoadBalancerHTTPSTargetPort, _ := strconv.ParseUint(testNginxLoadBalancerHttpsTargetPort, 10, 16)
	assert.Equal(t, expectedLoadBalancerHTTPSTargetPort, nginx.LoadBalancer.Ports.HTTPS.TargetPort)
	expectedEnableLoadBalancerAnnotation, _ := strconv.ParseBool(testNginxLoadBalancerAnnotationsEnabled)
	assert.Equal(t, expectedEnableLoadBalancerAnnotation, nginx.LoadBalancer.Annotations.Enabled)
	assert.Equal(t, testNginxLoadBalancerAnnotationsExtDnsHostname, nginx.LoadBalancer.Annotations.ExternalDnsHostname)
	expectedNginxLoadBalancerAnnotationsExtDnsTtl, _ := strconv.ParseUint(testNginxLoadBalancerAnnotationsExtDnsTtl, 10, 16)
	assert.Equal(t, expectedNginxLoadBalancerAnnotationsExtDnsTtl, nginx.LoadBalancer.Annotations.ExternalDnsTtl)
}

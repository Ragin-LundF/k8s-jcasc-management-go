package validator

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/configuration"
	"testing"
)

func init() {
	setup()
}

func setup() {
	configuration.LoadConfiguration("../../../", false, false)
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = nil

	var newIpAndNamespace = configuration.DeploymentStruct{
		IPAddress: "1.2.3.4",
		Namespace: "existing-namespace",
	}
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = append(configuration.GetConfiguration().K8SManagement.IPConfig.Deployments, newIpAndNamespace)
	newIpAndNamespace = configuration.DeploymentStruct{
		IPAddress: "1.2.3.5",
		Namespace: "valid-namespace",
	}
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = append(configuration.GetConfiguration().K8SManagement.IPConfig.Deployments, newIpAndNamespace)
	newIpAndNamespace = configuration.DeploymentStruct{
		IPAddress: "1.2.3.6",
		Namespace: "product-dev",
	}
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = append(configuration.GetConfiguration().K8SManagement.IPConfig.Deployments, newIpAndNamespace)
}

func TestValidateNamespaceAvailableInConfig(t *testing.T) {
	setup()
	exists := ValidateNamespaceAvailableInConfig("existing-namespace")

	assert.True(t, exists)
}

func TestValidateNamespaceAvailableInConfigOtherNamespace(t *testing.T) {
	setup()
	exists := ValidateNamespaceAvailableInConfig("valid-namespace")

	assert.True(t, exists)
}

func TestValidateNamespaceAvailableInConfigWithWrongNamespace(t *testing.T) {
	setup()
	exists := ValidateNamespaceAvailableInConfig("valid-namespace2")

	assert.False(t, exists)
}

func TestValidateNamespaceAvailableInConfigWithWrongNamespaceLessCharacters(t *testing.T) {
	setup()
	exists := ValidateNamespaceAvailableInConfig("valid-namespac")

	assert.False(t, exists)
}

func TestValidateNewNamespace(t *testing.T) {
	setup()
	var namespace = "correct-namespace"
	err := ValidateNewNamespace(namespace)

	assert.NoError(t, err)
}

func TestValidateNewNamespaceInvalid(t *testing.T) {
	setup()
	var namespace = "correct-namespace-"
	err := ValidateNewNamespace(namespace)

	assert.Error(t, err)
}

func TestValidateNewNamespaceTooLong(t *testing.T) {
	setup()
	var namespace = "incorrect-namespace-with-more-than-allowed-63-characters-should-fail"
	err := ValidateNewNamespace(namespace)

	assert.Error(t, err)
}

func TestValidateNewNamespaceAlreadyExisting(t *testing.T) {
	setup()
	var namespace = "product-dev"
	err := ValidateNewNamespace(namespace)

	assert.Error(t, err)
}

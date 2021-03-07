package validator

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/actions/migration"
	"testing"
)

func init() {
	migration.ResetIPAndNamespaces()
	migration.AddIPAndNamespaceToConfiguration("existing-namespace", "1.2.3.4")
	migration.AddIPAndNamespaceToConfiguration("valid-namespace", "1.2.3.5")
	migration.AddIPAndNamespaceToConfiguration("product-dev", "1.2.3.6")
}

func TestValidateNamespaceAvailableInConfig(t *testing.T) {
	exists := ValidateNamespaceAvailableInConfig("existing-namespace")

	assert.True(t, exists)
}

func TestValidateNamespaceAvailableInConfigOtherNamespace(t *testing.T) {
	exists := ValidateNamespaceAvailableInConfig("valid-namespace")

	assert.True(t, exists)
}

func TestValidateNamespaceAvailableInConfigWithWrongNamespace(t *testing.T) {
	exists := ValidateNamespaceAvailableInConfig("valid-namespace2")

	assert.False(t, exists)
}

func TestValidateNamespaceAvailableInConfigWithWrongNamespaceLessCharacters(t *testing.T) {
	exists := ValidateNamespaceAvailableInConfig("valid-namespac")

	assert.False(t, exists)
}

func TestValidateNewNamespace(t *testing.T) {
	var namespace = "correct-namespace"
	err := ValidateNewNamespace(namespace)

	assert.NoError(t, err)
}

func TestValidateNewNamespaceInvalid(t *testing.T) {
	var namespace = "correct-namespace-"
	err := ValidateNewNamespace(namespace)

	assert.Error(t, err)
}

func TestValidateNewNamespaceTooLong(t *testing.T) {
	var namespace = "incorrect-namespace-with-more-than-allowed-63-characters-should-fail"
	err := ValidateNewNamespace(namespace)

	assert.Error(t, err)
}

func TestValidateNewNamespaceAlreadyExisting(t *testing.T) {
	var namespace = "product-dev"
	err := ValidateNewNamespace(namespace)

	assert.Error(t, err)
}

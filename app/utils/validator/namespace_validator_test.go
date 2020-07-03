package validator

import (
	"k8s-management-go/app/models"
	"testing"
)

func init() {
	models.ResetIPAndNamespaces()
	models.AddIPAndNamespaceToConfiguration("existing-namespace", "1.2.3.4")
	models.AddIPAndNamespaceToConfiguration("valid-namespace", "1.2.3.5")
	models.AddIPAndNamespaceToConfiguration("product-dev", "1.2.3.6")
}

func TestValidateNamespaceAvailableInConfig(t *testing.T) {
	exists := ValidateNamespaceAvailableInConfig("existing-namespace")

	if exists {
		t.Log("Success. Found existing-namespace")
	} else {
		t.Error("Failed. Validator did not recognize the existing-namespace")
	}
}

func TestValidateNamespaceAvailableInConfigOtherNamespace(t *testing.T) {
	exists := ValidateNamespaceAvailableInConfig("valid-namespace")

	if exists {
		t.Log("Success. Found existing valid-namespace")
	} else {
		t.Error("Failed. Validator did not recognize the existing valid-namespace")
	}
}

func TestValidateNamespaceAvailableInConfigWithWrongNamespace(t *testing.T) {
	exists := ValidateNamespaceAvailableInConfig("valid-namespace2")

	if exists {
		t.Error("Failed. Validator did found non existing valid-namespace2")
	} else {
		t.Log("Success. Validator recognized that valid-namespace2 is not existing")
	}
}

func TestValidateNamespaceAvailableInConfigWithWrongNamespaceLessCharacters(t *testing.T) {
	exists := ValidateNamespaceAvailableInConfig("valid-namespac")

	if exists {
		t.Error("Failed. Validator did found non existing valid-namespac")
	} else {
		t.Log("Success. Validator recognized that valid-namespac is not existing")
	}
}

func TestValidateNewNamespace(t *testing.T) {
	var namespace = "correct-namespace"
	err := ValidateNewNamespace(namespace)

	if err != nil {
		t.Error("Failed. Valid namespace was recognized as invalid.")
	} else {
		t.Log("Success. Valid namespace accepted.")
	}
}

func TestValidateNewNamespaceInvalid(t *testing.T) {
	var namespace = "correct-namespace-"
	err := ValidateNewNamespace(namespace)

	if err != nil {
		t.Log("Success. Invalid namespace rejected.")
	} else {
		t.Error("Failed. Invalid namespace was accepted.")
	}
}

func TestValidateNewNamespaceTooLong(t *testing.T) {
	var namespace = "incorrect-namespace-with-more-than-allowed-63-characters-should-fail"
	err := ValidateNewNamespace(namespace)

	if err != nil {
		t.Log("Success. Too long namespace rejected.")
	} else {
		t.Error("Failed. Too long namespace was accepted.")
	}
}

package validator

import (
	"errors"
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/configuration"
	"regexp"
	"strings"
)

// ValidateNamespaceAvailableInConfig checks selected namespace against namespace list
func ValidateNamespaceAvailableInConfig(namespaceToValidate string) bool {
	for _, ip := range configuration.GetConfiguration().K8SManagement.IPConfig.Deployments {
		if ip.Namespace == namespaceToValidate {
			return true
		}
	}
	return false
}

// ValidateNewNamespace validates the namespace
func ValidateNewNamespace(input string) error {
	// a namespace name cannot be longer than 63 characters
	if len(input) > 63 {
		return errors.New("Namespace name is too long! You can only use max. 63 characters. ")
	}
	// Regex to have DNS compatible string
	var regex = regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	if !regex.Match([]byte(input)) {
		return errors.New("Namespace is not valid! It must fit to DNS specification! ")
	}
	// check, that namespace was not already used
	for _, ipConfig := range configuration.GetConfiguration().K8SManagement.IPConfig.Deployments {
		if strings.ToLower(ipConfig.Namespace) == strings.ToLower(input) {
			return errors.New("Namespace already in use! ")
		}
	}
	return nil
}

// ValidateAdditionalNamespaces validates the namespace
func ValidateAdditionalNamespaces(input string) error {
	if len(input) > 0 {
		var err error
		var additionalNamespaces = project.ProcessAdditionalNamespaces(input)
		for _, additionalNamespace := range additionalNamespaces {
			err = ValidateNewNamespace(additionalNamespace)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

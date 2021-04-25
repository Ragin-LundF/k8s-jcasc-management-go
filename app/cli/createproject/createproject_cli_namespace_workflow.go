package createproject

import (
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

// NamespaceWorkflow represents the namespace workflow
func NamespaceWorkflow() (namespace string, err error) {
	// Validator for namespace name
	var validateNamespace = validator.ValidateNewNamespace

	// get namespace
	// Prepare prompt
	dialogs.ClearScreen()
	namespace, err = dialogs.DialogPrompt(constants.TextEnterNamespaceName, validateNamespace)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(constants.LogUnableToGetNameOfNewNamespace, err.Error())
		return namespace, err
	}

	return namespace, nil
}

// AdditionalNamespaceWorkflow represents the additional namespace workflow
func AdditionalNamespaceWorkflow() (additionalNamespacesArr []string, err error) {
	// Validator for additional namespace name
	var validateAddNamespace = validator.ValidateAdditionalNamespaces

	// get additional namespaces
	// Prepare prompt
	dialogs.ClearScreen()
	additionalNamespaces, err := dialogs.DialogPrompt(constants.TextEnterAdditionalNamespaceName, validateAddNamespace)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(constants.LogUnableToGetNameOfNewNamespace, err.Error())
		return []string{}, err
	}

	return project.ProcessAdditionalNamespaces(additionalNamespaces), nil
}

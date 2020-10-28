package createproject

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

// NamespaceWorkflow represents the namespace workflow
func NamespaceWorkflow() (namespace string, err error) {
	// Validator for namespace name
	validate := validator.ValidateNewNamespace

	// Prepare prompt
	dialogs.ClearScreen()
	namespace, err = dialogs.DialogPrompt(constants.TextEnterNamespaceName, validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(constants.LogUnableToGetNameOfNewNamespace, err.Error())
		return namespace, err
	}

	return namespace, nil
}

package createproject

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

// JenkinsDomainWorkflow represents the domain for Jenkins workflow
func JenkinsDomainWorkflow() (jenkinsUrl string, err error) {
	// Validator for IP address
	var validate = validator.ValidateIP

	// Prepare prompt
	dialogs.ClearScreen()
	jenkinsUrl, err = dialogs.DialogPrompt(constants.TextEnterJenkinsUrl, validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(constants.LogErrUnableToGetIPAddress, err.Error())
		return jenkinsUrl, err
	}

	return jenkinsUrl, nil
}

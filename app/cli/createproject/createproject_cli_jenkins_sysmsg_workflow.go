package createproject

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

// JenkinsSystemMessageWorkflow represents the Jenkins system message workflow
func JenkinsSystemMessageWorkflow() (jenkinsSysMsg string, err error) {
	// Validator for jenkins system message
	validate := validator.ValidateJenkinsSystemMessage

	// Prepare prompt
	dialogs.ClearScreen()
	jenkinsSysMsg, err = dialogs.DialogPrompt("Enter the Jenkins system message or leave empty for default", validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get the Jenkins system message.", err.Error())
		return jenkinsSysMsg, err
	}

	// check if system message is empty, set default
	if jenkinsSysMsg == "" {
		jenkinsSysMsg = constants.CommonJenkinsSystemMessage
	}

	return jenkinsSysMsg, nil
}

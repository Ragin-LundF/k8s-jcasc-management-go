package createproject

import (
	"fmt"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

func JenkinsSystemMessageWorkflow(namespace string) (jenkinsSysMsg string, err error) {
	// Validator for jenkins system message
	validate := validator.JenkinsSystemMessageValidator

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
		jenkinsSysMsg = fmt.Sprintf("Jenkins instance for namespace [%s]", namespace)
	}

	return jenkinsSysMsg, nil
}

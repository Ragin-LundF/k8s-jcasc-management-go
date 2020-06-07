package createproject

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/utils/logger"
)

func ProjectWizardAskForJenkinsSystemMessage(namespace string) (jenkinsSysMsg string, err error) {
	log := logger.Log()
	// Validator for jenkins system message
	validate := func(input string) error {
		// a namespace name cannot be longer than 63 characters
		if len(input) > 255 {
			return errors.New("Should not be longer than 255 characters. ")
		}
		return nil
	}

	// Prepare prompt
	dialogs.ClearScreen()
	jenkinsSysMsg, err = dialogs.DialogPrompt("Enter the Jenkins system message or leave empty for default", validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get the Jenkins system message.", err.Error())
		log.Errorf("[ProjectWizardAskForJenkinsSystemMessage] Unable to get the Jenkins system message. %s\n", err.Error())
		return jenkinsSysMsg, err
	}

	// check if system message is empty, set default
	if jenkinsSysMsg == "" {
		jenkinsSysMsg = "Jenkins instance for namespace [" + namespace + "]"
	}

	return jenkinsSysMsg, nil
}

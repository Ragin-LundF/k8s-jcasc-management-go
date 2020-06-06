package createproject

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
)

func ProjectWizardAskForJenkinsSystemMessage(namespace string) (jenkinsSysMsg string, err error) {
	log := logger.Log()
	// Validator for jenkins system message
	validate := func(input string) error {
		// a namespace name can not be longer than 63 characters
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
		log.Error("[ProjectWizardAskForJenkinsSystemMessage] Unable to get the Jenkins system message. %v\n", err)
	}

	// check if system message is empty, set default
	if jenkinsSysMsg == "" {
		jenkinsSysMsg = "Jenkins instance for namespace [" + namespace + "]"
	}

	return jenkinsSysMsg, err
}

// Replace jenkins system message
func ProcessJenkinsSystemMessage(projectDirectory string, jenkinsSysMsg string) (success bool, err error) {
	log := logger.Log()
	jenkinsHelmValuesFile := files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode)
	if files.FileOrDirectoryExists(jenkinsHelmValuesFile) {
		successful, err := files.ReplaceStringInFile(jenkinsHelmValuesFile, constants.TemplateJenkinsSystemMessage, jenkinsSysMsg)
		if !successful || err != nil {
			log.Error("[ProcessJenkinsSystemMessage] Can not replace Jenkins System msg in file [%v], \n%v", jenkinsHelmValuesFile, err)
			return false, err
		}
	}
	return true, err
}

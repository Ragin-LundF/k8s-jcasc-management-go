package createproject

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

// JenkinsJobsConfigRepositoryWorkflow represents the config repository workflow
func JenkinsJobsConfigRepositoryWorkflow() (jenkinsJobsCfgRepo string, err error) {
	// Validator
	var validate = validator.ValidateJenkinsJobConfig

	// Prepare prompt
	dialogs.ClearScreen()
	jenkinsJobsCfgRepo, err = dialogs.DialogPrompt(constants.TextEnterJobsConfigurationRepository, validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(constants.LogUnableToGetJobsConfigurationRepository, err.Error())
		return jenkinsJobsCfgRepo, err
	}

	return jenkinsJobsCfgRepo, nil
}

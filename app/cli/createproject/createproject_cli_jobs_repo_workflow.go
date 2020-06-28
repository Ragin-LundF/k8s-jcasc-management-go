package createproject

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

func JenkinsJobsConfigRepositoryWorkflow() (jenkinsJobsCfgRepo string, err error) {
	// Validator
	validate := validator.ValidateJenkinsJobConfig

	// Prepare prompt
	dialogs.ClearScreen()
	jenkinsJobsCfgRepo, err = dialogs.DialogPrompt("Enter jobs configuration repository", validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get the jobs configuration repository.", err.Error())
		return jenkinsJobsCfgRepo, err
	}

	return jenkinsJobsCfgRepo, nil
}

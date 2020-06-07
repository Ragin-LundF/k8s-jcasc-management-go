package createproject

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"regexp"
)

func ProjectWizardAskForJobsConfigurationRepository() (jenkinsSysMsg string, err error) {
	log := logger.Log()
	// Validator
	validate := func(input string) error {
		// Job repository should not be longer than 512 characters
		if len(input) > 512 {
			return errors.New("Should not be longer than 512 characters. ")
		}
		// Regex regex to validate repository
		regex := regexp.MustCompile(models.GetConfiguration().Jenkins.JobDSL.RepoValidatePattern)
		if !regex.Match([]byte(input)) {
			return errors.New("Wrong repository name! ")
		}

		return nil
	}

	// Prepare prompt
	dialogs.ClearScreen()
	jenkinsSysMsg, err = dialogs.DialogPrompt("Enter jobs configuration repository", validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get the jobs configuration repository.", err.Error())
		log.Error("[ProjectWizardAskForJenkinsSystemMessage] Unable to get the jobs configuration repository. %v\n", err)
	}

	return jenkinsSysMsg, err
}

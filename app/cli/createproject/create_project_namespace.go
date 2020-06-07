package createproject

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"regexp"
	"strings"
)

func ProjectWizardAskForNamespace() (namespace string, err error) {
	log := logger.Log()
	// Validator for namespace name
	validate := func(input string) error {
		// a namespace name cannot be longer than 63 characters
		if len(input) > 63 {
			return errors.New("Namespace name is too long! You can only use max. 63 characters. ")
		}
		// Regex to have DNS compatible string
		regex := regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
		if !regex.Match([]byte(input)) {
			return errors.New("Namespace is not valid! It must fit to DNS specification! ")
		}
		// check, that namespace was not already used
		for _, ipConfig := range models.GetIpConfiguration().Ips {
			if strings.ToLower(ipConfig.Namespace) == strings.ToLower(input) {
				return errors.New("Namespace already in use! ")
			}
		}
		return nil
	}

	// Prepare prompt
	dialogs.ClearScreen()
	namespace, err = dialogs.DialogPrompt("Enter namespace name", validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get name of new namespace!", err.Error())
		log.Error("[ProjectWizardAskForNamespace] Unable to get name of new namespace. %v\n", err)
		return namespace, err
	}

	return namespace, nil
}

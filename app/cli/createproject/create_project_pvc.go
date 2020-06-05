package createproject

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/logger"
)

func ProjectWizardAskForExistingPersistentVolumeClaim() (namespace string, err error) {
	log := logger.Log()
	// Validator for pvc
	validate := func(input string) error {
		// a pvc name can not be longer than 253 characters
		if len(input) > 253 {
			return errors.New("PVC name is too long! You can only use max. 253 characters. ")
		}
		return nil
	}

	// Prepare prompt
	dialogs.ClearScreen()
	namespace, err = dialogs.DialogPrompt("Enter existing Persistent Volume Claim (PVC) or leave empty for emptyDir", validate)
	// check if everything was ok
	if err != nil {
		log.Error("[ProjectWizardAskForNamespace] Unable to get persistent volume claim. %v\n", err)
	}

	return namespace, err
}

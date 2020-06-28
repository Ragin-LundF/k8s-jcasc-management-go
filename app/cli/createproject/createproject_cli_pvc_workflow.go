package createproject

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

func PersistentVolumeClaimWorkflow() (pvcName string, err error) {
	log := logger.Log()
	// Validator for pvc
	validate := validator.ValidatePersistentVolumeClaim

	// Prepare prompt
	dialogs.ClearScreen()
	pvcName, err = dialogs.DialogPrompt("Enter existing Persistent Volume Claim (PVC) or leave empty for emptyDir", validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get persistent volume claim.", err.Error())
		log.Errorf("[PersistentVolumeClaimWorkflow] Unable to get persistent volume claim. %s\n", err.Error())
		return pvcName, err
	}

	return pvcName, nil
}

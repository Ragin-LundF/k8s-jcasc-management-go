package createproject

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
	"k8s-management-go/app/utils/validator"
)

// PersistentVolumeClaimWorkflow represents the PVC workflow
func PersistentVolumeClaimWorkflow() (pvcName string, err error) {
	// Validator for pvc
	validate := validator.ValidatePersistentVolumeClaim

	// Prepare prompt
	dialogs.ClearScreen()
	pvcName, err = dialogs.DialogPrompt(constants.TextEnterExistingPvcOrLeaveEmpty, validate)
	// check if everything was ok
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(constants.LogUnableToGetPvc, err.Error())
		return pvcName, err
	}

	return pvcName, nil
}

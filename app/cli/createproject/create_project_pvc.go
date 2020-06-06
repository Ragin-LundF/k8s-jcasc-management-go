package createproject

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
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

// Replace PVC Name
func ProcessTemplatePvcExistingClaim(projectDirectory string, pvcName string) (success bool, err error) {
	log := logger.Log()
	jenkinsHelmValuesFile := files.AppendPath(projectDirectory, constants.FilenameJenkinsHelmValues)
	if files.FileOrDirectoryExists(jenkinsHelmValuesFile) {
		successful, err := files.ReplaceStringInFile(jenkinsHelmValuesFile, constants.TemplatePvcExistingVolumeClaim, pvcName)
		if !successful || err != nil {
			log.Error("[ProcessTemplatePvcExistingClaim] Can not replace PVC name in file [%v], \n%v", jenkinsHelmValuesFile, err)
			return false, err
		}
	}

	pvcValuesFile := files.AppendPath(projectDirectory, constants.FilenamePvcClaim)
	if files.FileOrDirectoryExists(pvcValuesFile) {
		successful, err := files.ReplaceStringInFile(pvcValuesFile, constants.TemplatePvcName, pvcName)
		if !successful || err != nil {
			log.Error("[ProcessTemplatePvcExistingClaim] Can not replace PVC name in file [%v], \n%v", pvcValuesFile, err)
			return false, err
		}
	}
	return true, err
}

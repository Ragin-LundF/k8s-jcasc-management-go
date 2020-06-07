package createproject

import (
	"fmt"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
)

// project wizard dialog for cloud templates
func ProjectWizardAskForCloudTemplates() (cloudTemplates []string, err error) {
	log := logger.Log()

	// look if cloud templates are available
	var cloudTemplatePath = files.AppendPath(models.GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)
	if !files.FileOrDirectoryExists(cloudTemplatePath) {
		loggingstate.AddInfoEntry("  -> No cloud template directory found. Skip this step.")
		log.Info("[ProjectWizardAskForCloudTemplates] No cloud template directory found. Skip this step.")

		return cloudTemplates, nil
	}

	// The cloud-templates directory is existing -> read files
	fileArray, err := files.ListFilesOfDirectory(cloudTemplatePath)
	if fileArray != nil && cap(*fileArray) > 0 {
		// Prepare prompt
		var cloudTemplateDialog dialogs.CloudTemplatesDialog
		cloudTemplateFiles := []string{"===== Select templates below or continue here ======"}
		cloudTemplateFiles = append(cloudTemplateFiles, *fileArray...)
		cloudTemplateDialog.CloudTemplateFiles = cloudTemplateFiles

		err = dialogs.DialogAskForCloudTemplates(&cloudTemplateDialog)
		cloudTemplates = cloudTemplateDialog.SelectedCloudTemplates

		fmt.Println(cloudTemplateDialog.SelectedCloudTemplates)

		// check if everything was ok
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Unable to get the cloud templates.", err.Error())
			log.Error("[ProjectWizardAskForCloudTemplates] Unable to get the cloud templates. %v\n", err)
			return cloudTemplates, err
		}
	} else {
		// no files found -> skip
		loggingstate.AddInfoEntry("  -> No cloud templates found. Skip this step")
		log.Info("[ProjectWizardAskForCloudTemplates] No cloud templates found. Skip this step")

		return cloudTemplates, nil
	}

	return cloudTemplates, nil
}

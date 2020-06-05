package createproject

import (
	"fmt"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
)

func ProjectWizardAskForCloudTemplates() (cloudTemplates []string, info string, err error) {
	log := logger.Log()

	// look if cloud templates are available
	cloudTemplatePath := files.AppendPath(config.GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)
	if !files.FileOrDirectoryExists(cloudTemplatePath) {
		info = info + constants.NewLine + "No cloud template directory found. Skip this step "
		log.Info("[ProjectWizardAskForCloudTemplates] %v", info)

		return cloudTemplates, info, err
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
			log.Error("[ProjectWizardAskForCloudTemplates] Unable to get the cloud templates. %v\n", err)
		}
	} else {
		// no files found -> skip
		info = info + constants.NewLine + "No cloud templates found. Skip this step "
		log.Info("[ProjectWizardAskForCloudTemplates] No cloud templates found. Skip this step")

		return cloudTemplates, info, err
	}

	return cloudTemplates, info, err
}

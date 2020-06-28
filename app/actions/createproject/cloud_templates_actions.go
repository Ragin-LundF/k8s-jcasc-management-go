package createproject

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
)

// read cloud templates and return list
func ActionReadCloudTemplates() (cloudTemplates []string) {
	// look if cloud templates are available
	var cloudTemplatePath = files.AppendPath(models.GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)
	if !files.FileOrDirectoryExists(cloudTemplatePath) {
		loggingstate.AddInfoEntry("  -> No cloud template directory found. Skip this step.")

		return cloudTemplates
	}

	// The cloud-templates directory is existing -> read files
	fileArray, _ := files.ListFilesOfDirectory(cloudTemplatePath)
	if fileArray != nil && cap(*fileArray) > 0 {
		// Prepare prompt
		cloudTemplates = append(cloudTemplates, *fileArray...)
	}

	return cloudTemplates
}

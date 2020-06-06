package createproject

import (
	"fmt"
	"github.com/goware/prefixer"
	"io/ioutil"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"strings"
)

// project wizard dialog for cloud templates
func ProjectWizardAskForCloudTemplates() (cloudTemplates []string, info string, err error) {
	log := logger.Log()

	// look if cloud templates are available
	var cloudTemplatePath = files.AppendPath(models.GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)
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

// add cloud templates to project template
func ProcessCloudTemplates(projectDirectory string, cloudTemplateFiles []string) (success bool, err error) {
	log := logger.Log()
	targetFile := files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode)
	// if file exists -> try to replace files
	if files.FileOrDirectoryExists(targetFile) {
		// first check if there are templates which should be processed
		if cap(cloudTemplateFiles) > 0 {
			// prepare vars and directory
			var cloudTemplateContent string
			var cloudTemplatePath = files.AppendPath(models.GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)

			// first read every template into a variable
			for _, cloudTemplate := range cloudTemplateFiles {
				cloudTemplateFileWithPath := files.AppendPath(cloudTemplatePath, cloudTemplate)
				read, err := ioutil.ReadFile(cloudTemplateFileWithPath)
				if err != nil {
					log.Error("[ProcessCloudTemplates] Can not read cloud template [%v] \n%v", cloudTemplateFileWithPath, err)
					return false, err
				}
				cloudTemplateContent = cloudTemplateContent + "\n" + string(read)
			}

			// Prefix the lines with spaces for correct yaml template
			prefixReader := prefixer.New(strings.NewReader(cloudTemplateContent), "          ")
			buffer, _ := ioutil.ReadAll(prefixReader)
			cloudTemplateContent = string(buffer)

			// replace target template
			success, err = files.ReplaceStringInFile(targetFile, constants.TemplateJenkinsCloudTemplates, cloudTemplateContent)
			if !success || err != nil {
				return false, err
			}
		} else {
			// replace placeholder
			success, err = files.ReplaceStringInFile(targetFile, constants.TemplateJenkinsCloudTemplates, "")
			if !success || err != nil {
				return false, err
			}
		}
	}
	return true, err
}

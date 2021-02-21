package createprojectactions

import (
	"fmt"
	"github.com/goware/prefixer"
	"io/ioutil"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// ActionReadCloudTemplates reads cloud templates and return list
func ActionReadCloudTemplates() (cloudTemplates []string) {
	// look if cloud templates are available
	var cloudTemplatePath = files.AppendPath(models.GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)
	if !files.FileOrDirectoryExists(cloudTemplatePath) {
		loggingstate.AddInfoEntry("  -> No cloud template directory found. Skip this step.")

		return cloudTemplates
	}

	// The cloud-templates directory is existing -> read files
	fileArray, _ := files.ListFilesOfDirectory(cloudTemplatePath)
	if fileArray != nil && len(*fileArray) > 0 {
		// Prepare prompt
		cloudTemplates = append(cloudTemplates, *fileArray...)
	}

	return cloudTemplates
}

// ActionProcessTemplateCloudTemplates adds cloud templates to project template
func ActionProcessTemplateCloudTemplates(projectDirectory string, cloudTemplateFiles []string) (success bool, err error) {
	targetFile := files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode)
	// if file exists -> try to replace files
	if files.FileOrDirectoryExists(targetFile) {
		// first check if there are templates which should be processed
		if len(cloudTemplateFiles) > 0 {
			cloudTemplateContent, err := ActionReadCloudTemplatesAsString(cloudTemplateFiles)
			if err != nil {
				loggingstate.AddErrorEntryAndDetails("  -> Unable to read cloud template files", err.Error())
				return false, err
			}

			// replace target template
			success, err = files.ReplaceStringInFile(targetFile, constants.TemplateJenkinsCloudTemplates, cloudTemplateContent)
			if !success || err != nil {
				loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to replace [%s] in [%s]", constants.TemplateJenkinsCloudTemplates, constants.FilenameJenkinsConfigurationAsCode), err.Error())
				return false, err
			}
		} else {
			// replace placeholder
			success, err = files.ReplaceStringInFile(targetFile, constants.TemplateJenkinsCloudTemplates, "")
			if !success || err != nil {
				loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to replace [%s] in [%s]", constants.TemplateJenkinsCloudTemplates, constants.FilenameJenkinsConfigurationAsCode), err.Error())
				return false, err
			}
		}
	}
	return true, nil
}

// ActionReadCloudTemplatesAsString : Rad cloud templates as string for further processing
func ActionReadCloudTemplatesAsString(cloudTemplateFiles []string) (cloudTemplateContent string, err error) {
	// prepare vars and directory
	var cloudTemplatePath = files.AppendPath(models.GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)

	// first read every template into a variable
	for _, cloudTemplate := range cloudTemplateFiles {
		cloudTemplateFileWithPath := files.AppendPath(cloudTemplatePath, cloudTemplate)
		read, err := ioutil.ReadFile(cloudTemplateFileWithPath)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to read cloud template [%s]", cloudTemplateFileWithPath), err.Error())
			return "", err
		}
		cloudTemplateContent = cloudTemplateContent + constants.NewLine + string(read)
	}

	// Prefix the lines with spaces for correct yaml template
	prefixReader := prefixer.New(strings.NewReader(cloudTemplateContent), "          ")
	buffer, _ := ioutil.ReadAll(prefixReader)
	cloudTemplateContent = string(buffer)

	return cloudTemplateContent, nil
}

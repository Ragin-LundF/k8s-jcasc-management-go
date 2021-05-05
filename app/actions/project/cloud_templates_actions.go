package project

import (
	"fmt"
	"github.com/goware/prefixer"
	"io/ioutil"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// ActionReadCloudTemplates reads cloud templates and return list
func ActionReadCloudTemplates() (cloudTemplates []string) {
	// look if cloud templates are available
	var cloudTemplatePath = files.AppendPath(configuration.GetConfiguration().GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)
	if !files.FileOrDirectoryExists(cloudTemplatePath) {
		loggingstate.AddInfoEntry("  -> No cloud template directory found. Skip this step.")

		return cloudTemplates
	}

	// The cloud-templates directory is existing -> read files
	var fileArray, _ = files.ListFilesOfDirectory(cloudTemplatePath)
	if fileArray != nil && len(*fileArray) > 0 {
		// Prepare prompt
		cloudTemplates = append(cloudTemplates, *fileArray...)
	}

	return cloudTemplates
}

// ActionReadCloudTemplatesAsString : Rad cloud templates as string for further processing
func ActionReadCloudTemplatesAsString(cloudTemplateFiles []string) (cloudTemplateContent string, err error) {
	// prepare vars and directory
	var cloudTemplatePath = files.AppendPath(configuration.GetConfiguration().GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)

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
	var prefixReader = prefixer.New(strings.NewReader(cloudTemplateContent), "          ")
	var buffer, _ = ioutil.ReadAll(prefixReader)
	cloudTemplateContent = string(buffer)

	return cloudTemplateContent, nil
}

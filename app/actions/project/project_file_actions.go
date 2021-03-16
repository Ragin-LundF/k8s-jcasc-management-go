package project

import (
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"path/filepath"
	"strings"
)

// ActionCreateNewProjectDirectory creates new project directory
func ActionCreateNewProjectDirectory(newProjectDir string) (err error) {
	// create directory
	err = os.MkdirAll(newProjectDir, os.ModePerm)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Failed to create a new project directory.", err.Error())
		return err
	}

	return nil
}

// ActionCopyTemplatesToNewDirectory copies files to new directory
func (prj *Project) ActionCopyTemplatesToNewDirectory(projectDirectory string) error {
	for _, file := range prj.CalculateRequiredDeploymentFiles() {
		if err := CopyTemplate(projectDirectory, file, false); err != nil {
			return err
		}
	}
	return nil
}

// CopyTemplate : copies a template to the target directory
func CopyTemplate(projectDirectory string, filename string, useTemplatePrefix bool) (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Copy [%s]...", filename))
	var targetFile string
	if useTemplatePrefix {
		targetFile = files.AppendPath(projectDirectory, constants.FilenameTempPrefix+filename)
	} else {
		targetFile = files.AppendPath(projectDirectory, filename)
	}

	_, err = files.CopyFile(
		files.AppendPath(configuration.GetConfiguration().GetProjectTemplateDirectory(), filename),
		targetFile,
	)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Copy [%s]...failed. See errors.", filename), err.Error())
		return err
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Copy [%s]...done", filename))
	return nil
}

// RemoveTempFile : removes a temporary file if it exists
func RemoveTempFile(tempFile string) {
	var _, file = filepath.Split(tempFile)
	if strings.HasPrefix(file, constants.FilenameTempPrefix) {
		if files.FileOrDirectoryExists(tempFile) {
			_ = os.Remove(tempFile)
		}
	}
}

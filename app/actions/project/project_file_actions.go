package project

import (
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
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
func (prj *Project) ActionCopyTemplatesToNewDirectory(projectDirectory string) (err error) {
	err = copyTemplates(prj.CalculateRequiredDeploymentFiles(), projectDirectory)
	return err
}

// copy the filenames
func copyTemplates(fileNames []string, projectDirectory string) (err error) {
	for _, fileName := range fileNames {
		loggingstate.AddInfoEntry(fmt.Sprintf("  -> Copy [%s]...", fileName))
		_, err = files.CopyFile(
			files.AppendPath(configuration.GetConfiguration().GetProjectTemplateDirectory(), fileName),
			files.AppendPath(projectDirectory, fileName),
		)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Copy [%s]...failed. See errors.", fileName), err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf("  -> Copy [%s]...done", fileName))
	}
	return nil
}

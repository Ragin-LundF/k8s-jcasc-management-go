package createproject

import (
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
)

// create new project directory
func ActionCreateNewProjectDirectory(newProjectDir string) (err error) {
	// create directory
	err = os.MkdirAll(newProjectDir, os.ModePerm)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Failed to create a new project directory.", err.Error())
		return err
	}

	return nil
}

// copy files to new directory
func ActionCopyTemplatesToNewDirectory(projectDirectory string, copyPersistentVolume bool, createDeploymentOnlyProject bool) (err error) {
	var fileNamesToCopy []string

	// copy nginx-ingress-controller values.yaml
	fileNamesToCopy = append(fileNamesToCopy, constants.FilenameNginxIngressControllerHelmValues)

	// if it is not a deployment only project, copy more files
	if !createDeploymentOnlyProject {
		// copy Jenkins values.yaml
		fileNamesToCopy = append(fileNamesToCopy, constants.FilenameJenkinsHelmValues)
		// copy Jenkins JCasC config.yaml
		fileNamesToCopy = append(fileNamesToCopy, constants.FilenameJenkinsConfigurationAsCode)
		// copy existing PVC values.yaml
		if copyPersistentVolume {
			fileNamesToCopy = append(fileNamesToCopy, constants.FilenamePvcClaim)
		}
		// copy secrets to project
		if models.GetConfiguration().GlobalSecretsFile == "" {
			fileNamesToCopy = append(fileNamesToCopy, constants.FilenameSecrets)
		}
	}
	err = copyTemplates(fileNamesToCopy, projectDirectory)
	return err
}

// copy the filenames
func copyTemplates(fileNames []string, projectDirectory string) (err error) {
	for _, fileName := range fileNames {
		loggingstate.AddInfoEntry(fmt.Sprintf("  -> Copy [%s]...", fileName))
		_, err = files.CopyFile(
			files.AppendPath(models.GetProjectTemplateDirectory(), fileName),
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

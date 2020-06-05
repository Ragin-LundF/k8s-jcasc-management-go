package createproject

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os"
)

// Processing project creation
func ProcessProjectCreate(namespace string, ipAddress string, jenkinsSystemMsg string, jobsCfgRepo string, existingPvc string, selectedCloudTemplates []string, createDeploymentOnlyProject bool) (info string, err error) {
	// calculate the target directory
	newProjectDir := files.AppendPath(models.GetProjectBaseDirectory(), namespace)

	// create new project directory
	infoLog, err := createNewProjectDirectory(newProjectDir)
	info = info + constants.NewLine + infoLog
	if err != nil {
		return info, err
	}

	// copy necessary files
	infoLog, err = copyTemplatesToNewDirectory(newProjectDir, len(existingPvc) > 0, createDeploymentOnlyProject)
	info = info + constants.NewLine + infoLog
	if err != nil {
		os.RemoveAll(newProjectDir)
		return info, err
	}

	// add IP and namespace to IP configuration
	successful, err := config.AddToIpConfigFile(namespace, ipAddress)
	if !successful || err != nil {
		os.RemoveAll(newProjectDir)
		return info, err
	}

	// processing cloud templates
	successful, err = ProcessCloudTemplates(newProjectDir, selectedCloudTemplates)
	if !successful || err != nil {
		os.RemoveAll(newProjectDir)
		return info, err
	}

	return info, err
}

// create new project directory
func createNewProjectDirectory(newProjectDir string) (info string, err error) {
	log := logger.Log()
	log.Info("[createNewProjectDirectory] Trying to create a new project directory...")
	info = "Trying to create a new project directory..."

	// create directory
	err = os.Mkdir(newProjectDir, os.ModePerm)
	if err != nil {
		log.Error("[createNewProjectDirectory] Trying to create a new project directory [%v]...error. \n%v", newProjectDir, err)
		info = info + constants.NewLine + "Error while creating project directory."
		return info, err
	}
	// successful
	log.Info("[createNewProjectDirectory] Trying to create a new project directory...done")
	info = info + constants.NewLine + "Trying to create a new project directory...done"

	return info, err
}

// copy files to new directory
func copyTemplatesToNewDirectory(projectDirectory string, copyPersistentVolume bool, createDeploymentOnlyProject bool) (info string, err error) {
	templateDirectory := models.GetProjectTemplateDirectory()
	// copy nginx-ingress-controller values.yaml
	_, err = files.CopyFile(
		files.AppendPath(templateDirectory, constants.FilenameNginxIngressControllerHelmValues),
		files.AppendPath(projectDirectory, constants.FilenameNginxIngressControllerHelmValues),
	)
	if err != nil {
		return info, err
	}

	// if it is not a deployment only project, copy more files
	if !createDeploymentOnlyProject {
		// copy Jenkins values.yaml
		_, err = files.CopyFile(
			files.AppendPath(templateDirectory, constants.FilenameJenkinsHelmValues),
			files.AppendPath(projectDirectory, constants.FilenameJenkinsHelmValues),
		)
		if err != nil {
			return info, err
		}

		// copy Jenkins JCasC config.yaml
		_, err = files.CopyFile(
			files.AppendPath(templateDirectory, constants.FilenameJenkinsConfigurationAsCode),
			files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode),
		)
		if err != nil {
			return info, err
		}

		// copy existing PVC values.yaml
		if copyPersistentVolume {
			_, err = files.CopyFile(
				files.AppendPath(templateDirectory, constants.FilenamePvcClaim),
				files.AppendPath(projectDirectory, constants.FilenamePvcClaim),
			)
			if err != nil {
				return info, err
			}
		}

		// copy secrets to project
		if models.GetConfiguration().GlobalSecretsFile == "" {
			_, err = files.CopyFile(
				files.AppendPath(templateDirectory, constants.FilenameSecrets),
				files.AppendPath(projectDirectory, constants.FilenameSecrets),
			)
			if err != nil {
				return info, err
			}
		}
	}

	return info, err
}

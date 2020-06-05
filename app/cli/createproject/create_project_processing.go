package createproject

import (
	"fmt"
	"io"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os"
)

func ProcessProjectCreate(namespace string, ipAddress string, jenkinsSystemMsg string, jobsCfgRepo string, existingPvc string, selectedCloudTemplates []string, createDeploymentOnlyProject bool) (info string, err error) {
	// calculate the target directory
	newProjectDir := files.AppendPath(config.GetProjectBaseDirectory(), namespace)

	// create new project directory
	info, err = createNewProjectDirectory(newProjectDir)
	if err != nil {
		return info, err
	}

	info, err = copyTemplatesToNewDirectory(newProjectDir, len(existingPvc) > 0, createDeploymentOnlyProject)
	// dummy processing
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
	templateDirectory := config.GetProjectTemplateDirectory()
	// copy nginx-ingress-controller values.yaml
	_, err = copyFile(
		files.AppendPath(templateDirectory, constants.FilenameNginxIngressControllerHelmValues),
		files.AppendPath(projectDirectory, constants.FilenameNginxIngressControllerHelmValues),
	)
	if err != nil {
		return info, err
	}

	// if it is not a deployment only project, copy more files
	if !createDeploymentOnlyProject {
		// copy Jenkins values.yaml
		_, err = copyFile(
			files.AppendPath(templateDirectory, constants.FilenameJenkinsHelmValues),
			files.AppendPath(projectDirectory, constants.FilenameJenkinsHelmValues),
		)
		if err != nil {
			return info, err
		}

		// copy Jenkins JCasC config.yaml
		_, err = copyFile(
			files.AppendPath(templateDirectory, constants.FilenameJenkinsConfigurationAsCode),
			files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode),
		)
		if err != nil {
			return info, err
		}

		// copy existing PVC values.yaml
		if copyPersistentVolume {
			_, err = copyFile(
				files.AppendPath(templateDirectory, constants.FilenamePvcClaim),
				files.AppendPath(projectDirectory, constants.FilenamePvcClaim),
			)
			if err != nil {
				return info, err
			}
		}

		// copy secrets to project
		if config.GetConfiguration().GlobalSecretsFile == "" {
			_, err = copyFile(
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

func copyFile(src string, dst string) (bytesWritten int64, err error) {
	srcFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !srcFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()
	nBytes, err := io.Copy(dstFile, srcFile)

	return nBytes, err
}

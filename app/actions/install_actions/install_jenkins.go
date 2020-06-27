package install_actions

import (
	"errors"
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

// install_actions Jenkins with Helm
func ActionHelmInstallJenkins(command string, namespace string, deploymentName string) (err error) {
	log := logger.Log()
	log.Infof("[Install Jenkins] Try to %s Jenkins on namespace [%s] with deployment name [%s]...", command, namespace, deploymentName)
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to %s Jenkins on namespace [%s] with deployment name [%s]...", command, namespace, deploymentName))

	// check if command is ok
	if command == constants.HelmCommandInstall || command == constants.HelmCommandUpgrade {
		// prepare files and directories
		helmChartsJenkinsDirectory := models.FilePathWithBasePath(constants.DirHelmJenkinsMaster)
		helmChartsJenkinsValuesFile := files.AppendPath(
			files.AppendPath(
				models.GetProjectBaseDirectory(),
				namespace,
			),
			constants.FilenameJenkinsHelmValues,
		)

		// execute Helm command
		argsForCommand := []string{
			deploymentName,
			helmChartsJenkinsDirectory,
		}
		if models.GetConfiguration().K8sManagement.DryRunOnly {
			argsForCommand = append(argsForCommand, "--dry-run", "--debug")
		}
		argsForCommand = append(argsForCommand, "-n", namespace, "-f", helmChartsJenkinsValuesFile)

		// executing jenkins helm install_actions
		log.Infof("[Install Jenkins] Start installing/upgrading Jenkins with Helm on namespace [%s]...", namespace)
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Start installing/upgrading Jenkins with Helm on namespace [%s]...", namespace))
		err := helm.ExecutorHelm(command, argsForCommand)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("-> Unable to install_actions/upgrade Jenkins on namespace [%s] with deployment name [%s]", namespace, deploymentName), err.Error())
			log.Errorf("[Install Jenkins] Unable to install_actions/upgrade Jenkins on namespace [%s] with deployment name [%s]. Errors: \n%s", namespace, deploymentName, err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Start installing/upgrading Jenkins with Helm on namespace [%s]...done", namespace))
		log.Infof("[Install Jenkins] Start installing/upgrading Jenkins with Helm on namespace [%s]...done", namespace)
	} else {
		// helm command was wrong -> abort
		log.Errorf("[Install Jenkins] Try to install_actions/upgrade Jenkins on namespace [%s] with deployment name [%s]...failed. Wrong command [%s]", namespace, deploymentName, command)
		loggingstate.AddErrorEntry(fmt.Sprintf("-> Try to install_actions/upgrade Jenkins on namespace [%s] with deployment name [%s]...Wrong command [%s]", namespace, deploymentName, command))
		return errors.New("Helm command [" + command + "] unknown.")
	}
	log.Infof("[Install Jenkins] Try to %s Jenkins on namespace [%s] with deployment name [%s]...done", command, namespace, deploymentName)
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to %s Jenkins on namespace [%s] with deployment name [%s]...done", command, namespace, deploymentName))

	return err
}

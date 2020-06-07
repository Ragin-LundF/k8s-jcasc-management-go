package install

import (
	"errors"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/logger"
)

// install Jenkins with Helm
func HelmInstallJenkins(command string, namespace string, deploymentName string) (err error) {
	log := logger.Log()
	log.Info("[Install Jenkins] Try to %v Jenkins on namespace [%v] with deployment name [%v]...", command, namespace, deploymentName)
	loggingstate.AddInfoEntry("-> Try to " + command + " Jenkins on namespace [" + namespace + "] with deployment name [" + deploymentName + "]...")

	// check if command is ok
	if command == constants.HelmCommandInstall || command == constants.HelmCommandUpgrade {
		// prepare files and directories
		helmChartsJenkinsDirectory := files.AppendPath(models.GetConfiguration().BasePath, constants.DirHelmJenkinsMaster)
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

		// executing jenkins helm install
		log.Info("[Install Jenkins] Start installing/upgrading Jenkins with Helm on namespace [%v]...", namespace)
		loggingstate.AddInfoEntry("-> Start installing/upgrading Jenkins with Helm on namespace [" + namespace + "]...")
		err := helm.ExecutorHelm(command, argsForCommand)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("-> Unable to install/upgrade Jenkins on namespace ["+namespace+"] with deployment name ["+deploymentName+"]", err.Error())
			log.Error("[Install Jenkins] Unable to install/upgrade Jenkins on namespace [%v] with deployment name [%v]. Errors: \n%v", namespace, deploymentName, err)
			return err
		}
		loggingstate.AddInfoEntry("-> Start installing/upgrading Jenkins with Helm on namespace [" + namespace + "]...done")
		log.Info("[Install Jenkins] Start installing/upgrading Jenkins with Helm on namespace [%v]...done", namespace)
	} else {
		// helm command was wrong -> abort
		log.Error("[Install Jenkins] Try to install/upgrade Jenkins on namespace [%v] with deployment name [%v]...failed. Wrong command [%v]", namespace, deploymentName, command)
		loggingstate.AddErrorEntry("-> Try to install/upgrade Jenkins on namespace [" + namespace + "] with deployment name [" + deploymentName + "]...Wrong command [" + command + "]")
		return errors.New("Helm command [" + command + "] unknown.")
	}
	log.Info("[Install Jenkins] Try to %v Jenkins on namespace [%v] with deployment name [%v]...done", command, namespace, deploymentName)
	loggingstate.AddInfoEntry("-> Try to " + command + " Jenkins on namespace [" + namespace + "] with deployment name [" + deploymentName + "]...done")

	return err
}

package installactions

import (
	"errors"
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/loggingstate"
)

// ActionHelmInstallJenkins installs Jenkins with Helm
func ActionHelmInstallJenkins(command string, namespace string, deploymentName string) (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to %s Jenkins on namespace [%s] with deployment name [%s]...", command, namespace, deploymentName))

	// check if command is ok
	if command == constants.HelmCommandInstall || command == constants.HelmCommandUpgrade {
		// prepare files and directories
		helmChartsJenkinsDirectory := configuration.GetConfiguration().FilePathWithBasePath(constants.DirHelmJenkinsMaster)
		helmChartsJenkinsValuesFile := files.AppendPath(
			files.AppendPath(
				configuration.GetConfiguration().GetProjectBaseDirectory(),
				namespace,
			),
			constants.FilenameJenkinsHelmValues,
		)

		// execute Helm command
		argsForCommand := []string{
			deploymentName,
			helmChartsJenkinsDirectory,
		}
		if configuration.GetConfiguration().K8SManagement.DryRunOnly {
			argsForCommand = append(argsForCommand, "--dry-run", "--debug")
		}
		argsForCommand = append(argsForCommand, "-n", namespace, "-f", helmChartsJenkinsValuesFile)

		// executing jenkins helm install
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Start installing/upgrading Jenkins with Helm on namespace [%s]...", namespace))
		err := helm.ExecutorHelm(command, argsForCommand)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("-> Unable to install/upgrade Jenkins on namespace [%s] with deployment name [%s]", namespace, deploymentName), err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Start installing/upgrading Jenkins with Helm on namespace [%s]...done", namespace))
	} else {
		// helm command was wrong -> abort
		loggingstate.AddErrorEntry(fmt.Sprintf("-> Try to install/upgrade Jenkins on namespace [%s] with deployment name [%s]...Wrong command [%s]", namespace, deploymentName, command))
		return errors.New("Helm command [" + command + "] unknown.")
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to %s Jenkins on namespace [%s] with deployment name [%s]...done", command, namespace, deploymentName))

	return err
}

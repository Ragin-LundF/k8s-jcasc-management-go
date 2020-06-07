package install

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/logger"
)

// install Jenkins with Helm
func HelmInstallJenkins(command string, namespace string, deploymentName string) (info string, err error) {
	log := logger.Log()
	log.Info("[Install Jenkins] Try to install Jenkins on namespace [" + namespace + "] with deployment name [" + deploymentName + "]...")
	info = info + constants.NewLine + "Try to install Jenkins..."

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

		log.Info("[Install Jenkins] Start installing Jenkins with Helm...")
		err := helm.ExecutorHelm(command, argsForCommand)
		info = info + constants.NewLine + infoLog

		// first write output of dry-run...
		if models.GetConfiguration().K8sManagement.DryRunOnly {
			log.Info("[Install Jenkins] Output of dry-run for namespace [" + namespace + "]")
			log.Info(helmCmdOutput)
		}
		if err != nil {
			log.Error("[Install Jenkins] Cannot install Jenkins on namespace [" + namespace + "] with deployment name [" + deploymentName + "]")
			info = "Jenkins installation not successful. See errors." + constants.NewLine + info
			return info, err
		}
		log.Info("[Install Jenkins] Done installing Jenkins with Helm...")
		info = info + constants.NewLine + "Helm output:"
		info = info + constants.NewLine + helmCmdOutput
	} else {
		// helm command was wrong -> abort
		log.Error("[Install Jenkins] Helm command [" + command + "] unknown.")
		return info, errors.New("Helm command [" + command + "] unknown.")
	}

	return info, err
}

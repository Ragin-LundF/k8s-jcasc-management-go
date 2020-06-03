package install

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os/exec"
)

// install Jenkins with Helm
func HelmInstallJenkins(command string, deploymentName string, namespace string) (info string, err error) {
	log := logger.Log()
	info = info + "Try to install Jenkins..."
	// check if command is ok
	if command == constants.HelmCommandInstall || command == constants.HelmCommandUpgrade {
		// prepare files and directories
		helmChartsJenkinsDirectory := files.AppendPath(config.GetConfiguration().BasePath, constants.DirHelmJenkinsMaster)
		helmChartsJenkinsValuesFile := files.AppendPath(
			files.AppendPath(
				config.GetProjectBaseDirectory(),
				namespace,
			),
			constants.FilenameJenkinsHelmValues,
		) // execute Helm command
		outputCmd, err := exec.Command("helm", command, deploymentName, helmChartsJenkinsDirectory, "-n", namespace, "-f", helmChartsJenkinsValuesFile).Output()
		if err != nil {
			log.Error(err)
			return info, err
		}

		info = info + "\nHelm Jenkins install output:"
		info = info + "\n==============="
		info = info + string(outputCmd)
		info = info + "\n==============="
	} else {
		// helm command was wrong -> abort
		return info, errors.New("Helm command [" + command + "] unknown.")
	}
	return info, err
}

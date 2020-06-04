package install

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os/exec"
	"strings"
)

// install Jenkins with Helm
func HelmInstallJenkins(command string, namespace string, deploymentName string) (info string, err error) {
	log := logger.Log()

	info = info + constants.NewLine + "Try to install Jenkins..."
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
		)

		// execute Helm command
		argsForCommand := []string{
			command,
			deploymentName,
			helmChartsJenkinsDirectory,
		}
		if config.GetConfiguration().K8sManagement.DryRunOnly {
			argsForCommand = append(argsForCommand, "--dry-run", "--debug")
		}
		argsForCommand = append(argsForCommand, "-n", namespace, "-f", helmChartsJenkinsValuesFile)

		cmd := exec.Command("helm", argsForCommand...)
		log.Info("Executing command: [helm " + strings.Join(argsForCommand, " ") + "]")

		outputCmd, err := cmd.CombinedOutput()
		if err != nil {
			log.Error("Failed to execute: " + cmd.String())
			info = info + constants.NewLine + "Jenkins installation aborted. See errors."
			err = errors.New(string(outputCmd) + constants.NewLine + err.Error())
			return info, err
		}

		info = info + constants.NewLine + "Helm Jenkins install output:"
		info = info + constants.NewLine + "==============="
		info = info + string(outputCmd)
		info = info + constants.NewLine + "==============="
	} else {
		// helm command was wrong -> abort
		return info, errors.New("Helm command [" + command + "] unknown.")
	}

	return info, err
}

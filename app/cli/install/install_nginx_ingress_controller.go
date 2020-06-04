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
func HelmInstallNginxIngressController(command string, namespace string, jenkinsIngressEnabled bool) (info string, err error) {
	log := logger.Log()

	info = info + constants.NewLine + "Try to install Nginx Ingress Controller..."

	// create var with path to ingress controller helm values
	helmChartsNginxIngressCtrlValuesFile := files.AppendPath(
		files.AppendPath(
			config.GetProjectBaseDirectory(),
			namespace,
		),
		constants.FilenameNginxIngressControllerHelmValues,
	)

	// check if nginx ingress controller helm values are existing
	// if this is the case -> install it
	if files.FileOrDirectoryExists(helmChartsNginxIngressCtrlValuesFile) {
		// check if command is ok
		if command == constants.HelmCommandInstall || command == constants.HelmCommandUpgrade {
			// prepare files and directories
			helmChartsNginxIngressCtrlDirectory := files.AppendPath(config.GetConfiguration().BasePath, constants.DirHelmJenkinsMaster)
			// execute Helm command
			nginxIngressCtrlDeploymentName := config.GetConfiguration().Nginx.Ingress.Controller.DeploymentName

			// if Jenkins is not active for this namespace, then disable Jenkins ingress routing
			var jenkinsIngressSetParam = ""
			if jenkinsIngressEnabled == false {
				info = info + constants.NewLine + "Disabling Jenkins ingress routing, because Jenkins is not available."
				jenkinsIngressSetParam = "--set k8sJenkinsMgmt.ingress.enabled=false"
			}
			// execute command
			cmd := exec.Command("helm", command, nginxIngressCtrlDeploymentName, helmChartsNginxIngressCtrlDirectory, "-n", namespace, "-f", helmChartsNginxIngressCtrlValuesFile, jenkinsIngressSetParam)
			outputCmd, err := cmd.CombinedOutput()
			if err != nil {
				log.Error("Failed to execute: " + cmd.String())
				info = info + constants.NewLine + "Nginx Ingress Controller installation aborted. See errors."
				err = errors.New(string(outputCmd) + constants.NewLine + err.Error())
				return info, err
			}

			info = info + constants.NewLine + "Helm Nginx Ingress Controller install output:"
			info = info + constants.NewLine + "==============="
			info = info + string(outputCmd)
			info = info + constants.NewLine + "==============="
		} else {
			// helm command was wrong -> abort
			return info, errors.New("Helm command [" + command + "] unknown.")
		}
	} else {
		info = info + constants.NewLine + "No nginx-ingress-controller values.yaml for namespace [" + namespace + "] found."
		info = info + constants.NewLine + "Skipping nginx-ingress-controller values.yaml installation."
	}

	return info, err
}

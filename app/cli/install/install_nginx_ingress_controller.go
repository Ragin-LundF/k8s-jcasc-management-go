package install

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/logger"
	"strconv"
)

// install Jenkins with Helm
func HelmInstallNginxIngressController(command string, namespace string, jenkinsIngressEnabled bool) (info string, err error) {
	log := logger.Log()
	log.Info("[Install NginxIngressCtrl] Trying to install nginx-ingress-controller on namespace [" + namespace + "] while Jenkins exists [" + strconv.FormatBool(jenkinsIngressEnabled) + "]...")
	info = info + constants.NewLine + "Trying to install nginx-ingress-controller..."

	// create var with path to ingress controller helm values
	helmChartsNginxIngressCtrlValuesFile := files.AppendPath(
		files.AppendPath(
			models.GetProjectBaseDirectory(),
			namespace,
		),
		constants.FilenameNginxIngressControllerHelmValues,
	)

	// check if nginx ingress controller helm values are existing
	// if this is the case -> install it
	if files.FileOrDirectoryExists(helmChartsNginxIngressCtrlValuesFile) {
		log.Info("[Install NginxIngressCtrl] nginx-ingress-controller values.yaml file found for namespace [" + namespace + "].")
		// check if command is ok
		if command == constants.HelmCommandInstall || command == constants.HelmCommandUpgrade {
			// prepare files and directories
			helmChartsNginxIngressCtrlDirectory := files.AppendPath(models.GetConfiguration().BasePath, constants.DirHelmNginxIngressCtrl)
			// execute Helm command
			nginxIngressCtrlDeploymentName := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName

			// execute Helm command
			argsForCommand := []string{
				nginxIngressCtrlDeploymentName,
				helmChartsNginxIngressCtrlDirectory,
				"-n", namespace,
				"-f", helmChartsNginxIngressCtrlValuesFile,
			}

			// if Jenkins is not active for this namespace, then disable Jenkins ingress routing
			if jenkinsIngressEnabled == false {
				log.Info("[Install NginxIngressCtrl] Jenkins is not available for the namespace [" + namespace + "]. Disabling Jenkins ingress routing.")
				info = info + constants.NewLine + "Jenkins is not available for the namespace [" + namespace + "]. Disabling Jenkins ingress routing."
				argsForCommand = append(argsForCommand, "--set", "k8sJenkinsMgmt.ingress.enabled=false")
			}

			// add dry-run and debug if necessary
			if models.GetConfiguration().K8sManagement.DryRunOnly {
				argsForCommand = append(argsForCommand, "--dry-run", "--debug")
			}

			// execute command
			log.Info("[Install NginxIngressCtrl] Start installing nginx-ingress-controller on namespace [" + namespace + "] with Helm...")
			helmCmdOutput, infoLog, err := helm.ExecutorHelm(command, argsForCommand)
			info = info + constants.NewLine + infoLog

			// first write output of dry-run...
			if models.GetConfiguration().K8sManagement.DryRunOnly {
				log.Info("[Install NginxIngressCtrl] Output of dry-run for namespace [" + namespace + "]")
				log.Info(helmCmdOutput)
			}

			if err != nil {
				info = info + constants.NewLine + "nginx-ingress-controller installation aborted. See errors."
				return info, err
			}
			log.Info("[Install NginxIngressCtrl] Done installing Jenkins on namespace [" + namespace + "] with Helm...")
			info = info + constants.NewLine + "nginx-ingress-controller install Helm output:"
			info = info + constants.NewLine + helmCmdOutput
		} else {
			// helm command was wrong -> abort
			log.Error("[Install NginxIngressCtrl] nginx-ingress-controller installation: Helm command [" + command + "] unknown.")
			return info, errors.New("nginx-ingress-controller installation: Helm command [" + command + "] unknown.")
		}
	} else {
		log.Info("[Install NginxIngressCtrl] No nginx-ingress-controller values.yaml file found for namespace [" + namespace + "].")
		log.Info("[Install NginxIngressCtrl] Skipping nginx-ingress-controller installation for namespace [" + namespace + "].")
		info = info + constants.NewLine + "No nginx-ingress-controller values.yaml for namespace [" + namespace + "] found."
		info = info + constants.NewLine + "Skipping nginx-ingress-controller values.yaml installation."
	}

	return info, err
}

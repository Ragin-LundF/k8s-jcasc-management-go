package install

import (
	"errors"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/logger"
	"strconv"
)

// install Jenkins with Helm
func HelmInstallNginxIngressController(command string, namespace string, jenkinsIngressEnabled bool) (err error) {
	log := logger.Log()
	log.Info("[Install NginxIngressCtrl] Trying to %v nginx-ingress-controller on namespace [%v] while Jenkins exists [%v]...", command, namespace, strconv.FormatBool(jenkinsIngressEnabled))
	loggingstate.AddInfoEntry("[Install NginxIngressCtrl] Trying to " + command + " nginx-ingress-controller on namespace [" + namespace + "] while Jenkins exists [" + strconv.FormatBool(jenkinsIngressEnabled) + "]...")

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
		log.Info("[Install NginxIngressCtrl] nginx-ingress-controller values.yaml file found for namespace [%v].", namespace)
		loggingstate.AddInfoEntry("[Install NginxIngressCtrl] nginx-ingress-controller values.yaml file found for namespace [" + namespace + "].")
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
				log.Info("[Install NginxIngressCtrl] Jenkins is not available for the namespace [%v]. Disabling Jenkins ingress routing.", namespace)
				loggingstate.AddInfoEntry("[Install NginxIngressCtrl] Jenkins is not available for the namespace [" + namespace + "]. Disabling Jenkins ingress routing.")
				argsForCommand = append(argsForCommand, "--set", "k8sJenkinsMgmt.ingress.enabled=false")
			}

			// add dry-run and debug if necessary
			if models.GetConfiguration().K8sManagement.DryRunOnly {
				argsForCommand = append(argsForCommand, "--dry-run", "--debug")
			}

			// execute command
			log.Info("[Install NginxIngressCtrl] Start installing/upgrading nginx-ingress-controller with Helm on namespace [%v]...", namespace)
			loggingstate.AddInfoEntry("-> Start installing/upgrading nginx-ingress-controller with Helm on namespace [" + namespace + "]...")
			err := helm.ExecutorHelm(command, argsForCommand)
			if err != nil {
				loggingstate.AddErrorEntryAndDetails("-> Unable to install/upgrade nginx-ingress-controller on namespace ["+namespace+"]", err.Error())
				log.Error("[Install NginxIngressCtrl] Unable to install/upgrade nginx-ingress-controller on namespace [%v]. Errors: \n%v", namespace, err)
				return err
			}
			loggingstate.AddInfoEntry("-> Start installing/upgrading nginx-ingress-controller with Helm on namespace [" + namespace + "]...done")
			log.Info("[Install NginxIngressCtrl] Start installing/upgrading nginx-ingress-controller with Helm on namespace [%v]...done", namespace)
		} else {
			// helm command was wrong -> abort
			log.Error("[Install NginxIngressCtrl] Try to install/upgrade nginx-ingress-controller on namespace [%v]...failed. Wrong command [%v]", namespace, command)
			loggingstate.AddErrorEntry("-> Try to install/upgrade nginx-ingress-controller on namespace [" + namespace + "]...Wrong command [" + command + "]")
			return errors.New("Helm command [" + command + "] unknown.")
		}
	} else {
		loggingstate.AddInfoEntry("[Install NginxIngressCtrl] No nginx-ingress-controller values.yaml file found for namespace [" + namespace + "]. Skip installing.")
		log.Info("[Install NginxIngressCtrl] No nginx-ingress-controller values.yaml file found for namespace [%v]. Skip installing.", namespace)
	}

	loggingstate.AddInfoEntry("[Install NginxIngressCtrl] Trying to " + command + " nginx-ingress-controller on namespace [" + namespace + "] while Jenkins exists [" + strconv.FormatBool(jenkinsIngressEnabled) + "]...done")
	log.Info("[Install NginxIngressCtrl] Trying to %v nginx-ingress-controller on namespace [%v] while Jenkins exists [%v]...done", command, namespace, strconv.FormatBool(jenkinsIngressEnabled))

	return err
}

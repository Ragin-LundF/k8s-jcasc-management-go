package install_actions

import (
	"errors"
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/loggingstate"
	"strconv"
)

// install_actions Jenkins with Helm
func ActionHelmInstallNginxIngressController(command string, namespace string, jenkinsIngressEnabled bool) (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf("[Install NginxIngressCtrl] Trying to %s nginx-ingress-controller on namespace [%s] while Jenkins exists [%s]...", command, namespace, strconv.FormatBool(jenkinsIngressEnabled)))

	// create var with path to ingress controller helm values
	helmChartsNginxIngressCtrlValuesFile := files.AppendPath(
		files.AppendPath(
			models.GetProjectBaseDirectory(),
			namespace,
		),
		constants.FilenameNginxIngressControllerHelmValues,
	)

	// check if nginx ingress controller helm values are existing
	// if this is the case -> install_actions it
	if files.FileOrDirectoryExists(helmChartsNginxIngressCtrlValuesFile) {
		loggingstate.AddInfoEntry(fmt.Sprintf("[Install NginxIngressCtrl] nginx-ingress-controller values.yaml file found for namespace [%s].", namespace))
		// check if command is ok
		if command == constants.HelmCommandInstall || command == constants.HelmCommandUpgrade {
			// prepare files and directories
			helmChartsNginxIngressCtrlDirectory := models.FilePathWithBasePath(constants.DirHelmNginxIngressCtrl)
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
				loggingstate.AddInfoEntry(fmt.Sprintf("[Install NginxIngressCtrl] Jenkins is not available for the namespace [%s]. Disabling Jenkins ingress routing.", namespace))
				argsForCommand = append(argsForCommand, "--set", "k8sJenkinsMgmt.ingress.enabled=false")
			}

			// add dry-run and debug if necessary
			if models.GetConfiguration().K8sManagement.DryRunOnly {
				argsForCommand = append(argsForCommand, "--dry-run", "--debug")
			}

			// execute command
			loggingstate.AddInfoEntry(fmt.Sprintf("-> Start installing/upgrading nginx-ingress-controller with Helm on namespace [%s]...", namespace))
			err := helm.ExecutorHelm(command, argsForCommand)
			if err != nil {
				loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("-> Unable to install_actions/upgrade nginx-ingress-controller on namespace [%s]", namespace), err.Error())
				return err
			}
			loggingstate.AddInfoEntry(fmt.Sprintf("-> Start installing/upgrading nginx-ingress-controller with Helm on namespace [%s]...done", namespace))
		} else {
			// helm command was wrong -> abort
			loggingstate.AddErrorEntry(fmt.Sprintf("-> Try to install_actions/upgrade nginx-ingress-controller on namespace [%s]...Wrong command [%s]", namespace, command))
			return errors.New(fmt.Sprintf("Helm command [%s] unknown.", command))
		}
	} else {
		loggingstate.AddInfoEntry(fmt.Sprintf("[Install NginxIngressCtrl] No nginx-ingress-controller values.yaml file found for namespace [%s]. Skip installing.", namespace))
	}

	loggingstate.AddInfoEntry(fmt.Sprintf("[Install NginxIngressCtrl] Trying to %s nginx-ingress-controller on namespace [%s] while Jenkins exists [%s]...done", command, namespace, strconv.FormatBool(jenkinsIngressEnabled)))

	return err
}

package install

import (
	"fmt"
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/loggingstate"
	"strconv"
)

// ActionHelmInstallNginxIngressController installs Jenkins with Helm
func (projectConfig *ProjectConfig) ActionHelmInstallNginxIngressController() (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"[Install NginxIngressCtrl] Trying to %s nginx-ingress-controller on namespace [%s] while Jenkins exists [%s]...",
		projectConfig.HelmCommand,
		projectConfig.Project.Base.Namespace,
		strconv.FormatBool(projectConfig.Project.CalculateIfDeploymentFileIsRequired(constants.FilenameJenkinsHelmValues))))

	// create var with path to ingress controller helm values
	var helmChartsNginxIngressCtrlValuesFile string
	helmChartsNginxIngressCtrlValuesFile, err = projectConfig.PrepareInstallYAML(constants.FilenameNginxIngressControllerHelmValues)

	// check if nginx ingress controller helm values are existing
	// if this is the case -> install it
	if files.FileOrDirectoryExists(helmChartsNginxIngressCtrlValuesFile) {
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"[Install NginxIngressCtrl] nginx-ingress-controller values.yaml file found for namespace [%s].",
			projectConfig.Project.Base.Namespace))
		// check if command is ok
		if projectConfig.HelmCommand == constants.HelmCommandInstall || projectConfig.HelmCommand == constants.HelmCommandUpgrade {
			// prepare files and directories
			var helmChartsNginxIngressCtrlDirectory = configuration.GetConfiguration().FilePathWithBasePath(constants.DirHelmNginxIngressCtrl)
			// execute Helm command
			var nginxIngressCtrlDeploymentName = configuration.GetConfiguration().Nginx.Ingress.Deployment.DeploymentName

			// execute Helm command
			var argsForCommand = []string{
				nginxIngressCtrlDeploymentName,
				helmChartsNginxIngressCtrlDirectory,
				"-n", projectConfig.Project.Base.Namespace,
				"-f", helmChartsNginxIngressCtrlValuesFile,
			}

			// if Jenkins is not active for this namespace, then disable Jenkins ingress routing
			if projectConfig.Project.CalculateIfDeploymentFileIsRequired(constants.FilenameJenkinsHelmValues) == false {
				loggingstate.AddInfoEntry(fmt.Sprintf(
					"[Install NginxIngressCtrl] Jenkins is not available for the namespace [%s]. Disabling Jenkins ingress routing.",
					projectConfig.Project.Base.Namespace))
				argsForCommand = append(argsForCommand, "--set", "k8sJenkinsMgmt.ingress.enabled=false")
			}

			// add dry-run and debug if necessary
			if configuration.GetConfiguration().K8SManagement.DryRunOnly {
				argsForCommand = append(argsForCommand, "--dry-run", "--debug")
			}

			// execute command
			loggingstate.AddInfoEntry(fmt.Sprintf(
				"-> Start installing/upgrading nginx-ingress-controller with Helm on namespace [%s]...",
				projectConfig.Project.Base.Namespace))
			err = helm.ExecutorHelm(projectConfig.HelmCommand, argsForCommand)

			project.RemoveTempFile(helmChartsNginxIngressCtrlValuesFile)
			if err != nil {
				loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
					"-> Unable to install/upgrade nginx-ingress-controller on namespace [%s]",
					projectConfig.Project.Base.Namespace), err.Error())
				return err
			}
			loggingstate.AddInfoEntry(fmt.Sprintf(
				"-> Start installing/upgrading nginx-ingress-controller with Helm on namespace [%s]...done",
				projectConfig.Project.Base.Namespace))
		} else {
			// helm command was wrong -> abort
			loggingstate.AddErrorEntry(fmt.Sprintf(
				"-> Try to install/upgrade nginx-ingress-controller on namespace [%s]...Wrong command [%s]",
				projectConfig.Project.Base.Namespace,
				projectConfig.HelmCommand))
			return fmt.Errorf("Helm command [%s] unknown. ", projectConfig.HelmCommand)
		}
	}

	loggingstate.AddInfoEntry(fmt.Sprintf(
		"[Install NginxIngressCtrl] No nginx-ingress-controller values.yaml file found for namespace [%s]. Skip installing.",
		projectConfig.Project.Base.Namespace))
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"[Install NginxIngressCtrl] Trying to %s nginx-ingress-controller on namespace [%s] while Jenkins exists [%s]...done",
		projectConfig.HelmCommand,
		projectConfig.Project.Base.Namespace,
		strconv.FormatBool(projectConfig.Project.CalculateIfDeploymentFileIsRequired(constants.FilenameJenkinsHelmValues))))

	project.RemoveTempFile(helmChartsNginxIngressCtrlValuesFile)

	return err
}

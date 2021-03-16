package install

import (
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
)

// ProcessJenkinsUninstallIfExists processes Jenkins uninstall if Jenkins Helm values exists
func (projectConfig *ProjectConfig) ProcessJenkinsUninstallIfExists() (err error) {
	if projectConfig.Project.CalculateIfDeploymentFileIsRequired(constants.FilenameJenkinsHelmValues) {
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"-> Uninstalling deployment [%s] on namespace [%s]...",
			projectConfig.Project.Base.DeploymentName,
			projectConfig.Project.Base.Namespace))
		if err = projectConfig.ActionHelmUninstallJenkins(); err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
				"  -> Unable to uninstall deployment [%s] on namespace [%s]...failed",
				projectConfig.Project.Base.DeploymentName,
				projectConfig.Project.Base.Namespace), err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"-> Uninstalling deployment [%s] on namespace [%s]...done",
			projectConfig.Project.Base.DeploymentName,
			projectConfig.Project.Base.Namespace))
	}
	return nil
}

// ProcessNginxIngressControllerUninstall processes the nginx ingress controller uninstall
func (projectConfig *ProjectConfig) ProcessNginxIngressControllerUninstall() (err error) {
	if projectConfig.Project.CalculateIfDeploymentFileIsRequired(constants.FilenameNginxIngressControllerHelmValues) {
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"-> Uninstalling nginx-ingress-controller on namespace [%s]...",
			projectConfig.Project.Base.Namespace))
		if err = projectConfig.ActionHelmUninstallNginxIngressController(); err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
				"-> Uninstalling nginx-ingress-controller on namespace [%s]...abort",
				projectConfig.Project.Base.Namespace), err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"-> Uninstalling nginx-ingress-controller on namespace [%s]...done",
			projectConfig.Project.Base.Namespace))
	}
	return nil
}

// ProcessScriptsUninstallIfExists processes the uninstall of the scripts if it exists
func (projectConfig *ProjectConfig) ProcessScriptsUninstallIfExists() {
	// try to uninstall scripts
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"-> Try to execute uninstall scripts on [%s]...",
		projectConfig.Project.Base.Namespace))
	// we ignore errors. They will be logged, but we keep on doing the uninstall for the scripts
	_ = projectConfig.ActionShellScriptsUninstall()
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"-> Try to execute uninstall scripts on [%s]...done",
		projectConfig.Project.Base.Namespace))
}

// ProcessK8sCleanup processes the K8S cleanup
func (projectConfig *ProjectConfig) ProcessK8sCleanup() {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"-> Try to cleanup configuration of [%s]...",
		projectConfig.Project.Base.Namespace))
	projectConfig.ActionCleanupK8sNginxIngressController()
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"-> Try to cleanup configuration of [%s]...done",
		projectConfig.Project.Base.Namespace))
}

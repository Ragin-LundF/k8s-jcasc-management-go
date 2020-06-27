package uninstall_actions

import (
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
)

func JenkinsUninstallIfExists(state models.StateData) (err error) {
	if state.JenkinsHelmValuesExist {
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling deployment [%s] on namespace [%s]...", state.DeploymentName, state.Namespace))
		if err = HelmUninstallJenkins(state.Namespace, state.DeploymentName); err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to uninstall deployment [%s] on namespace [%s]...failed", state.DeploymentName, state.Namespace), err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling deployment [%s] on namespace [%s]...done", state.DeploymentName, state.Namespace))
	}
	return nil
}

func NginxIngressControllerUninstall(state models.StateData) (err error) {
	if state.NginxHelmValuesExist {
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling nginx-ingress-controller on namespace [%s]...", state.Namespace))
		if err = HelmUninstallNginxIngressController(state.Namespace); err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("-> Uninstalling nginx-ingress-controller on namespace [%s]...abort", state.Namespace), err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling nginx-ingress-controller on namespace [%s]...done", state.Namespace))
	}
	return nil
}

func ScriptsUninstallIfExists(state models.StateData) {
	// try to uninstall scripts
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute uninstall scripts on [%s]...", state.Namespace))
	// we ignore errors. They will be logged, but we keep on doing the uninstall for the scripts
	_ = ShellScriptsUninstall(state.Namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute uninstall scripts on [%s]...done", state.Namespace))
}

func K8sCleanup(state models.StateData) {
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to cleanup configuration of [%s]...", state.Namespace))
	CleanupK8sNginxIngressController(state.Namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to cleanup configuration of [%s]...done", state.Namespace))
}

func CheckNginxDirectoryExists(state models.StateData) models.StateData {
	nginxHelmValueFile := files.AppendPath(
		files.AppendPath(
			models.GetProjectBaseDirectory(),
			state.Namespace,
		),
		constants.FilenameNginxIngressControllerHelmValues,
	)
	state.NginxHelmValuesExist = files.FileOrDirectoryExists(nginxHelmValueFile)
	return state
}

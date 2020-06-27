package uninstall_actions

import (
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
)

func ProcessJenkinsUninstallIfExists(state models.StateData) (err error) {
	if state.JenkinsHelmValuesExist {
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling deployment [%s] on namespace [%s]...", state.DeploymentName, state.Namespace))
		if err = ActionHelmUninstallJenkins(state.Namespace, state.DeploymentName); err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to uninstall deployment [%s] on namespace [%s]...failed", state.DeploymentName, state.Namespace), err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling deployment [%s] on namespace [%s]...done", state.DeploymentName, state.Namespace))
	}
	return nil
}

func ProcessNginxIngressControllerUninstall(state models.StateData) (err error) {
	if state.NginxHelmValuesExist {
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling nginx-ingress-controller on namespace [%s]...", state.Namespace))
		if err = ActionHelmUninstallNginxIngressController(state.Namespace); err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("-> Uninstalling nginx-ingress-controller on namespace [%s]...abort", state.Namespace), err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling nginx-ingress-controller on namespace [%s]...done", state.Namespace))
	}
	return nil
}

func ProcessScriptsUninstallIfExists(state models.StateData) {
	// try to uninstall scripts
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute uninstall scripts on [%s]...", state.Namespace))
	// we ignore errors. They will be logged, but we keep on doing the uninstall for the scripts
	_ = ActionShellScriptsUninstall(state.Namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute uninstall scripts on [%s]...done", state.Namespace))
}

func ProcessK8sCleanup(state models.StateData) {
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to cleanup configuration of [%s]...", state.Namespace))
	ActionCleanupK8sNginxIngressController(state.Namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to cleanup configuration of [%s]...done", state.Namespace))
}

func ProcessCheckNginxDirectoryExists(state models.StateData) models.StateData {
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

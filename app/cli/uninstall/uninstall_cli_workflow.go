package uninstall

import (
	"fmt"
	"k8s-management-go/app/actions/uninstall_actions"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
)

// workflow for uninstall
func DoUninstall() (err error) {
	bar := dialogs.CreateProgressBar("Uninstalling Jenkins", 5)
	// ask for namespace
	loggingstate.AddInfoEntry("Starting Uninstall...")
	state, err := ShowUninstallDialogs()
	if err != nil {
		return err
	}

	bar.Describe("Uninstalling Jenkins deployment...")
	if state.JenkinsHelmValuesExist {
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling deployment [%s] on namespace [%s]...", state.DeploymentName, state.Namespace))
		if err = uninstall_actions.HelmUninstallJenkins(state.Namespace, state.DeploymentName); err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to uninstall deployment [%s] on namespace [%s]...failed", state.DeploymentName, state.Namespace), err.Error())
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling deployment [%s] on namespace [%s]...done", state.DeploymentName, state.Namespace))
	}
	_ = bar.Add(1)

	// uninstall nginx ingress controller
	bar.Describe("Check for Nginx installation...")
	nginxHelmValueFile := files.AppendPath(
		files.AppendPath(
			models.GetProjectBaseDirectory(),
			state.Namespace,
		),
		constants.FilenameNginxIngressControllerHelmValues,
	)
	nginxHelmValuesExist := files.FileOrDirectoryExists(nginxHelmValueFile)
	_ = bar.Add(1)

	bar.Describe("Nginx-ingress-controller found...Uninstalling...")
	if nginxHelmValuesExist {
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling nginx-ingress-controller on namespace [%s]...", state.Namespace))
		if err = uninstall_actions.HelmUninstallNginxIngressController(state.Namespace); err != nil {
			return err
		}
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Uninstalling nginx-ingress-controller on namespace [%s]...done", state.Namespace))
	}
	_ = bar.Add(1)

	// in dry-run we do not want to uninstall the scripts
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		bar.Describe("Try to execute uninstall scripts...")
		// try to uninstall scripts
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute uninstall scripts on [%s]...", state.Namespace))
		// we ignore errors. They will be logged, but we keep on doing the uninstall for the scripts
		_ = uninstall_actions.ShellScriptsUninstall(state.Namespace)
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute uninstall scripts on [%s]...done", state.Namespace))
		_ = bar.Add(1)

		// nginx-ingress-controller cleanup
		bar.Describe("Cleanup configuration...")
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to cleanup configuration of [%s]...", state.Namespace))
		uninstall_actions.CleanupK8sNginxIngressController(state.Namespace)
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to cleanup configuration of [%s]...done", state.Namespace))
		_ = bar.Add(1)
	} else {
		_ = bar.Add(2) // skipping previous steps
	}

	loggingstate.AddInfoEntry("Starting Uninstall...done")
	return nil
}

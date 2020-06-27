package uninstall

import (
	"k8s-management-go/app/actions/uninstall_actions"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

// workflow for uninstall
func DoUninstall() (err error) {
	bar := dialogs.CreateProgressBar("Uninstalling Jenkins", 5)
	// Show dialogs to catch needed information
	loggingstate.AddInfoEntry("Starting Uninstall...")
	state, err := ShowUninstallDialogs()
	if err != nil {
		return err
	}

	bar.Describe("Uninstalling Jenkins deployment...")
	err = uninstall_actions.JenkinsUninstallIfExists(state)
	if err != nil {
		return err
	}
	_ = bar.Add(1)

	// uninstall nginx ingress controller
	bar.Describe("Check for Nginx installation...")
	state = uninstall_actions.CheckNginxDirectoryExists(state)
	_ = bar.Add(1)

	bar.Describe("Nginx-ingress-controller found...Uninstalling...")
	err = uninstall_actions.NginxIngressControllerUninstall(state)
	if err != nil {
		return err
	}
	_ = bar.Add(1)

	// in dry-run we do not want to uninstall the scripts
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		bar.Describe("Try to execute uninstall scripts...")
		uninstall_actions.ScriptsUninstallIfExists(state)
		_ = bar.Add(1)

		// nginx-ingress-controller cleanup
		bar.Describe("Cleanup configuration...")
		uninstall_actions.K8sCleanup(state)
		_ = bar.Add(1)
	} else {
		_ = bar.Add(2) // skipping previous steps
	}

	loggingstate.AddInfoEntry("Starting Uninstall...done")
	return nil
}

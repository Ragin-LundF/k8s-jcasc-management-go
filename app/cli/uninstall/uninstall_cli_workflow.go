package uninstall

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/loggingstate"
)

// DoUninstall is the workflow for uninstall
func DoUninstall() (err error) {
	var bar = dialogs.CreateProgressBar("Uninstalling Jenkins", 5)
	// Show dialogs to catch needed information
	loggingstate.AddInfoEntry("Starting Uninstall...")
	projectConfig, err := ShowUninstallDialogs()
	if err != nil {
		return err
	}

	bar.Describe("Uninstalling Jenkins deployment...")
	err = projectConfig.ProcessJenkinsUninstallIfExists()
	if err != nil {
		return err
	}
	_ = bar.Add(1)

	// uninstall nginx ingress controller
	bar.Describe("Check for Nginx installation...")
	projectConfig.ProcessCheckNginxDirectoryExists()
	_ = bar.Add(1)

	bar.Describe("Nginx-ingress-controller found...Uninstalling...")
	err = projectConfig.ProcessNginxIngressControllerUninstall()
	if err != nil {
		return err
	}
	_ = bar.Add(1)

	// in dry-run we do not want to uninstall the scripts
	if !configuration.GetConfiguration().K8SManagement.DryRunOnly {
		bar.Describe("Try to execute uninstall scripts...")
		projectConfig.ProcessScriptsUninstallIfExists()
		_ = bar.Add(1)

		// nginx-ingress-controller cleanup
		bar.Describe("Cleanup configuration...")
		projectConfig.ProcessK8sCleanup()
		_ = bar.Add(1)
	} else {
		_ = bar.Add(2) // skipping previous steps
	}

	loggingstate.AddInfoEntry("Starting Uninstall...done")
	return nil
}

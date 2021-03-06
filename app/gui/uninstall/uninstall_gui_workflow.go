package uninstall

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"k8s-management-go/app/actions/uninstallactions"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

// ExecuteUninstallWorkflow executes the workflow
func ExecuteUninstallWorkflow(window fyne.Window, state models.StateData) (err error) {
	// Progress Bar
	var progressCnt = 1
	var progressMaxCnt = 4
	var bar = dialog.NewProgress(state.HelmCommand, "Uninstalling on namespace "+state.Namespace, window)
	bar.Show()

	// uninstall Jenkins if exists
	err = uninstallactions.ProcessJenkinsUninstallIfExists(state)
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		return err
	}

	// uninstall nginx ingress controller
	state = uninstallactions.ProcessCheckNginxDirectoryExists(state)

	// uninstall Nginx ingress controller is exists
	err = uninstallactions.ProcessNginxIngressControllerUninstall(state)
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		return err
	}

	// in dry-run we do not want to uninstall the scripts
	if !configuration.GetConfiguration().K8SManagement.DryRunOnly {
		uninstallactions.ProcessScriptsUninstallIfExists(state)
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++

		// nginx-ingress-controller cleanup
		uninstallactions.ProcessK8sCleanup(state)
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++
	} else {
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++
		progressCnt++
	}

	loggingstate.AddInfoEntry("Starting Uninstall...done")
	bar.Hide()
	return err
}

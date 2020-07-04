package uninstall

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"k8s-management-go/app/actions/uninstallactions"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

// ExecuteUninstallWorkflow executes the workflow
func ExecuteUninstallWorkflow(window fyne.Window, state models.StateData) (err error) {
	// Progress Bar
	progressCnt := 1
	progressMaxCnt := 4
	bar := dialog.NewProgress(state.HelmCommand, "Uninstalling on namespace "+state.Namespace, window)
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
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
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

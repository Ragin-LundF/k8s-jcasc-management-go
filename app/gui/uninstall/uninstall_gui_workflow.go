package uninstall

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"k8s-management-go/app/actions/uninstall_actions"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/gui/ui_elements"
	"k8s-management-go/app/models"
)

// execute the workflow
func ExecuteUninstallWorkflow(window fyne.Window, state models.StateData) (err error) {
	// Progress Bar
	progressCnt := 1
	progressMaxCnt := 4
	bar := dialog.NewProgress(state.HelmCommand, "Uninstalling on namespace "+state.Namespace, window)
	bar.Show()

	// uninstall Jenkins if exists
	err = uninstall_actions.JenkinsUninstallIfExists(state)
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		return err
	}

	// uninstall nginx ingress controller
	state = uninstall_actions.CheckNginxDirectoryExists(state)

	// uninstall Nginx ingress controller is exists
	err = uninstall_actions.NginxIngressControllerUninstall(state)
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		return err
	}

	// in dry-run we do not want to uninstall the scripts
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		uninstall_actions.ScriptsUninstallIfExists(state)
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++

		// nginx-ingress-controller cleanup
		uninstall_actions.K8sCleanup(state)
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++
	} else {
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++
		progressCnt++
	}

	loggingstate.AddInfoEntry("Starting Uninstall...done")
	bar.Hide()
	ui_elements.ShowLogOutput(window)
	return err
}

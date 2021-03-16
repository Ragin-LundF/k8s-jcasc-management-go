package uninstall

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/loggingstate"
)

// ExecuteUninstallWorkflow executes the workflow
func ExecuteUninstallWorkflow(window fyne.Window, projectConfig install.ProjectConfig) (err error) {
	// Progress Bar
	var progressCnt = 1
	var progressMaxCnt = 4
	var bar = dialog.NewProgress(
		projectConfig.HelmCommand,
		"Uninstalling on namespace "+projectConfig.Project.Base.Namespace,
		window)
	bar.Show()

	// uninstall Jenkins if exists
	err = projectConfig.ProcessJenkinsUninstallIfExists()
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		return err
	}

	// uninstall Nginx ingress controller is exists
	err = projectConfig.ProcessNginxIngressControllerUninstall()
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		return err
	}

	// in dry-run we do not want to uninstall the scripts
	if !configuration.GetConfiguration().K8SManagement.DryRunOnly {
		projectConfig.ProcessScriptsUninstallIfExists()
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++

		// nginx-ingress-controller cleanup
		projectConfig.ProcessK8sCleanup()
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

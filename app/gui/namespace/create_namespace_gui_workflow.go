package namespace

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/actions/namespaceactions"
	"time"
)

// ExecuteCreateNamespaceWorkflow executes the create namespace workflow
func ExecuteCreateNamespaceWorkflow(window fyne.Window, projectConfig install.ProjectConfig) (err error) {
	// Progress Bar
	var bar = dialog.NewProgress(
		projectConfig.HelmCommand,
		"Creating namespace "+projectConfig.Project.Base.Namespace,
		window)
	bar.Show()
	err = namespaceactions.ProcessNamespaceCreation(projectConfig)
	bar.SetValue(1)

	// wait 1 second to show user the dialog
	time.Sleep(time.Duration(1) * time.Second)
	bar.Hide()
	return err
}

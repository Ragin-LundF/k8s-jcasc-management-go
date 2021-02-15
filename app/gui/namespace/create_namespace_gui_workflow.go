package namespace

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/models"
	"time"
)

// ExecuteCreateNamespaceWorkflow executes the create namespace workflow
func ExecuteCreateNamespaceWorkflow(window fyne.Window, state models.StateData) (err error) {
	// Progress Bar
	bar := dialog.NewProgress(state.HelmCommand, "Creating namespace "+state.Namespace, window)
	bar.Show()
	err = namespaceactions.ProcessNamespaceCreation(state)
	bar.SetValue(1)

	// wait 1 second to show user the dialog
	time.Sleep(time.Duration(1) * time.Second)
	bar.Hide()
	return err
}

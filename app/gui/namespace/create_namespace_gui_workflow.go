package namespace

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"k8s-management-go/app/actions/namespace_actions"
	"k8s-management-go/app/models"
	"time"
)

func ExecuteCreateNamespaceWorkflow(window fyne.Window, state models.StateData) (err error) {
	// Progress Bar
	bar := dialog.NewProgress(state.HelmCommand, "Creating namespace "+state.Namespace, window)
	bar.Show()
	err = namespace_actions.ProcessNamespaceCreation(state)
	bar.SetValue(1)

	// wait 1 second to show user the dialog
	time.Sleep(1000)
	bar.Hide()
	return err
}

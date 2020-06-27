package namespace

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"k8s-management-go/app/actions/install_actions"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
	"time"
)

func ExecuteCreateNamespaceWorkflow(window fyne.Window, state models.StateData) (err error) {
	// Progress Bar
	bar := dialog.NewProgress(state.HelmCommand, "Creating namespace "+state.Namespace, window)
	bar.Show()

	loggingstate.AddInfoEntry("Start creating namespace...")

	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...")
	if err = install_actions.CheckAndCreateNamespace(state.Namespace); err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Check and create namespace if necessary...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...done")
	loggingstate.AddInfoEntry("Start creating namespace...done")
	bar.SetValue(1)
	time.Sleep(1000)
	bar.Hide()
	return nil
}

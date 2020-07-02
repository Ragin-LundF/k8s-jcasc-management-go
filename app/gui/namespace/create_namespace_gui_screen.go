package namespace

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/models"
)

// ScreenNamespaceCreate shows the create namespace screen
func ScreenNamespaceCreate(window fyne.Window) fyne.CanvasObject {
	var namespace string

	// Namespace
	namespaceErrorLabel := widget.NewLabel("")
	namespaceSelectEntry := uielements.CreateNamespaceSelectEntry(namespaceErrorLabel)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "", Widget: namespaceErrorLabel},
		},
		OnSubmit: func() {
			// get variables
			namespace = namespaceSelectEntry.Text

			// map state
			state := models.StateData{
				Namespace: namespace,
			}

			_ = ExecuteCreateNamespaceWorkflow(window, state)
			// show output
			uielements.ShowLogOutput(window)
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)

	return box
}

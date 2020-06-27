package namespace

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/ui_elements"
	"k8s-management-go/app/models"
)

func ScreenNamespaceCreate(window fyne.Window) fyne.CanvasObject {
	var namespace string

	// Namespace
	namespaceErrorLabel := widget.NewLabel("")
	namespaceSelectEntry := ui_elements.CreateNamespaceSelectEntry(namespaceErrorLabel)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "", Widget: namespaceErrorLabel},
		},
		OnSubmit: func() {
			// get variables
			namespace = namespaceSelectEntry.Text
			if !ui_elements.ValidateNamespace(namespace) {
				namespaceErrorLabel.SetText("Error: namespace is unknown!")
				namespaceErrorLabel.Show()
				return
			}

			// map state
			state := models.StateData{
				Namespace: namespace,
			}

			_ = ExecuteCreateNamespaceWorkflow(window, state)
			// show output
			ui_elements.ShowLogOutput(window)
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)

	return box
}

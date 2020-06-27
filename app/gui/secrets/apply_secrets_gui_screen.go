package secrets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/secrets_actions"
	"k8s-management-go/app/gui/ui_elements"
)

// apply to all namespaces
func ScreenApplySecretsToAllNamespace(window fyne.Window) fyne.CanvasObject {
	// secrets password
	passwordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			//_ = secrets_actions.ActionDecryptSecretsFile(passwordEntry.Text)
			ui_elements.ShowLogOutput(window)
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)
	return box
}

// apply to one selected namespace
func ScreenApplySecretsToNamespace(window fyne.Window) fyne.CanvasObject {
	// Namespace
	namespaceErrorLabel := widget.NewLabel("")
	namespaceSelectEntry := ui_elements.CreateNamespaceSelectEntry(namespaceErrorLabel)

	// password
	passwordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			_ = secrets_actions.ActionApplySecretsToNamespace(namespaceSelectEntry.Text)
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

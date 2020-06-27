package secrets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/secrets_actions"
	"k8s-management-go/app/gui/ui_elements"
)

func ScreenDecryptSecrets(window fyne.Window) fyne.CanvasObject {
	// secrets password
	passwordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			_ = secrets_actions.ActionDecryptSecretsFile(passwordEntry.Text)
			ui_elements.ShowLogOutput(window)
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)

	return box
}

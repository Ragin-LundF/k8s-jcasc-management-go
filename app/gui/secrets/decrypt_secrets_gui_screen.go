package secrets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/secretsactions"
	"k8s-management-go/app/gui/uielements"
)

// ScreenDecryptSecrets shows the decrypt secrets screen
func ScreenDecryptSecrets(window fyne.Window) fyne.CanvasObject {
	// secrets password
	passwordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			_ = secretsactions.ActionDecryptSecretsFile(passwordEntry.Text)
			uielements.ShowLogOutput(window)
		},
	}

	box := widget.NewVBox(
		form,
	)

	return box
}

package secrets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"k8s-management-go/app/actions/secretsactions"
	"k8s-management-go/app/gui/uielements"
)

// ScreenDecryptSecrets shows the decrypt secrets screen
func ScreenDecryptSecrets(window fyne.Window) fyne.CanvasObject {
	// secrets password
	var secretsFiles = uielements.CreateSecretsFileEntry()
	var passwordEntry = widget.NewPasswordEntry()

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Secrets file", Widget: secretsFiles},
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			_ = secretsactions.ActionDecryptSecretsFile(passwordEntry.Text, secretsFiles.Selected)
			uielements.ShowLogOutput(window)
		},
	}

	return container.NewVBox(
		widget.NewLabel(""),
		form,
	)
}

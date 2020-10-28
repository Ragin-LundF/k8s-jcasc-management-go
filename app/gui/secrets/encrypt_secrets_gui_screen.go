package secrets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/secretsactions"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/utils/validator"
)

// ScreenEncryptSecrets shows the encrypt secrets screen
func ScreenEncryptSecrets(window fyne.Window) fyne.CanvasObject {
	// UI elements
	var secretsFiles = uielements.CreateSecretsFileEntry()
	var passwordErrorLabel = widget.NewLabel("")
	// secrets password
	var passwordEntry = widget.NewPasswordEntry()
	var confirmPasswordEntry = widget.NewPasswordEntry()

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Secrets file", Widget: secretsFiles},
			{Text: "Password", Widget: passwordEntry},
			{Text: "Confirm password", Widget: confirmPasswordEntry},
			{Text: "", Widget: passwordErrorLabel},
		},
		OnSubmit: func() {
			isValid, errMessage := validator.ValidateConfirmPasswords(passwordEntry.Text, confirmPasswordEntry.Text)
			passwordErrorLabel.SetText(errMessage)
			if isValid {
				_ = secretsactions.ActionEncryptSecretsFile(passwordEntry.Text, secretsFiles.Selected)
				uielements.ShowLogOutput(window)
			}
		},
	}

	return widget.NewVBox(
		widget.NewLabel(""),
		form,
	)
}

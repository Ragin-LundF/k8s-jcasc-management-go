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
	passwordErrorLabel := widget.NewLabel("")
	// secrets password
	passwordEntry := widget.NewPasswordEntry()
	confirmPasswordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Password", Widget: passwordEntry},
			{Text: "Confirm password", Widget: confirmPasswordEntry},
			{Text: "", Widget: passwordErrorLabel},
		},
		OnSubmit: func() {
			isValid, errMessage := validator.ValidateConfirmPasswords(passwordEntry.Text, confirmPasswordEntry.Text)
			passwordErrorLabel.SetText(errMessage)
			if isValid {
				_ = secretsactions.ActionEncryptSecretsFile(passwordEntry.Text)
				uielements.ShowLogOutput(window)
			}
		},
	}

	box := widget.NewVBox(
		form,
	)

	return box
}

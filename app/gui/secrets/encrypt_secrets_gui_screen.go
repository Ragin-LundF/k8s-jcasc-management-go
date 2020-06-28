package secrets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/secrets_actions"
	"k8s-management-go/app/gui/ui_elements"
	"k8s-management-go/app/utils/validator"
)

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
				_ = secrets_actions.ActionEncryptSecretsFile(passwordEntry.Text)
				ui_elements.ShowLogOutput(window)
			}
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)

	return box
}

package secrets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/secretsactions"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/models"
)

// ScreenApplySecretsToAllNamespace shows the apply to all namespaces screen
func ScreenApplySecretsToAllNamespace(window fyne.Window) fyne.CanvasObject {
	// secrets password
	passwordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			// first try to decrypt the file
			if err := secretsactions.ActionDecryptSecretsFile(passwordEntry.Text); err == nil {
				// execute the file and apply to all namespaces
				bar := uielements.ProgressBar{
					Bar:        dialog.NewProgress("Apply secrets to all namespaces", "Progress", window),
					CurrentCnt: 0,
					MaxCount:   float64(len(models.GetIPConfiguration().Ips)),
				}
				bar.Bar.Show()
				_ = secretsactions.ActionApplySecretsToAllNamespaces(bar.AddCallback)
				bar.Bar.Hide()
			}

			uielements.ShowLogOutput(window)
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)
	return box
}

// ScreenApplySecretsToNamespace shows the apply to one selected namespace screen
func ScreenApplySecretsToNamespace(window fyne.Window) fyne.CanvasObject {
	// Namespace
	namespaceErrorLabel := widget.NewLabel("")
	namespaceSelectEntry := uielements.CreateNamespaceSelectEntry(namespaceErrorLabel)

	// password
	passwordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			// first try to decrypt the file
			err := secretsactions.ActionDecryptSecretsFile(passwordEntry.Text)
			if err == nil {
				// execute the file
				_ = secretsactions.ActionApplySecretsToNamespace(namespaceSelectEntry.Text)
			}
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

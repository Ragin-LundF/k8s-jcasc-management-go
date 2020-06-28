package secrets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/secrets_actions"
	"k8s-management-go/app/gui/ui_elements"
	"k8s-management-go/app/models"
)

type progressBar struct {
	Bar        *dialog.ProgressDialog
	MaxCount   float64
	CurrentCnt float64
}

// function to add progress. Will be used as callback
func (progress *progressBar) AddCallback() {
	progress.Bar.SetValue(float64(1) / progress.MaxCount * progress.CurrentCnt)
	progress.CurrentCnt = progress.CurrentCnt + 1
}

// apply to all namespaces
func ScreenApplySecretsToAllNamespace(window fyne.Window) fyne.CanvasObject {
	// secrets password
	passwordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			// first try to decrypt the file
			if err := secrets_actions.ActionDecryptSecretsFile(passwordEntry.Text); err == nil {
				// execute the file and apply to all namespaces
				bar := progressBar{
					Bar:        dialog.NewProgress("Apply secrets to all namespaces", "Progress", window),
					CurrentCnt: 0,
					MaxCount:   float64(len(models.GetIpConfiguration().Ips)),
				}
				bar.Bar.Show()
				_ = secrets_actions.ActionApplySecretsToAllNamespaces(bar.AddCallback)
				bar.Bar.Hide()
			}

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
			// first try to decrypt the file
			err := secrets_actions.ActionDecryptSecretsFile(passwordEntry.Text)
			if err == nil {
				// execute the file
				_ = secrets_actions.ActionApplySecretsToNamespace(namespaceSelectEntry.Text)
			}
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

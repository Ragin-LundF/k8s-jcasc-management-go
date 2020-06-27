package install

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/install_actions"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/gui/ui_elements"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/validator"
)

func ScreenInstall(window fyne.Window) fyne.CanvasObject {
	var namespace string
	var deploymentName string
	var installTypeOption string
	var dryRunOption string
	var secretsPasswords string

	// Namespace
	namespaceErrorLabel := widget.NewLabel("")
	namespaceSelectEntry := ui_elements.CreateNamespaceSelectEntry(namespaceErrorLabel)

	// Deployment name
	deploymentNameEntry := ui_elements.CreateDeploymentNameEntry()

	// Install or update
	installTypeRadio := ui_elements.CreateInstallTypeRadio()

	// Dry-run or execute
	dryRunRadio := ui_elements.CreateDryRunRadio()

	// secrets password
	secretsPasswordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "Deployment Name", Widget: deploymentNameEntry},
			{Text: "Installation type", Widget: installTypeRadio},
			{Text: "Execute or dry run", Widget: dryRunRadio},
		},
		OnSubmit: func() {
			// get variables
			namespace = namespaceSelectEntry.Text
			deploymentName = deploymentNameEntry.Text
			installTypeOption = installTypeRadio.Selected
			dryRunOption = dryRunRadio.Selected
			if dryRunOption == constants.InstallDryRunActive {
				models.AssignDryRun(true)
			} else {
				models.AssignDryRun(false)
			}
			if !validator.ValidateNamespace(namespace) {
				namespaceErrorLabel.SetText("Error: namespace is unknown!")
				namespaceErrorLabel.Show()
				return
			}

			// map state
			state := models.StateData{
				Namespace:       namespace,
				DeploymentName:  deploymentName,
				HelmCommand:     installTypeOption,
				SecretsPassword: &secretsPasswords,
			}

			// Directories
			err, state := install_actions.CalculateDirectoriesForInstall(state, state.Namespace)
			if err != nil {
				dialog.ShowError(err, window)
			}

			// Check Jenkins directories
			state = install_actions.CheckJenkinsDirectories(state)

			// ask for password
			if dryRunOption == constants.InstallDryRunInactive {
				openSecretsPasswordDialog(window, secretsPasswordEntry, state)
			} else {
				_ = ExecuteInstallWorkflow(window, state)
				// show output
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

// Secrets password dialog
func openSecretsPasswordDialog(window fyne.Window, secretsPasswordEntry *widget.Entry, state models.StateData) {
	secretsPasswordWindow := widget.NewForm(widget.NewFormItem("Password", secretsPasswordEntry))
	secretsPasswordWindow.Resize(fyne.Size{Width: 300, Height: 90})

	dialog.ShowCustomConfirm("Password for Secrets...", "Ok", "Cancel", secretsPasswordWindow, func(confirmed bool) {
		if confirmed {
			state.SecretsPassword = &secretsPasswordEntry.Text
			_ = ExecuteInstallWorkflow(window, state)
		} else {
			return
		}
	}, window)
}

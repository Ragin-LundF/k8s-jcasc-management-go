package uninstall

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

func ScreenUninstall(window fyne.Window) fyne.CanvasObject {
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

	// Dry-run or execute
	dryRunRadio := ui_elements.CreateDryRunRadio()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "Deployment Name", Widget: deploymentNameEntry},
			{Text: "Execute or dry run", Widget: dryRunRadio},
		},
		OnSubmit: func() {
			// get variables
			namespace = namespaceSelectEntry.Text
			deploymentName = deploymentNameEntry.Text
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

			_ = ExecuteUninstallWorkflow(window, state)
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

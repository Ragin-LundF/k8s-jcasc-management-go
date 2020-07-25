package install

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/installactions"
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/events"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/validator"
	"time"
)

var namespaceSelectEntry *widget.SelectEntry
var namespaceErrorLabel = widget.NewLabel("")

// ScreenInstall shows the install screen
func ScreenInstall(window fyne.Window) fyne.CanvasObject {
	var namespace string
	var deploymentName string
	var installTypeOption string
	var dryRunOption string
	var secretsPasswords string

	// Namespace
	namespaceSelectEntry = uielements.CreateNamespaceSelectEntry(namespaceErrorLabel)

	// Deployment name
	deploymentNameEntry := uielements.CreateDeploymentNameEntry()

	// Install or update
	installTypeRadio := uielements.CreateInstallTypeRadio()

	// Dry-run or execute
	dryRunRadio := uielements.CreateDryRunRadio()

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
			if !validator.ValidateNamespaceAvailableInConfig(namespace) {
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
			state, err := installactions.CalculateDirectoriesForInstall(state, state.Namespace)
			if err != nil {
				dialog.ShowError(err, window)
			}

			// Check Jenkins directories
			state = installactions.CheckJenkinsDirectories(state)

			// ask for password
			if dryRunOption == constants.InstallDryRunInactive {
				openSecretsPasswordDialog(window, secretsPasswordEntry, state)
			} else {
				_ = ExecuteInstallWorkflow(window, state)
				// show output
				uielements.ShowLogOutput(window)
			}
		},
	}

	return widget.NewVBox(
		form,
	)
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

func init() {
	createNamespaceNotifier := namespaceCreatedNotifier{}
	events.NamespaceCreated.Register(createNamespaceNotifier)
}

type namespaceCreatedNotifier struct {
	namespace string
}

func (notifier namespaceCreatedNotifier) Handle(payload events.NamespaceCreatedPayload) {
	logger.Log().Info("[install_gui] -> Retrieved event to that new namespace was created")
	namespaceSelectEntry.SetOptions(namespaceactions.ActionReadNamespaceWithFilter(nil))

	events.RefreshTabs.Trigger(events.RefreshTabsPayload{
		Time: time.Now(),
	})
}

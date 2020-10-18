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
	var secretsFile string
	var secretsPasswords string

	// Entries
	var namespaceSelectEntry = uielements.CreateNamespaceSelectEntry(namespaceErrorLabel)
	var secretsFileSelect = uielements.CreateSecretsFileEntry()
	var deploymentNameEntry = uielements.CreateDeploymentNameEntry()
	var installTypeRadio = uielements.CreateInstallTypeRadio()
	var dryRunRadio = uielements.CreateDryRunRadio()
	var secretsPasswordEntry = widget.NewPasswordEntry()

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "Secrets file", Widget: secretsFileSelect},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "Deployment Name", Widget: deploymentNameEntry},
			{Text: "Installation type", Widget: installTypeRadio},
			{Text: "Execute or dry run", Widget: dryRunRadio},
		},
		OnSubmit: func() {
			// get variables
			secretsFile = secretsFileSelect.Selected
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
			var state = models.StateData{
				Namespace:       namespace,
				DeploymentName:  deploymentName,
				HelmCommand:     installTypeOption,
				SecretsPassword: &secretsPasswords,
				SecretsFileName: secretsFile,
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
		widget.NewLabel(""),
		form,
	)
}

// Secrets password dialog
func openSecretsPasswordDialog(window fyne.Window, secretsPasswordEntry *widget.Entry, state models.StateData) {
	var secretsPasswordWindow = widget.NewForm(widget.NewFormItem("Password", secretsPasswordEntry))
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
	var createNamespaceNotifier = namespaceCreatedNotifier{}
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

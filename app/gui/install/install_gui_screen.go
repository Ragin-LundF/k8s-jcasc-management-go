package install

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/events"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/validator"
	"time"
)

var namespaceErrorLabel = widget.NewLabel("")
var namespaceSelectEntry = widget.NewSelectEntry([]string{})

// ScreenInstall shows the install screen
func ScreenInstall(window fyne.Window) fyne.CanvasObject {
	var namespace string
	var deploymentName string
	var installTypeOption string
	var dryRunOption string
	var secretsFile string
	var secretsPasswords string

	// Entries
	namespaceSelectEntry = uielements.CreateNamespaceSelectEntry(namespaceErrorLabel)
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
				configuration.GetConfiguration().SetDryRun(true)
			} else {
				configuration.GetConfiguration().SetDryRun(false)
			}
			if !validator.ValidateNamespaceAvailableInConfig(namespace) {
				namespaceErrorLabel.SetText("Error: namespace is unknown!")
				namespaceErrorLabel.Show()
				return
			}

			// map state
			var projectConfig = install.NewInstallProjectConfig()
			var err = projectConfig.LoadProjectConfigIfExists(namespace)
			if err != nil {
				dialog.ShowError(err, window)
			}
			projectConfig.Project.Base.DeploymentName = deploymentName
			projectConfig.HelmCommand = installTypeOption
			projectConfig.SecretsPassword = &secretsPasswords
			projectConfig.SecretsFileName = secretsFile

			// ask for password
			if dryRunOption == constants.InstallDryRunInactive {
				openSecretsPasswordDialog(window, secretsPasswordEntry, projectConfig)
			} else {
				_ = ExecuteInstallWorkflow(window, projectConfig)
				// show output
				uielements.ShowLogOutput(window)
			}
		},
	}

	return container.NewVBox(
		widget.NewLabel(""),
		form,
	)
}

// Secrets password dialog
func openSecretsPasswordDialog(window fyne.Window, secretsPasswordEntry *widget.Entry, projectConfig install.ProjectConfig) {
	var secretsPasswordWindow = widget.NewForm(widget.NewFormItem("Password", secretsPasswordEntry))
	secretsPasswordWindow.Resize(fyne.Size{Width: 300, Height: 90})

	dialog.ShowCustomConfirm("Password for Secrets...", "Ok", "Cancel", secretsPasswordWindow, func(confirmed bool) {
		if confirmed {
			projectConfig.SecretsPassword = &secretsPasswordEntry.Text
			_ = ExecuteInstallWorkflow(window, projectConfig)
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
	if namespaceSelectEntry != nil {
		namespaceSelectEntry.SetOptions(namespaceactions.ActionReadNamespaceWithFilter(nil))
	}

	events.RefreshTabs.Trigger(events.RefreshTabsPayload{
		Time: time.Now(),
	})
}

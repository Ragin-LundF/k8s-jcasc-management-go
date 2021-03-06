package uninstall

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"k8s-management-go/app/actions/installactions"
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/events"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/validator"
	"time"
)

// Namespace
var namespaceErrorLabel = widget.NewLabel("")
var namespaceSelectEntry = widget.NewSelectEntry([]string{})

// ScreenUninstall shows the uninstall screen
func ScreenUninstall(window fyne.Window) fyne.CanvasObject {
	var namespace string
	var deploymentName string
	var installTypeOption string
	var dryRunOption string
	var secretsPasswords string
	namespaceSelectEntry = uielements.CreateNamespaceSelectEntry(namespaceErrorLabel)

	// Deployment name
	var deploymentNameEntry = uielements.CreateDeploymentNameEntry()

	// Dry-run or execute
	var dryRunRadio = uielements.CreateDryRunRadio()

	var form = &widget.Form{
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
			var state = models.StateData{
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

			_ = ExecuteUninstallWorkflow(window, state)
			// show output
			uielements.ShowLogOutput(window)
		},
	}

	return container.NewVBox(
		widget.NewLabel(""),
		form,
	)
}

func init() {
	var createNamespaceNotifier = namespaceCreatedNotifier{}
	events.NamespaceCreated.Register(createNamespaceNotifier)
}

type namespaceCreatedNotifier struct {
	namespace string
}

func (notifier namespaceCreatedNotifier) Handle(payload events.NamespaceCreatedPayload) {
	logger.Log().Info("[uninstall_gui] -> Retrieved event to that new namespace was created")
	if namespaceSelectEntry != nil {
		namespaceSelectEntry.SetOptions(namespaceactions.ActionReadNamespaceWithFilter(nil))
	}

	events.RefreshTabs.Trigger(events.RefreshTabsPayload{
		Time: time.Now(),
	})
}

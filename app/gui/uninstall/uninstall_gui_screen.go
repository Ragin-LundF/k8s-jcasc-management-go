package uninstall

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
)

// Namespace
var namespaceErrorLabel = widget.NewLabel("")
var namespaceSelectEntry *widget.SelectEntry

// ScreenUninstall shows the uninstall screen
func ScreenUninstall(window fyne.Window) fyne.CanvasObject {
	var namespace string
	var deploymentName string
	var installTypeOption string
	var dryRunOption string
	var secretsPasswords string

	// Namespace
	namespaceSelectEntry = uielements.CreateNamespaceSelectEntry(namespaceErrorLabel)

	// Deployment name
	deploymentNameEntry := uielements.CreateDeploymentNameEntry()

	// Dry-run or execute
	dryRunRadio := uielements.CreateDryRunRadio()

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

			_ = ExecuteUninstallWorkflow(window, state)
			// show output
			uielements.ShowLogOutput(window)
		},
	}

	box := widget.NewVBox(
		form,
	)

	return box
}

func init() {
	createNamespaceNotifier := namespaceCreatedNotifier{}
	events.NamespaceCreated.Register(createNamespaceNotifier)
}

type namespaceCreatedNotifier struct {
	namespace string
}

func (notifier namespaceCreatedNotifier) updateView() {
	logger.Log().Info("[uninstall_gui] -> Retrieved event to that new namespace was created")
	namespaceSelectEntry.SetOptions(namespaceactions.ActionReadNamespaceWithFilter(nil))
	namespaceSelectEntry.Refresh()
}

func (notifier namespaceCreatedNotifier) Handle(payload events.NamespaceCreatedPayload) {
	notifier.updateView()
}

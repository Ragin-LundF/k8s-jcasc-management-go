package namespace

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/events"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/utils/logger"
	"time"
)

var namespaceErrorLabel = widget.NewLabel("")
var namespaceSelectEntry = widget.NewSelectEntry([]string{})

// ScreenNamespaceCreate shows the create namespace screen
func ScreenNamespaceCreate(window fyne.Window) fyne.CanvasObject {
	var namespace string
	namespaceSelectEntry = uielements.CreateNamespaceSelectEntry(namespaceErrorLabel)

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "", Widget: namespaceErrorLabel},
		},
		OnSubmit: func() {
			// get variables
			namespace = namespaceSelectEntry.Text

			// map state
			var projectConfig = install.NewInstallProjectConfig()
			projectConfig.Project.SetNamespace(namespace)

			_ = ExecuteCreateNamespaceWorkflow(window, projectConfig)
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
	logger.Log().Info("[namespace_gui] -> Retrieved event to that new namespace was created")
	if namespaceSelectEntry != nil {
		namespaceSelectEntry.SetOptions(namespaceactions.ActionReadNamespaceWithFilter(nil))
	}

	events.RefreshTabs.Trigger(events.RefreshTabsPayload{
		Time: time.Now(),
	})
}

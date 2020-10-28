package namespace

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/events"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"time"
)

var namespaceErrorLabel = widget.NewLabel("")
var namespaceSelectEntry = uielements.CreateNamespaceSelectEntry(namespaceErrorLabel)

// ScreenNamespaceCreate shows the create namespace screen
func ScreenNamespaceCreate(window fyne.Window) fyne.CanvasObject {
	var namespace string

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "", Widget: namespaceErrorLabel},
		},
		OnSubmit: func() {
			// get variables
			namespace = namespaceSelectEntry.Text

			// map state
			var state = models.StateData{
				Namespace: namespace,
			}

			_ = ExecuteCreateNamespaceWorkflow(window, state)
			// show output
			uielements.ShowLogOutput(window)
		},
	}

	return widget.NewVBox(
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
	namespaceSelectEntry.SetOptions(namespaceactions.ActionReadNamespaceWithFilter(nil))

	events.RefreshTabs.Trigger(events.RefreshTabsPayload{
		Time: time.Now(),
	})
}

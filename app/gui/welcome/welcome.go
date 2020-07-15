package welcome

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/kubernetesactions"
	"k8s-management-go/app/gui/resources"
	"k8s-management-go/app/gui/uiconstants"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/utils/loggingstate"
)

// ScreenWelcome shows the welcome screen
func ScreenWelcome(window fyne.Window, info string) fyne.CanvasObject {
	// set label
	var labelInfo *widget.Label

	logo := canvas.NewImageFromResource(resources.K8sJcascMgmtIcon())
	logo.SetMinSize(fyne.NewSize(128, 128))

	if info == "" {
		labelInfo = widget.NewLabelWithStyle("You are on the latest version.", fyne.TextAlignCenter, fyne.TextStyle{
			Italic: true,
		})
	} else {
		labelInfo = widget.NewLabelWithStyle(info, fyne.TextAlignCenter, fyne.TextStyle{
			Bold: true,
		})
	}

	// Context switcher
	var k8sCurrentContext = widget.NewLabel(kubernetesactions.GetKubernetesConfig().CurrentContext())
	var k8sContextErrorLabel = widget.NewLabel("")
	var k8sContextSwitchSelectEntry = uielements.CreateKubernetesContextSelectEntry(k8sContextErrorLabel)
	// clear logging state to reset the log entries till this point
	loggingstate.ClearLoggingState()

	// create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Current Kubernetes Config", Widget: k8sCurrentContext},
			{Text: "Kubernetes Context", Widget: k8sContextSwitchSelectEntry},
			{Text: "", Widget: k8sContextErrorLabel},
		},
		OnSubmit: func() {
			if err := kubernetesactions.SwitchKubernetesConfig(k8sContextSwitchSelectEntry.Text); err == nil {
				k8sCurrentContext.SetText(kubernetesactions.GetKubernetesConfig().CurrentContext())
				window.SetTitle(uiconstants.K8sJcasCMgmtTitle + kubernetesactions.GetKubernetesConfig().CurrentContext())
				k8sContextSwitchSelectEntry.SetText("")
			} else {
				k8sContextErrorLabel.SetText("Unable to switch context...")
			}

			uielements.ShowLogOutput(window)
		},
	}

	return widget.NewVBox(
		layout.NewSpacer(),
		labelInfo,
		widget.NewLabelWithStyle("Welcome to Kubernetes Jenkins Configuration as Code", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		form,
		layout.NewSpacer(),
	)
}

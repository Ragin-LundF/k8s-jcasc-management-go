package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/cmd/fyne_demo/screens"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/k8s_screens"
)

func CreateTabMenu(k8sJcascApp fyne.App, k8sJcascWindow fyne.Window) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), k8s_screens.WelcomeScreen()),
		widget.NewTabItemWithIcon("Deployments", theme.ConfirmIcon(), InstallScreen(k8sJcascWindow, k8sJcascApp.Preferences())),
		widget.NewTabItemWithIcon("Secrets", theme.ContentAddIcon(), SecretsScreen(k8sJcascApp.Preferences())),
		widget.NewTabItemWithIcon("Create Project", theme.ViewRestoreIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Create Deployment-Only Project", theme.ViewRestoreIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Create Namespace", theme.ViewRestoreIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Create Jenkins User Password", theme.ViewFullScreenIcon(), screens.DialogScreen(k8sJcascWindow)))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(k8sJcascApp.Preferences().Int(PreferencesMenuMainTab))

	return tabs
}
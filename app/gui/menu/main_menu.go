package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/cmd/fyne_demo/screens"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	namespace2 "k8s-management-go/app/gui/jenkinsuser"
	"k8s-management-go/app/gui/namespace"
	"k8s-management-go/app/gui/welcome"
)

func CreateTabMenu(k8sJcascApp fyne.App, k8sJcascWindow fyne.Window) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), welcome.ScreenWelcome()),
		widget.NewTabItemWithIcon("Deployments", theme.ConfirmIcon(), InstallScreen(k8sJcascWindow, k8sJcascApp.Preferences())),
		widget.NewTabItemWithIcon("Secrets", theme.ContentAddIcon(), SecretsScreen(k8sJcascApp.Preferences())),
		widget.NewTabItemWithIcon("Create Project", theme.ViewRestoreIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Create Deployment-Only Project", theme.ViewRestoreIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Create Namespace", theme.ViewRestoreIcon(), namespace.ScreenNamespaceCreate(k8sJcascWindow)),
		widget.NewTabItemWithIcon("Create Jenkins User Password", theme.ViewFullScreenIcon(), namespace2.ScreenJenkinsUserPasswordCreate(k8sJcascWindow)))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(k8sJcascApp.Preferences().Int(PreferencesMenuMainTab))

	return tabs
}

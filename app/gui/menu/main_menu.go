package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/welcome"
)

func CreateTabMenu(k8sJcascApp fyne.App, k8sJcascWindow fyne.Window) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), welcome.ScreenWelcome()),
		widget.NewTabItemWithIcon("Deployments", theme.ConfirmIcon(), InstallScreen(k8sJcascWindow, k8sJcascApp.Preferences())),
		widget.NewTabItemWithIcon("Secrets", theme.WarningIcon(), SecretsScreen(k8sJcascWindow, k8sJcascApp.Preferences())),
		widget.NewTabItemWithIcon("Create Project", theme.DocumentCreateIcon(), ProjectsScreen(k8sJcascWindow, k8sJcascApp.Preferences())),
		widget.NewTabItemWithIcon("Tools", theme.SettingsIcon(), ToolsScreen(k8sJcascWindow, k8sJcascApp.Preferences())))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(k8sJcascApp.Preferences().Int(PreferencesMenuMainTab))

	return tabs
}

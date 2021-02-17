package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"k8s-management-go/app/gui/welcome"
)

// CreateTabMenu creates the main tab menu
func CreateTabMenu(k8sJcascApp fyne.App, k8sJcascWindow fyne.Window, info string) (tabs *container.AppTabs) {
	tabs = container.NewAppTabs(
		container.NewTabItemWithIcon("Welcome", theme.HomeIcon(), welcome.ScreenWelcome(k8sJcascWindow, info)),
		container.NewTabItemWithIcon("Deployments", theme.ConfirmIcon(), InstallScreen(k8sJcascWindow, k8sJcascApp.Preferences())),
		container.NewTabItemWithIcon("Secrets", theme.WarningIcon(), SecretsScreen(k8sJcascWindow, k8sJcascApp.Preferences())),
		container.NewTabItemWithIcon("Create Project", theme.DocumentCreateIcon(), ProjectsScreen(k8sJcascWindow, k8sJcascApp.Preferences())),
		container.NewTabItemWithIcon("Tools", theme.SettingsIcon(), ToolsScreen(k8sJcascWindow, k8sJcascApp.Preferences())))
	tabs.SetTabLocation(container.TabLocationTop)

	return tabs
}

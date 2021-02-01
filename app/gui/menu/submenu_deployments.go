package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"k8s-management-go/app/gui/install"
	"k8s-management-go/app/gui/uiconstants"
	"k8s-management-go/app/gui/uninstall"
)

// InstallScreen shows the installation screen
func InstallScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return container.NewVBox(
		installSubMenu(window, preferences),
		layout.NewSpacer())
}

func installSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *container.AppTabs) {
	tabs = container.NewAppTabs(
		container.NewTabItemWithIcon("Install", theme.MediaPlayIcon(), install.ScreenInstall(window)),
		container.NewTabItemWithIcon("Uninstall", theme.MediaPauseIcon(), uninstall.ScreenUninstall(window)))

	tabs.SetTabLocation(container.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(uiconstants.PreferencesSubMenuDeploymentsTab))

	return tabs
}

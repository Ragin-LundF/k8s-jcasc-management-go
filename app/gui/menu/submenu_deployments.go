package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/install"
	"k8s-management-go/app/gui/uiconstants"
	"k8s-management-go/app/gui/uninstall"
)

// InstallScreen shows the installation screen
func InstallScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		installSubMenu(window, preferences),
		layout.NewSpacer())
}

func installSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Install", theme.MediaPlayIcon(), install.ScreenInstall(window)),
		widget.NewTabItemWithIcon("Uninstall", theme.MediaPauseIcon(), uninstall.ScreenUninstall(window)))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(uiconstants.PreferencesSubMenuDeploymentsTab))

	return tabs
}

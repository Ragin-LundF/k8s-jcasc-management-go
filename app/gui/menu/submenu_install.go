package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/install"
	"k8s-management-go/app/gui/uninstall"
)

func InstallScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		installSubMenu(window, preferences),
	)
}

func installSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Install", theme.ConfirmIcon(), install.ScreenInstall(window)),
		widget.NewTabItemWithIcon("Uninstall", theme.CancelIcon(), uninstall.ScreenUninstall(window)))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(PreferencesSubMenuInstallTab))

	return tabs
}

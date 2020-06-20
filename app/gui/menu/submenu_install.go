package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/cmd/fyne_demo/screens"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/install"
)

func InstallScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		installSubMenu(window, preferences),
	)
}

func installSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Install", theme.ConfirmIcon(), install.ScreenInstall(window)),
		widget.NewTabItemWithIcon("Uninstall", theme.CancelIcon(), screens.WidgetScreen()),
		widget.NewTabItemWithIcon("Upgrade", theme.ViewRestoreIcon(), screens.ContainerScreen()))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(PreferencesSubMenuInstallTab))

	return tabs
}

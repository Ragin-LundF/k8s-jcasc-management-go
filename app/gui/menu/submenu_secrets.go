package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/cmd/fyne_demo/screens"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func SecretsScreen(preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		secretsSubMenu(preferences),
	)
}

func secretsSubMenu(preferences fyne.Preferences) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Apply Secrets", theme.ContentAddIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Apply Secrets to all Namespaces", theme.ConfirmIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Encrypt Secrets", theme.ContentCutIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Decrypt Secrets", theme.ContentRedoIcon(), screens.ContainerScreen()))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(PreferencesSubMenuSecretsTab))

	return tabs
}

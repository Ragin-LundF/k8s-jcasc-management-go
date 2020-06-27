package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/cmd/fyne_demo/screens"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func SecretsScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		secretsSubMenu(preferences),
		layout.NewSpacer())
}

func secretsSubMenu(preferences fyne.Preferences) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Apply Secrets", theme.MailReplyIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Apply Secrets to all Namespaces", theme.MailReplyAllIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Encrypt Secrets", theme.VisibilityOffIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Decrypt Secrets", theme.VisibilityIcon(), screens.ContainerScreen()))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(PreferencesSubMenuSecretsTab))

	return tabs
}

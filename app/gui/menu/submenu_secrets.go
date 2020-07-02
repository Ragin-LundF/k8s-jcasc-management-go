package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/secrets"
)

// SecretsScreen shows the secrets screen
func SecretsScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		secretsSubMenu(window, preferences),
		layout.NewSpacer())
}

func secretsSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Apply Secrets", theme.MailReplyIcon(), secrets.ScreenApplySecretsToNamespace(window)),
		widget.NewTabItemWithIcon("Apply Secrets to all Namespaces", theme.MailReplyAllIcon(), secrets.ScreenApplySecretsToAllNamespace(window)),
		widget.NewTabItemWithIcon("Encrypt Secrets", theme.VisibilityOffIcon(), secrets.ScreenEncryptSecrets(window)),
		widget.NewTabItemWithIcon("Decrypt Secrets", theme.VisibilityIcon(), secrets.ScreenDecryptSecrets(window)))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(PreferencesSubMenuSecretsTab))

	return tabs
}

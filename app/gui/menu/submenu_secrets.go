package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"k8s-management-go/app/gui/secrets"
	"k8s-management-go/app/gui/uiconstants"
)

// SecretsScreen shows the secrets screen
func SecretsScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return container.NewVBox(
		secretsSubMenu(window, preferences),
		layout.NewSpacer())
}

func secretsSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *container.AppTabs) {
	tabs = container.NewAppTabs(
		container.NewTabItemWithIcon("Apply Secrets", theme.MailReplyIcon(), secrets.ScreenApplySecretsToNamespace(window)),
		container.NewTabItemWithIcon("Apply Secrets to all Namespaces", theme.MailReplyAllIcon(), secrets.ScreenApplySecretsToAllNamespace(window)),
		container.NewTabItemWithIcon("Encrypt Secrets", theme.VisibilityOffIcon(), secrets.ScreenEncryptSecrets(window)),
		container.NewTabItemWithIcon("Decrypt Secrets", theme.VisibilityIcon(), secrets.ScreenDecryptSecrets(window)))

	tabs.SetTabLocation(container.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(uiconstants.PreferencesSubMenuSecretsTab))

	return tabs
}

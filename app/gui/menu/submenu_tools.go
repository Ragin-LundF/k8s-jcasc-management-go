package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"k8s-management-go/app/gui/jenkinsuser"
	"k8s-management-go/app/gui/namespace"
	"k8s-management-go/app/gui/uiconstants"
)

// ToolsScreen shows the tools scrren
func ToolsScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return container.NewVBox(
		toolsSubMenu(window, preferences),
		layout.NewSpacer())
}

func toolsSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *container.AppTabs) {
	tabs = container.NewAppTabs(
		container.NewTabItemWithIcon("Create Namespace", theme.ContentPasteIcon(), namespace.ScreenNamespaceCreate(window)),
		container.NewTabItemWithIcon("Create Jenkins User Password", theme.ConfirmIcon(), jenkinsuser.ScreenJenkinsUserPasswordCreate(window)))

	tabs.SetTabLocation(container.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(uiconstants.PreferencesSubMenuToolsTab))

	return tabs
}

package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/jenkinsuser"
	"k8s-management-go/app/gui/namespace"
)

// ToolsScreen shows the tools scrren
func ToolsScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		toolsSubMenu(window, preferences),
		layout.NewSpacer())
}

func toolsSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Create Namespace", theme.ContentPasteIcon(), namespace.ScreenNamespaceCreate(window)),
		widget.NewTabItemWithIcon("Create Jenkins User Password", theme.ConfirmIcon(), jenkinsuser.ScreenJenkinsUserPasswordCreate(window)))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(PreferencesSubMenuToolsTab))

	return tabs
}

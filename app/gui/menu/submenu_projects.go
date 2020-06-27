package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/jenkinsuser"
	"k8s-management-go/app/gui/namespace"
)

func ProjectsScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		projectsSubMenu(window, preferences),
		layout.NewSpacer())
}

func projectsSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Create Project", theme.MediaRecordIcon(), namespace.ScreenNamespaceCreate(window)),
		widget.NewTabItemWithIcon("Create Deployment-Only Project", theme.MediaReplayIcon(), jenkinsuser.ScreenJenkinsUserPasswordCreate(window)))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(PreferencesSubMenuProjectsTab))

	return tabs
}

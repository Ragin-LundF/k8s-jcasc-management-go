package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/createproject"
)

// ProjectsScreen shows the projects screen
func ProjectsScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		projectsSubMenu(window, preferences),
		layout.NewSpacer())
}

func projectsSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Create Project", theme.MediaRecordIcon(), createproject.ScreenCreateFullProject(window)),
		widget.NewTabItemWithIcon("Create Deployment-Only Project", theme.MediaReplayIcon(), createproject.ScreenCreateDeployOnlyProject(window)))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(PreferencesSubMenuProjectsTab))

	return tabs
}

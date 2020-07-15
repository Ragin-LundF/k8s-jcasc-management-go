package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/createproject"
	"k8s-management-go/app/gui/uiconstants"
)

// ProjectsScreen shows the projects screen
func ProjectsScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return widget.NewVBox(
		projectsSubMenu(window, preferences),
		layout.NewSpacer())
}

func projectsSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *widget.TabContainer) {
	// form create full project
	formScreenCreateFullProject := createproject.ScreenCreateFullProject(window)
	boxScreenCreateFullProject := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		formScreenCreateFullProject,
	)
	// form create deploy only project
	formScreenCreateDeployOnlyProject := createproject.ScreenCreateDeployOnlyProject(window)
	boxScreenCreateDeployOnlyProject := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		formScreenCreateDeployOnlyProject,
	)

	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Create Project", theme.MediaRecordIcon(), boxScreenCreateFullProject),
		widget.NewTabItemWithIcon("Create Deployment-Only Project", theme.MediaReplayIcon(), boxScreenCreateDeployOnlyProject))

	tabs.SetTabLocation(widget.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(uiconstants.PreferencesSubMenuProjectsTab))

	return tabs
}

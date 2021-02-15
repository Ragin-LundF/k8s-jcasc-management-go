package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"k8s-management-go/app/gui/createproject"
	"k8s-management-go/app/gui/uiconstants"
)

// ProjectsScreen shows the projects screen
func ProjectsScreen(window fyne.Window, preferences fyne.Preferences) fyne.CanvasObject {
	return container.NewVBox(
		projectsSubMenu(window, preferences),
		layout.NewSpacer())
}

func projectsSubMenu(window fyne.Window, preferences fyne.Preferences) (tabs *container.AppTabs) {
	// form create full project
	var formScreenCreateFullProject = createproject.ScreenCreateFullProject(window)
	var boxScreenCreateFullProject = container.NewVBox(
		widget.NewLabel(""),
		container.NewHBox(layout.NewSpacer()),
		formScreenCreateFullProject,
	)
	// form create deploy only project
	var formScreenCreateDeployOnlyProject = createproject.ScreenCreateDeployOnlyProject(window)
	var boxScreenCreateDeployOnlyProject = container.NewVBox(
		widget.NewLabel(""),
		container.NewHBox(layout.NewSpacer()),
		formScreenCreateDeployOnlyProject,
	)

	tabs = container.NewAppTabs(
		container.NewTabItemWithIcon("Create Project", theme.MediaRecordIcon(), boxScreenCreateFullProject),
		container.NewTabItemWithIcon("Create Deployment-Only Project", theme.MediaReplayIcon(), boxScreenCreateDeployOnlyProject))

	tabs.SetTabLocation(container.TabLocationTop)
	tabs.SelectTabIndex(preferences.Int(uiconstants.PreferencesSubMenuProjectsTab))

	return tabs
}

// +build !ignore

package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"k8s-management-go/app/cli"
	"k8s-management-go/app/gui/menu"
	"k8s-management-go/app/gui/resources"
)

// StartApp will start app with GUI
func StartApp(info string) {
	k8sJcascApp := app.NewWithID("k8s_jcasc_mgmt_go_ui")
	k8sJcascApp.SetIcon(resources.K8sJcascMgmtIcon())

	// set theme
	setTheme(k8sJcascApp)

	k8sJcascWindow := k8sJcascApp.NewWindow("K8S JCasC Management")
	k8sJcascWindow.SetIcon(resources.K8sJcascMgmtIcon())
	mainMenu := menu.CreateMainMenu(k8sJcascApp, k8sJcascWindow)

	k8sJcascWindow.SetMainMenu(mainMenu)
	k8sJcascWindow.SetMaster()

	tabs := menu.CreateTabMenu(k8sJcascApp, k8sJcascWindow, info)

	k8sJcascWindow.SetContent(tabs)
	k8sJcascWindow.Resize(fyne.Size{
		Width:  900,
		Height: 400,
	})
	k8sJcascWindow.ShowAndRun()
}

// StartCli will start App with CLI
func StartCli(info string) {
	cli.Workflow(info, nil)
}

func setTheme(app fyne.App) {
	if app.Preferences().String(menu.PreferencesTheme) == menu.PreferencesThemeLight {
		app.Settings().SetTheme(theme.LightTheme())
	} else {
		app.Settings().SetTheme(theme.DarkTheme())
	}
}

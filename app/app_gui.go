// +build darwin

package app

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"k8s-management-go/app/cli"
	"k8s-management-go/app/gui/menu"
)

// start app with GUI
func StartApp(info string) {
	k8sJcascApp := app.NewWithID("k8s_jcasc_mgmt_go_ui")
	k8sJcascApp.SetIcon(theme.FyneLogo())

	k8sJcascWindow := k8sJcascApp.NewWindow("K8S JCasC Management")
	mainMenu := menu.CreateMainMenu(k8sJcascApp)

	k8sJcascWindow.SetMainMenu(mainMenu)
	k8sJcascWindow.SetMaster()

	tabs := menu.CreateTabMenu(k8sJcascApp, k8sJcascWindow)

	k8sJcascWindow.SetContent(tabs)
	k8sJcascWindow.ShowAndRun()
	k8sJcascApp.Preferences().SetInt(menu.PreferencesMenuMainTab, tabs.CurrentTabIndex())
}

// Start App with CLI
func StartCli(info string) {
	cli.Workflow(info, nil)
}

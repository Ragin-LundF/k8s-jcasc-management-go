// +build !cli

package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"k8s-management-go/app/actions/kubernetesactions"
	"k8s-management-go/app/cli"
	"k8s-management-go/app/events"
	"k8s-management-go/app/gui/menu"
	"k8s-management-go/app/gui/resources"
	"k8s-management-go/app/gui/uiconstants"
	"k8s-management-go/app/utils/logger"
)

var tabs *container.AppTabs

// StartApp will start app with GUI
func StartApp(info string) {
	var k8sJcascApp = app.NewWithID("k8s_jcasc_mgmt_go_ui")
	k8sJcascApp.SetIcon(resources.K8sJcascMgmtIcon())

	// set theme
	setTheme(k8sJcascApp)

	var k8sJcascWindow = k8sJcascApp.NewWindow(uiconstants.K8sJcasCMgmtTitle + kubernetesactions.GetKubernetesConfig().CurrentContext())
	k8sJcascWindow.SetIcon(resources.K8sJcascMgmtIcon())
	var mainMenu = menu.CreateMainMenu(k8sJcascApp, k8sJcascWindow)

	k8sJcascWindow.SetMainMenu(mainMenu)
	k8sJcascWindow.SetMaster()

	tabs = menu.CreateTabMenu(k8sJcascApp, k8sJcascWindow, info)

	k8sJcascWindow.SetContent(tabs)
	k8sJcascWindow.Resize(fyne.Size{
		Width:  980,
		Height: 400,
	})
	k8sJcascWindow.ShowAndRun()
}

// StartCli will start App with CLI
func StartCli(info string) {
	cli.Workflow(info, nil)
}

func setTheme(app fyne.App) {
	if app.Preferences().String(uiconstants.PreferencesTheme) == uiconstants.PreferencesThemeLight {
		app.Settings().SetTheme(theme.LightTheme())
	} else {
		app.Settings().SetTheme(theme.DarkTheme())
	}
}

func init() {
	// register as finalizer
	var notifierRefreshTabs = tabsRefreshNotifier{}
	events.RefreshTabs.Register(notifierRefreshTabs)
}

type tabsRefreshNotifier struct {
}

func (notifier tabsRefreshNotifier) Handle(payload events.RefreshTabsPayload) {
	logger.Log().Info("[app] -> Retrieved event to refresh tabs")
	tabs.Refresh()
}

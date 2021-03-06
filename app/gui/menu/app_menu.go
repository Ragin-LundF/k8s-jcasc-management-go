package menu

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"k8s-management-go/app/actions/migration"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/gui/uiconstants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
)

// CreateMainMenu creates the main menu
func CreateMainMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {
	// K8S Management Menu
	var settingsItem = fyne.NewMenuItem("Configuration", func() { printConfiguration(window) })
	var toolsItem = fyne.NewMenuItem("Migrate templates v2 -> v3", func() { migrateTemplatesV2(window) })
	var quitItem = fyne.NewMenuItem("Quit", func() { app.Quit() })
	var darkThemeItem = fyne.NewMenuItem("Dark Theme", func() {
		app.Settings().SetTheme(theme.DarkTheme())
		app.Preferences().SetString(uiconstants.PreferencesTheme, uiconstants.PreferencesThemeDark)
	})
	var lightThemeItem = fyne.NewMenuItem("Light Theme", func() {
		app.Settings().SetTheme(theme.LightTheme())
		app.Preferences().SetString(uiconstants.PreferencesTheme, uiconstants.PreferencesThemeLight)
	})

	var mainMenu = fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		fyne.NewMenu("K8S Management", fyne.NewMenuItemSeparator(), settingsItem, fyne.NewMenuItemSeparator(), quitItem),
		fyne.NewMenu("Theme", darkThemeItem, lightThemeItem),
		fyne.NewMenu("Tools", toolsItem),
	)

	return mainMenu
}

func migrateTemplatesV2(window fyne.Window) {
	var progressBar = widget.NewProgressBarInfinite()
	var progressBox = container.NewVBox(progressBar)
	var migrationDialog = dialog.NewCustom("Migration", "Please wait...", progressBox, window)
	progressBar.Show()
	progressBar.Start()
	migrationDialog.Show()
	var information = migration.MigrateTemplatesToV3()
	migrationDialog.Hide()
	progressBar.Stop()
	progressBar.Hide()

	dialog.ShowInformation("Migration", information, window)
}

func printConfiguration(window fyne.Window) {
	// System config
	var configSystemAsJSON, _ = json.MarshalIndent(configuration.GetConfiguration(), "", "\t")

	// textgrid for system config
	var textGridSystemConfig = widget.NewTextGrid()
	textGridSystemConfig.SetText(string(configSystemAsJSON))

	// IP config
	var configIP = models.GetIPConfiguration()
	var configIPAsJSON, _ = json.MarshalIndent(configIP, "", "\t")

	// writing into log
	var log = logger.Log()
	log.Info("---- Printing system configuration start -----")
	log.Info("\n" + string(configSystemAsJSON))
	log.Info("---- Printing system configuration end   -----")
	log.Info("---- Printing IP configuration start -----")
	log.Info("\n" + string(configIPAsJSON))
	log.Info("---- Printing IP configuration end   -----")

	dialog.ShowInformation("Configuration", "Your configuration was saved into your logs!", window)
}

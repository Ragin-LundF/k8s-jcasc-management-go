package menu

import (
	"encoding/json"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
)

// CreateMainMenu creates the main menu
func CreateMainMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {
	// K8S Management Menu
	settingsItem := fyne.NewMenuItem("Configuration", func() { printConfiguration(window) })
	quitItem := fyne.NewMenuItem("Quit", func() { app.Quit() })

	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		fyne.NewMenu("K8S Management", fyne.NewMenuItemSeparator(), settingsItem, fyne.NewMenuItemSeparator(), quitItem),
	)

	return mainMenu
}

func printConfiguration(window fyne.Window) {
	// System config
	configSystem := models.GetConfiguration()
	configSystemAsJSON, _ := json.MarshalIndent(configSystem, "", "\t")

	// textgrid for system config
	textGridSystemConfig := widget.NewTextGrid()
	textGridSystemConfig.SetText(string(configSystemAsJSON))

	// IP config
	configIP := models.GetIPConfiguration()
	configIPAsJSON, _ := json.MarshalIndent(configIP, "", "\t")

	// writing into log
	log := logger.Log()
	log.Info("---- Printing system configuration start -----")
	log.Info("\n" + string(configSystemAsJSON))
	log.Info("---- Printing system configuration end   -----")
	log.Info("---- Printing IP configuration start -----")
	log.Info("\n" + string(configIPAsJSON))
	log.Info("---- Printing IP configuration end   -----")

	dialog.ShowInformation("Configuration", "Your configuration was saved into your logs!", window)
}

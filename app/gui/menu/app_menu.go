package menu

import (
	"fmt"
	"fyne.io/fyne"
)

func CreateMainMenu(app fyne.App) *fyne.MainMenu {
	// K8S Management Menu
	settingsItem := fyne.NewMenuItem("Configuration", func() { fmt.Println("Menu Settings") })
	quitItem := fyne.NewMenuItem("Quit", func() { app.Quit() })

	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		fyne.NewMenu("K8S Management", fyne.NewMenuItemSeparator(), settingsItem, fyne.NewMenuItemSeparator(), quitItem),
	)

	return mainMenu
}

package menu

import (
	"fmt"
	"fyne.io/fyne"
)

func CreateMainMenu(app fyne.App) *fyne.MainMenu {
	// K8S Management Menu
	newItem := createNewSegment()
	settingsItem := fyne.NewMenuItem("Configuration", func() { fmt.Println("Menu Settings") })
	quitItem := fyne.NewMenuItem("Quit", func() { app.Quit() })

	findItem := fyne.NewMenuItem("Find", func() { fmt.Println("Menu Find") })
	helpMenu := fyne.NewMenu("Help", fyne.NewMenuItem("Help", func() { fmt.Println("Help Menu") }))

	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		fyne.NewMenu("K8S Management", newItem, fyne.NewMenuItemSeparator(), settingsItem, fyne.NewMenuItemSeparator(), quitItem),
		fyne.NewMenu("Edit", fyne.NewMenuItemSeparator(), findItem),
		helpMenu,
	)

	return mainMenu
}

func createNewSegment() *fyne.MenuItem {
	newItem := fyne.NewMenuItem("New Project", nil)
	newItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Project with Jenkins", func() { fmt.Println("Menu New->Other->Project") }),
		fyne.NewMenuItem("Project for deployment", func() { fmt.Println("Menu New->Other->Mail") }),
	)

	return newItem
}

package menu

import (
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMainMenu(t *testing.T) {
	app := test.NewApp()
	window := app.NewWindow("test")
	mainMenu := CreateMainMenu(app, window)

	assert.Equal(t, "K8S Management", mainMenu.Items[0].Label)
	assert.Equal(t, "", mainMenu.Items[0].Items[0].Label) // separator
	assert.Equal(t, "Configuration", mainMenu.Items[0].Items[1].Label)
	assert.Equal(t, "", mainMenu.Items[0].Items[2].Label) // separator
	assert.Equal(t, "Quit", mainMenu.Items[0].Items[3].Label)

	assert.Equal(t, "Theme", mainMenu.Items[1].Label)
	assert.Equal(t, "Dark Theme", mainMenu.Items[1].Items[0].Label)
	assert.Equal(t, "Light Theme", mainMenu.Items[1].Items[1].Label)

	assert.Equal(t, "Tools", mainMenu.Items[2].Label)
	assert.Equal(t, "Migrate config v2 -> v3", mainMenu.Items[2].Items[0].Label)
	assert.Equal(t, "Migrate templates v2 -> v3", mainMenu.Items[2].Items[1].Label)

}

package menu

import (
	"fyne.io/fyne/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTabMenu(t *testing.T) {
	app := test.NewApp()
	window := app.NewWindow("test")
	tabMenu := CreateTabMenu(app, window, "")

	assert.Equal(t, "Welcome", tabMenu.Items[0].Text)
	assert.Equal(t, "Deployments", tabMenu.Items[1].Text)
	assert.Equal(t, "Secrets", tabMenu.Items[2].Text)
	assert.Equal(t, "Create Project", tabMenu.Items[3].Text)
	assert.Equal(t, "Tools", tabMenu.Items[4].Text)
}

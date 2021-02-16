package createproject

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScreenCreateFullProject(t *testing.T) {
	var deployOnlyPrjForm = ScreenCreateFullProject(test.NewApp().NewWindow("test"))
	var i = 0

	assert.Len(t, deployOnlyPrjForm.Items, 13)
	assert.Equal(t, "Namespace", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "IP address", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "Domain name", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "Jenkins system message", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "Jenkins Jobs Repository", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "Jenkins Existing PVC", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "Cloud Templates", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &container.Scroll{}, deployOnlyPrjForm.Items[i].Widget)
}

func TestScreenCreateDeployOnlyProject(t *testing.T) {
	var deployOnlyPrjForm = ScreenCreateDeployOnlyProject(test.NewApp().NewWindow("test"))
	var i = 0

	assert.Len(t, deployOnlyPrjForm.Items, 6)
	assert.Equal(t, "Namespace", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "IP address", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "Domain name", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[i].Widget)
	i++
	assert.Equal(t, "", deployOnlyPrjForm.Items[i].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[i].Widget)
}

func TestCheckCloudboxes(t *testing.T) {
	var selectLabelB = "label B"
	var checkboxes []*widget.Check
	var checkboxA = widget.NewCheck("label A", func(b bool) {})
	var checkboxB = widget.NewCheck(selectLabelB, func(b bool) {})
	var checkboxC = widget.NewCheck("label C", func(b bool) {})

	// set B as checked
	checkboxB.Checked = true

	// append boxes to array
	checkboxes = append(checkboxes, checkboxA, checkboxB, checkboxC)

	// calculate result
	var stringArr = checkCloudboxes(checkboxes)

	assert.Len(t, stringArr, 1)
	assert.Equal(t, selectLabelB, stringArr[0])
}

func TestCheckCloudboxesWithoutCheck(t *testing.T) {
	var selectLabelB = "label B"
	var checkboxes []*widget.Check
	var checkboxA = widget.NewCheck("label A", func(b bool) {})
	var checkboxB = widget.NewCheck(selectLabelB, func(b bool) {})
	var checkboxC = widget.NewCheck("label C", func(b bool) {})

	// append boxes to array
	checkboxes = append(checkboxes, checkboxA, checkboxB, checkboxC)

	// calculate result
	var stringArr = checkCloudboxes(checkboxes)

	assert.Len(t, stringArr, 0)
}

func TestCheckCloudboxesNil(t *testing.T) {
	var result = checkCloudboxes(nil)

	assert.Empty(t, result)
}

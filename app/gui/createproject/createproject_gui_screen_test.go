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

	assert.Len(t, deployOnlyPrjForm.Items, 11)
	assert.Equal(t, "Namespace", deployOnlyPrjForm.Items[0].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[0].Widget)
	assert.Equal(t, "", deployOnlyPrjForm.Items[1].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[1].Widget)
	assert.Equal(t, "IP address or domain", deployOnlyPrjForm.Items[2].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[2].Widget)
	assert.Equal(t, "", deployOnlyPrjForm.Items[3].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[3].Widget)
	assert.Equal(t, "Jenkins system message", deployOnlyPrjForm.Items[4].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[4].Widget)
	assert.Equal(t, "", deployOnlyPrjForm.Items[5].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[5].Widget)
	assert.Equal(t, "Jenkins Jobs Repository", deployOnlyPrjForm.Items[6].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[6].Widget)
	assert.Equal(t, "", deployOnlyPrjForm.Items[7].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[7].Widget)
	assert.Equal(t, "Jenkins Existing PVC", deployOnlyPrjForm.Items[8].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[8].Widget)
	assert.Equal(t, "", deployOnlyPrjForm.Items[9].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[9].Widget)
	assert.Equal(t, "Cloud Templates", deployOnlyPrjForm.Items[10].Text)
	assert.IsType(t, &container.Scroll{}, deployOnlyPrjForm.Items[10].Widget)
}

func TestScreenCreateDeployOnlyProject(t *testing.T) {
	var deployOnlyPrjForm = ScreenCreateDeployOnlyProject(test.NewApp().NewWindow("test"))

	assert.Len(t, deployOnlyPrjForm.Items, 4)
	assert.Equal(t, "Namespace", deployOnlyPrjForm.Items[0].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[0].Widget)
	assert.Equal(t, "", deployOnlyPrjForm.Items[1].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[1].Widget)
	assert.Equal(t, "IP address or domain", deployOnlyPrjForm.Items[2].Text)
	assert.IsType(t, &widget.Entry{}, deployOnlyPrjForm.Items[2].Widget)
	assert.Equal(t, "", deployOnlyPrjForm.Items[3].Text)
	assert.IsType(t, &widget.Label{}, deployOnlyPrjForm.Items[3].Widget)
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

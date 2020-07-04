package createproject

import (
	"fyne.io/fyne/widget"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

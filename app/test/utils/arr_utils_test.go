package utils

import (
	"k8s-management-go/app/utils/arrays"
	"strings"
	"testing"
)

func TestIndexOfArr(t *testing.T) {
	// prepare array
	var testArray = createTestArray()

	// execute function
	var idx = arrays.IndexOfArr("World", testArray)

	// validate result
	if idx != 1 {
		t.Errorf("Function IndexOf has not found the right value (%v instead of 1", idx)
	} else {
		t.Log("Success finding index of element")
	}
}

func TestRemoveElementFromStringArr(t *testing.T) {
	// prepare array
	var testArray = createTestArray()

	// execute function
	var resultArr = arrays.RemoveElementFromStringArr(testArray, 1)

	// validate results
	if cap(resultArr) != 4 && strings.Join(resultArr, " ") != "Hello Here I Am" {
		t.Errorf("Right element was not removed from slice: [%s]", strings.Join(resultArr, " "))
	} else {
		t.Log("Success removing elements from slice")
	}
}

func createTestArray() []string {
	return []string{
		"Hello",
		"World",
		"Here",
		"I",
		"Am",
	}
}

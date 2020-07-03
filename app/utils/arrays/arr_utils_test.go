package arrays

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIndexOfArr(t *testing.T) {
	// prepare array
	var testArray = createTestArray()

	// execute function
	var idx = IndexOfArr("World", testArray)

	// validate result
	assert.Equal(t, 1, idx)
}

func TestIndexOfArrWithNonExistingValue(t *testing.T) {
	// prepare array
	var testArray = createTestArray()

	// execute function
	var idx = IndexOfArr("NotAvailable", testArray)

	// validate result
	assert.Equal(t, -1, idx)
}

func TestRemoveElementFromStringArr(t *testing.T) {
	// prepare array
	var testArray = createTestArray()

	// execute function
	var resultArr = RemoveElementFromStringArr(testArray, 1)

	// validate results
	assert.Len(t, resultArr, 4)
	assert.Equal(t, strings.Join(resultArr, " "), "Hello Here I Am")
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

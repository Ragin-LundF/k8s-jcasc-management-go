package loggingstate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddInfoEntry(t *testing.T) {
	var msg = "This is an info"
	AddInfoEntry(msg)
	entry := GetLoggingStateEntries()[0]

	assert.Equal(t, "INFO", entry.Type)
	assert.Equal(t, "", entry.Details)
	assert.Equal(t, msg, entry.Entry)

	// clear state
	ClearLoggingState()
}

func TestAddInfoEntryAndDetails(t *testing.T) {
	var msg = "This is an info"
	var details = "Here are more details"
	AddInfoEntryAndDetails(msg, details)
	entry := GetLoggingStateEntries()[0]

	assert.Equal(t, "INFO", entry.Type)
	assert.Equal(t, details, entry.Details)
	assert.Equal(t, msg, entry.Entry)

	// clear state
	ClearLoggingState()
}

func TestAddErrorEntry(t *testing.T) {
	var msg = "This is an error"
	AddErrorEntry(msg)
	entry := GetLoggingStateEntries()[0]

	assert.Equal(t, "ERROR", entry.Type)
	assert.Equal(t, "", entry.Details)
	assert.Equal(t, msg, entry.Entry)

	// clear state
	ClearLoggingState()
}

func TestAddErrorEntryAndDetails(t *testing.T) {
	var msg = "This is an error"
	var details = "Here are more error details"
	AddErrorEntryAndDetails(msg, details)
	entry := GetLoggingStateEntries()[0]

	assert.Equal(t, "ERROR", entry.Type)
	assert.Equal(t, details, entry.Details)
	assert.Equal(t, msg, entry.Entry)

	// clear state
	ClearLoggingState()
}

func TestClearLoggingState(t *testing.T) {
	AddErrorEntry("Error")
	AddInfoEntry("Info")
	assert.Greater(t, len(GetLoggingStateEntries()), 0)

	ClearLoggingState()
	assert.Nil(t, GetLoggingStateEntries())
}

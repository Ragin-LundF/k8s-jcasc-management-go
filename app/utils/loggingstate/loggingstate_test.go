package loggingstate

import (
	"testing"
)

func TestAddInfoEntry(t *testing.T) {
	var msg = "This is an info"
	AddInfoEntry(msg)
	entry := GetLoggingStateEntries()[0]

	if entry.Type == "INFO" && entry.Details == "" && entry.Entry == msg {
		t.Log("Success. INFO message without details added.")
	} else {
		t.Error("Failed. Can not find correct INFO message without details.")
	}
	ClearLoggingState()
}

func TestAddInfoEntryAndDetails(t *testing.T) {
	var msg = "This is an info"
	var details = "Here are more details"
	AddInfoEntryAndDetails(msg, details)
	entry := GetLoggingStateEntries()[0]

	if entry.Type == "INFO" && entry.Details == details && entry.Entry == msg {
		t.Log("Success. INFO message with details added.")
	} else {
		t.Errorf("Failed. Can not find correct INFO message with details. Type [%v] Entry [%v] Details [%v]", entry.Type, entry.Entry, entry.Details)
	}
	ClearLoggingState()
}

func TestAddErrorEntry(t *testing.T) {
	var msg = "This is an error"
	AddErrorEntry(msg)
	entry := GetLoggingStateEntries()[0]

	if entry.Type == "ERROR" && entry.Details == "" && entry.Entry == msg {
		t.Log("Success. ERROR message without details added.")
	} else {
		t.Error("Failed. Can not find correct ERROR message without details.")
	}
	ClearLoggingState()
}

func TestAddErrorEntryAndDetails(t *testing.T) {
	var msg = "This is an error"
	var details = "Here are more error details"
	AddErrorEntryAndDetails(msg, details)
	entry := GetLoggingStateEntries()[0]

	if entry.Type == "ERROR" && entry.Details == details && entry.Entry == msg {
		t.Log("Success. ERROR message with details added.")
	} else {
		t.Errorf("Failed. Can not find correct ERROR message with details. Type [%v] Entry [%v] Details [%v]", entry.Type, entry.Entry, entry.Details)
	}
	ClearLoggingState()
}

func TestClearLoggingState(t *testing.T) {
	AddErrorEntry("Error")
	AddInfoEntry("Info")
	loggingEntries := GetLoggingStateEntries()

	if cap(loggingEntries) > 0 {
		ClearLoggingState()
		loggingEntries = GetLoggingStateEntries()
		if loggingEntries == nil {
			t.Log("Success. LoggingStateEntries are nil.")
		} else {
			t.Error("Failed. LoggingStateEntries still have values.")
		}
	} else {
		t.Error("Failed. LoggingStateEntries are empty, but should have 2 values.")
	}
}

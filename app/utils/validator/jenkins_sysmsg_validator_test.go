package validator

import (
	"testing"
)

func TestValidateJenkinsSystemMessage(t *testing.T) {
	var jenkinsSysMsg = "Hello. This is a valid system message."
	err := ValidateJenkinsSystemMessage(jenkinsSysMsg)

	if err != nil {
		t.Error("Failed. Validator for Jenkins system message returned error.")
	} else {
		t.Log("Success. Validator accepted string.")
	}
}

func TestValidateJenkinsSystemMessageWithErrir(t *testing.T) {
	var jenkinsSysMsg = "This Message is longer than 255 characters, which is too long. The validator should reject this message. If not, it is an error. Lets hope that it works...This Message is longer than 255 characters, which is too long. The validator should reject this message. If not, it is an error. Lets hope that it works..."
	err := ValidateJenkinsSystemMessage(jenkinsSysMsg)

	if err != nil {
		t.Log("Success. Validator rejected the long message.")
	} else {
		t.Error("Failed. Validator for Jenkins system message did not return an error.")
	}
}

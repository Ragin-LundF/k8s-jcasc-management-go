package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateJenkinsSystemMessage(t *testing.T) {
	var jenkinsSysMsg = "Hello. This is a valid system message."
	err := ValidateJenkinsSystemMessage(jenkinsSysMsg)

	assert.NoError(t, err)
}

func TestValidateJenkinsSystemMessageWithErrir(t *testing.T) {
	var jenkinsSysMsg = "This Message is longer than 255 characters, which is too long. The validator should reject this message. If not, it is an error. Lets hope that it works...This Message is longer than 255 characters, which is too long. The validator should reject this message. If not, it is an error. Lets hope that it works..."
	err := ValidateJenkinsSystemMessage(jenkinsSysMsg)

	assert.Error(t, err)
}

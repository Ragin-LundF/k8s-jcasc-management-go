package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateConfirmPasswords(t *testing.T) {
	var password = "mypass"
	var confirmPw = password

	valid, _ := ValidateConfirmPasswords(password, confirmPw)

	assert.True(t, valid)
}

func TestValidateConfirmPasswordsTooShort(t *testing.T) {
	var password = "1234"
	var confirmPw = password

	valid, _ := ValidateConfirmPasswords(password, confirmPw)

	assert.False(t, valid)
}

func TestValidateConfirmPasswordsNotEqual(t *testing.T) {
	var password = "123456"
	var confirmPw = "654321"

	valid, _ := ValidateConfirmPasswords(password, confirmPw)

	assert.False(t, valid)
}

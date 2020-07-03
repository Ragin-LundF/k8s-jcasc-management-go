package validator

import "testing"

func TestValidateConfirmPasswords(t *testing.T) {
	var password = "mypass"
	var confirmPw = password

	valid, _ := ValidateConfirmPasswords(password, confirmPw)
	if valid {
		t.Log("Success. Both passwords are equal and valid.")
	} else {
		t.Error("Failed. Equal passwords with correct length not accepted.")
	}
}

func TestValidateConfirmPasswordsTooShort(t *testing.T) {
	var password = "1234"
	var confirmPw = password

	valid, _ := ValidateConfirmPasswords(password, confirmPw)
	if valid {
		t.Error("Failed. Equal, but too small passwords accepted.")
	} else {
		t.Log("Success. Equal, but too small passwords rejected.")
	}
}

func TestValidateConfirmPasswordsNotEqual(t *testing.T) {
	var password = "123456"
	var confirmPw = "654321"

	valid, _ := ValidateConfirmPasswords(password, confirmPw)
	if valid {
		t.Error("Failed. Not equal passwords accepted.")
	} else {
		t.Log("Success. Not equal passwords rejected.")
	}
}

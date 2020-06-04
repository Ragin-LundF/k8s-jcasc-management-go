package jenkinsuser

import (
	"errors"
	"github.com/atotto/clipboard"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/encryption"
	"strings"
)

func CreateJenkinsUserPassword() (info string, err error) {
	// Validator for password (keep it simple for now)
	validate := func(input string) error {
		if len(input) < 4 {
			return errors.New("Password too short!")
		}
		if strings.Contains(input, " ") {
			return errors.New("Password should not contain spaces!")
		}
		return nil
	}

	plainPassword, err := dialogs.DialogAskForPassword("Password", validate)
	plainPasswordConfirm, err := dialogs.DialogAskForPassword("Retype your password", validate)

	if plainPassword == plainPasswordConfirm {
		// encrypt password with bcrypt
		hashedPassword, err := encryption.EncryptJenkinsUserPassword(plainPassword)
		if err != nil {
			return info, err
		}

		templateDetails := `
--------- Encrypted Password ----------
{{ "Password    :" | faint }}	` + hashedPassword

		resultConfirm := dialogs.DialogConfirm(
			"Do you want to copy the password to the clipboard?",
			"Selection",
			templateDetails,
			"Your password: "+hashedPassword,
		)

		if resultConfirm {
			// copy to clipboard
			err = clipboard.WriteAll(hashedPassword)
		}
		return "Created password: " + hashedPassword, err
	} else {
		return info, errors.New("Passwords did not match!")
	}
}

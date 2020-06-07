package jenkinsuser

import (
	"errors"
	"github.com/atotto/clipboard"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/logger"
	"strings"
)

func CreateJenkinsUserPassword() (info string, err error) {
	log := logger.Log()
	log.Info("[Encrypt JenkinsUser Password] Ask for plain password...")
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
		log.Info("[Encrypt JenkinsUser Password] Entered passwords did match...")
		// encrypt password with bcrypt
		hashedPassword, err := encryption.EncryptJenkinsUserPassword(plainPassword)
		if err != nil {
			return info, err
		}
		log.Info("[Encrypt JenkinsUser Password] Password successfully encrypted...")

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
			log.Error("[Encrypt JenkinsUser Password] Unable to copy password to clipboard... %v\n", err)
		}
		return "Created password: " + hashedPassword, err
	} else {
		log.Error("[Encrypt JenkinsUser Password] Passwords did not match...")
		return info, errors.New("Passwords did not match! ")
	}
}

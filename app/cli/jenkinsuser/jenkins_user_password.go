package jenkinsuser

import (
	"errors"
	"github.com/atotto/clipboard"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/logger"
	"strings"
)

func CreateJenkinsUserPassword() (err error) {
	log := logger.Log()
	log.Info("[CreateJenkinsUserPassword] Ask for plain passwords...")
	loggingstate.AddInfoEntry("-> Ask for plain passwords...")
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

	// Get plain passwords
	plainPassword, err := dialogs.DialogAskForPassword("Password", validate)
	if err != nil {
		log.Error("[CreateJenkinsUserPassword] Unable to get plain password. \n%v", err)
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get plain password.", err.Error())
		return err
	}
	plainPasswordConfirm, err := dialogs.DialogAskForPassword("Retype your password", validate)
	if err != nil {
		log.Error("[CreateJenkinsUserPassword] Unable to get plain confirmation password. \n%v", err)
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get plain confirmation password.", err.Error())
	}

	// compare plain passwords
	if plainPassword == plainPasswordConfirm {
		log.Info("[CreateJenkinsUserPassword] Entered passwords did match...")
		loggingstate.AddInfoEntry("  -> Entered passwords did match...Try to encrypt...")
		// encrypt password with bcrypt
		hashedPassword, err := encryption.EncryptJenkinsUserPassword(plainPassword)
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Password successfully encrypted")
		log.Info("[CreateJenkinsUserPassword] Password successfully encrypted...")

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
			loggingstate.AddErrorEntryAndDetails("-> Unable to copy password to clipboard", err.Error())
			log.Error("[CreateJenkinsUserPassword] Unable to copy password to clipboard... %v\n", err)
		}
		return err
	} else {
		err = errors.New("Passwords did not match! ")
		loggingstate.AddErrorEntry("-> " + err.Error())
		log.Error("[CreateJenkinsUserPassword] %v", err.Error())
		return err
	}
}

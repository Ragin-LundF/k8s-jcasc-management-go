package jenkinsuser

import (
	"errors"
	"github.com/atotto/clipboard"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

func CreateJenkinsUserPassword() (err error) {
	log := logger.Log()
	log.Infof("[CreateJenkinsUserPassword] Ask for plain passwords...")
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
		log.Errorf("[CreateJenkinsUserPassword] Unable to get plain password. \n%s", err.Error())
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get plain password.", err.Error())
		return err
	}
	plainPasswordConfirm, err := dialogs.DialogAskForPassword("Retype your password", validate)
	if err != nil {
		log.Errorf("[CreateJenkinsUserPassword] Unable to get plain confirmation password. \n%s", err.Error())
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get plain confirmation password.", err.Error())
	}

	// compare plain passwords
	if plainPassword == plainPasswordConfirm {
		log.Infof("[CreateJenkinsUserPassword] Entered passwords did match...")
		loggingstate.AddInfoEntry("  -> Entered passwords did match...Try to encrypt...")
		// encrypt password with bcrypt
		hashedPassword, err := encryption.EncryptJenkinsUserPassword(plainPassword)
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Password successfully encrypted")
		log.Infof("[CreateJenkinsUserPassword] Password successfully encrypted...")

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
			if err != nil {
				loggingstate.AddErrorEntryAndDetails("-> Unable to copy password to clipboard", err.Error())
				log.Errorf("[CreateJenkinsUserPassword] Unable to copy password to clipboard... %s\n", err.Error())
			}
		}
		return err
	} else {
		err = errors.New("Passwords did not match! ")
		loggingstate.AddErrorEntry("-> " + err.Error())
		log.Errorf("[CreateJenkinsUserPassword] %s", err.Error())
		return err
	}
}

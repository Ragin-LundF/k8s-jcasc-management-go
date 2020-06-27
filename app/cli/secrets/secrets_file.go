package secrets

import (
	"errors"
	"k8s-management-go/app/actions/secrets_actions"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// encrypt the secrets file with given password
func EncryptSecretsFile() (err error) {
	// read password
	loggingstate.AddInfoEntry("  -> Ask for the password for secret file...")
	password, err := AskForSecretsPassword("Password for secrets file")
	if err != nil {
		return err
	}

	// let password confirm
	loggingstate.AddInfoEntry("  -> Ask for the confirmation password for secret file...")
	passwordConfirm, err := AskForSecretsPassword("Confirmation password for secrets file")
	if err != nil {
		return err
	}

	// check if passwords match
	if password != passwordConfirm {
		loggingstate.AddErrorEntry("  -> Passwords did not match! ")
		return errors.New("Passwords did not match! ")
	} else {
		loggingstate.AddErrorEntry("  -> Passwords did match! Starting encryption....")
	}

	// encrypt secrets file
	err = secrets_actions.ActionEncryptSecretsFile(password)

	return err
}

func DecryptSecretsFile() (err error) {
	password, err := AskForSecretsPassword("Password for secrets file")
	if err != nil {
		return err
	}
	err = secrets_actions.ActionDecryptSecretsFile(password)
	return err
}

// ask for password
func AskForSecretsPassword(passwordText string) (password string, err error) {
	// Validator for password (keep it simple for now)
	validate := func(input string) error {
		if len(input) < 4 {
			return errors.New("Password too short! ")
		}
		if strings.Contains(input, " ") {
			return errors.New("Password should not contain spaces! ")
		}
		return nil
	}

	// ask for password
	loggingstate.AddInfoEntry("  -> Ask for the password for secret file...")
	password, err = dialogs.DialogAskForPassword(passwordText, validate)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for the password for secret files...failed", err.Error())

		return "", err
	}
	loggingstate.AddInfoEntry("  -> Ask for the password for secret file...done")

	return password, err
}

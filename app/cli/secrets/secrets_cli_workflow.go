package secrets

import (
	"errors"
	"k8s-management-go/app/actions/secrets_actions"
	"k8s-management-go/app/utils/loggingstate"
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
		loggingstate.AddInfoEntry("  -> Passwords did match! Starting encryption....")
	}

	// encrypt secrets file
	err = secrets_actions.ActionEncryptSecretsFile(password)

	return err
}

// Decrypt secrets file
func DecryptSecretsFile() (err error) {
	password, err := AskForSecretsPassword("Password for secrets file")
	if err != nil {
		return err
	}
	err = secrets_actions.ActionDecryptSecretsFile(password)
	return err
}

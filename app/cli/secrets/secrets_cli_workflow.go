package secrets

import (
	"errors"
	"k8s-management-go/app/actions/secretsactions"
	"k8s-management-go/app/utils/loggingstate"
)

// EncryptSecretsFile encrypts the secrets file with given password
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
	}

	loggingstate.AddInfoEntry("  -> Passwords did match! Starting encryption....")
	// encrypt secrets file
	err = secretsactions.ActionEncryptSecretsFile(password)

	return err
}

// DecryptSecretsFile decrypts the secrets file
func DecryptSecretsFile() (err error) {
	password, err := AskForSecretsPassword("Password for secrets file")
	if err != nil {
		return err
	}
	err = secretsactions.ActionDecryptSecretsFile(password)
	return err
}

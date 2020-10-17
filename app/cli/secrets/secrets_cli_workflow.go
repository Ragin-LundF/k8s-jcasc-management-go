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
	secretsFile, password, err := AskForSecretsPassword("Password for secrets file", true)
	if err != nil {
		return err
	}

	// let password confirm
	loggingstate.AddInfoEntry("  -> Ask for the confirmation password for secret file...")
	_, passwordConfirm, err := AskForSecretsPassword("Confirmation password for secrets file", false)
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
	err = secretsactions.ActionEncryptSecretsFile(password, secretsFile)

	return err
}

// DecryptSecretsFile decrypts the secrets file
func DecryptSecretsFile(secretsFile *string) (err error) {
	var password string
	if secretsFile == nil {
		*secretsFile, password, err = AskForSecretsPassword("Password for secrets file", true)
	} else {
		_, password, err = AskForSecretsPassword("Password for secrets file", false)
	}
	if err != nil {
		return err
	}
	err = secretsactions.ActionDecryptSecretsFile(*secretsFile, password)
	return err
}

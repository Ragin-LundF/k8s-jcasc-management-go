package secrets

import (
	"errors"
	"k8s-management-go/app/actions/secretsactions"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
)

// EncryptSecretsFile encrypts the secrets file with given password
func EncryptSecretsFile() (err error) {
	// read password
	loggingstate.AddInfoEntry(constants.LogAskForPasswordOfSecretsFile)
	secretsFile, password, err := AskForSecretsPassword(constants.TextPasswordForSecretsFile, true)
	if err != nil {
		return err
	}

	// let password confirm
	loggingstate.AddInfoEntry(constants.LogAskForConfirmationPasswordOfSecretsFile)
	_, passwordConfirm, err := AskForSecretsPassword(constants.TextPasswordForSecretsFileConfirmation, false)
	if err != nil {
		return err
	}

	// check if passwords match
	if password != passwordConfirm {
		loggingstate.AddErrorEntry(constants.LogErrPasswordDidNotMatch)
		return errors.New(constants.TextPasswordDidNotMatch)
	}

	loggingstate.AddInfoEntry(constants.LogInfoPasswordDidMatchStartEncrypting)
	// encrypt secrets file
	err = secretsactions.ActionEncryptSecretsFile(password, secretsFile)

	return err
}

// DecryptSecretsFile decrypts the secrets file
func DecryptSecretsFile(secretsFile *string) (err error) {
	var password string
	if secretsFile == nil {
		var secretsFilePath string
		secretsFilePath, password, err = AskForSecretsPassword(constants.TextPasswordForSecretsFile, true)
		secretsFile = &secretsFilePath
	} else {
		_, password, err = AskForSecretsPassword(constants.TextPasswordForSecretsFile, false)
	}
	if err != nil {
		return err
	}
	err = secretsactions.ActionDecryptSecretsFile(password, *secretsFile)
	return err
}

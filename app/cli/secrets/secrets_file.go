package secrets

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// encrypt the secrets file with given password
func EncryptSecretsFile() (err error) {
	log := logger.Log()

	// read password
	log.Infof("[EncryptSecretsFile] Ask for the password for secret file...")
	loggingstate.AddInfoEntry("  -> Ask for the password for secret file...")
	password, err := AskForSecretsPassword("Password for secrets file")
	if err != nil {
		return err
	}

	// let password confirm
	log.Infof("[EncryptSecretsFile] Ask for the confirmation password for secret file...")
	loggingstate.AddInfoEntry("  -> Ask for the confirmation password for secret file...")
	passwordConfirm, err := AskForSecretsPassword("Confirmation password for secrets file")
	if err != nil {
		return err
	}

	// check if passwords match
	if password != passwordConfirm {
		loggingstate.AddErrorEntry("  -> Passwords did not match! ")
		log.Errorf("[EncryptSecretsFile] Passwords did not match! ")
		return errors.New("Passwords did not match! ")
	} else {
		loggingstate.AddErrorEntry("  -> Passwords did match! Starting encryption....")
		log.Errorf("[EncryptSecretsFile] Passwords did match! Starting encryption....")
	}

	// encrypt secrets file
	secretsFilePath := models.GetGlobalSecretsFile()
	err = encryption.GpgEncryptSecrets(secretsFilePath, password)

	return err
}

func DecryptSecretsFile() (err error) {
	password, err := AskForSecretsPassword("Password for secrets file")
	if err != nil {
		return err
	}
	err = DecryptSecretsFileWithPass(password)
	return err
}

// decrypt secrets file with password
func DecryptSecretsFileWithPass(password string) (err error) {
	secretsFilePath := models.GetGlobalSecretsFile() + constants.SecretsFileEncodedEnding
	err = encryption.GpgDecryptSecrets(secretsFilePath, password)

	return err
}

// ask for password
func AskForSecretsPassword(passwordText string) (password string, err error) {
	log := logger.Log()

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
	log.Infof("[AskForSecretsPassword] Ask for the password for secret file...")
	loggingstate.AddInfoEntry("  -> Ask for the password for secret file...")
	password, err = dialogs.DialogAskForPassword(passwordText, validate)
	if err != nil {
		log.Errorf("[AskForSecretsPassword] Ask for the password for secret files...failed, \n%s", err.Error())
		loggingstate.AddErrorEntryAndDetails("  -> Ask for the password for secret files...failed", err.Error())

		return "", err
	}
	loggingstate.AddInfoEntry("  -> Ask for the password for secret file...done")
	log.Infof("[AskForSecretsPassword] Ask for the password for secret file...done")

	return password, err
}

package secrets

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/logger"
	"strings"
)

func EncryptSecretsFile() (err error) {
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

	// read password
	password, err := dialogs.DialogAskForPassword("Password for secrets file", validate)
	if err != nil {
		return info, err
	}
	// let password confirm
	passwordConfirm, err := dialogs.DialogAskForPassword("Confirm password for secrets file", validate)
	if err != nil {
		return info, err
	}

	// check if passwords match
	if password != passwordConfirm {
		return "", errors.New("Passwords did not match!")
	}

	// encrypt secrets file
	secretsFilePath := models.GetGlobalSecretsFile()
	err = encryption.GpgEncryptSecrets(secretsFilePath, password)

	return info, err
}

// ask for password and decrypt secrets file
func DecryptSecretsFile() (err error) {
	log := logger.Log()

	// ask for password
	log.Info("[DecryptSecretsFile] Ask for the password for secret files...")
	loggingstate.AddInfoEntry("  -> Ask for the password for secret files...")
	password, err := dialogs.DialogAskForPassword("Password for secrets file", nil)
	if err != nil {
		log.Error("[DecryptSecretsFile] Ask for the password for secret files...failed, \n%v", err)
		loggingstate.AddErrorEntryAndDetails("  -> Ask for the password for secret files...failed", err.Error())

		return err
	}
	loggingstate.AddInfoEntry("  -> Ask for the password for secret files...done")
	log.Info("[DecryptSecretsFile] Ask for the password for secret files...done")

	secretsFilePath := models.GetGlobalSecretsFile() + constants.SecretsFileEncodedEnding
	err = encryption.GpgDecryptSecrets(secretsFilePath, password)

	return err
}

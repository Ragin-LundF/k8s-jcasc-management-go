package secrets

import (
	"errors"
	"k8s-management-go/cli/dialogs"
	"k8s-management-go/constants"
	"k8s-management-go/models/config"
	"k8s-management-go/utils/encryption"
	"log"
	"strings"
)

func EncryptSecretsFile() (info string, err error) {
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
		log.Println(err)
		return info, err
	}
	// let password confirm
	passwordConfirm, err := dialogs.DialogAskForPassword("Confirm password for secrets file", validate)
	if err != nil {
		log.Println(err)
		return info, err
	}

	// check if passwords match
	if password != passwordConfirm {
		return "", errors.New("Passwords did not match!")
	}

	// encrypt secrets file
	secretsFilePath := config.GetGlobalSecretsFile()
	info, err = encryption.GpgEncryptSecrets(secretsFilePath, password)

	return info, err
}

func DecryptSecretsFile() (info string, err error) {
	password, err := dialogs.DialogAskForPassword("Password for secrets file", nil)
	if err != nil {
		log.Println(err)
		return info, err
	}
	secretsFilePath := config.GetGlobalSecretsFile() + constants.SecretsFileEncodedEnding
	info, err = encryption.GpgDecryptSecrets(secretsFilePath, password)

	return info, err
}

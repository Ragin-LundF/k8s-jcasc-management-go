package cli

import (
	"errors"
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
	password, err := dialogPassword("Password for secrets file", validate)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// let password confirm
	passwordConfirm, err := dialogPassword("Confirm password for secrets file", validate)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// check if passwords match
	if password != passwordConfirm {
		return "", errors.New("Passwords did not match!")
	}

	// encrypt secrets file
	secretsFilePath := config.GetConfiguration().BasePath + "/" + config.GetConfiguration().GlobalSecretsFile
	info, err = encryption.GpgEncryptSecrets(secretsFilePath, password)

	return info, err
}

func DecryptSecretsFile() (info string, err error) {
	password, err := dialogPassword("Password for secrets file", nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	secretsFilePath := config.GetConfiguration().BasePath + "/" + config.GetConfiguration().GlobalSecretsFile + ".gpg"
	info, err = encryption.GpgDecryptSecrets(secretsFilePath, password)

	return info, err
}

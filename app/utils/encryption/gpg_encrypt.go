package encryption

import (
	"k8s-management-go/app/utils/logger"
	"os"
	"os/exec"
)

func GpgEncryptSecrets(secretsFilePath string, password string) (info string, err error) {
	log := logger.Log()
	cmd := exec.Command("gpg", "--batch", "--yes", "--passphrase", password, "-c", secretsFilePath)
	err = cmd.Run()
	if err != nil {
		log.Error(err)
	} else {
		info = "File [" + secretsFilePath + "] encrypted."
		// after everything was ok -> delete original file
		err = os.Remove(secretsFilePath)
	}
	return info, err
}

func GpgDecryptSecrets(secretsFilePath string, password string) (info string, err error) {
	log := logger.Log()
	cmd := exec.Command("gpg", "--batch", "--yes", "--passphrase", password, secretsFilePath)
	err = cmd.Run()
	if err != nil {
		log.Error(err)
	} else {
		info = "File [" + secretsFilePath + "] decrypted."
	}
	return info, err
}

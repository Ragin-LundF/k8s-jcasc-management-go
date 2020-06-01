package encryption

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func GpgEncryptSecrets(secretsFilePath string, password string) (info string, err error) {
	secretsFilePath = strings.Replace(secretsFilePath, "/./", "/", -1)
	cmd := exec.Command("gpg", "--batch", "--yes", "--passphrase", password, "-c", secretsFilePath)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	} else {
		info = "File [" + secretsFilePath + "] encrypted."
		// after everything was ok -> delete original file
		err = os.Remove(secretsFilePath)
	}
	return info, err
}

func GpgDecryptSecrets(secretsFilePath string, password string) (info string, err error) {
	secretsFilePath = strings.Replace(secretsFilePath, "/./", "/", -1)
	cmd := exec.Command("gpg", "--batch", "--yes", "--passphrase", password, secretsFilePath)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	} else {
		info = "File [" + secretsFilePath + "] decrypted."
	}
	return info, err
}

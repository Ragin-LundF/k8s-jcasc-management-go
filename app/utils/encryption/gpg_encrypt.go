package encryption

import (
	"os"
	"os/exec"
)

// encrypt secrets
func GpgEncryptSecrets(secretsFilePath string, password string) (info string, err error) {
	cmd := exec.Command("gpg", "--batch", "--yes", "--passphrase", password, "-c", secretsFilePath)
	err = cmd.Run()
	if err != nil {
		return info, err
	} else {
		info = "File [" + secretsFilePath + "] encrypted."
		// after everything was ok -> delete original file
		err = os.Remove(secretsFilePath)
	}
	return info, err
}

// decrypt secrets
func GpgDecryptSecrets(secretsFilePath string, password string) (info string, err error) {
	cmd := exec.Command("gpg", "--batch", "--yes", "--passphrase", password, secretsFilePath)
	err = cmd.Run()
	if err != nil {
		return info, err
	} else {
		info = "File [" + secretsFilePath + "] decrypted."
	}
	return info, err
}

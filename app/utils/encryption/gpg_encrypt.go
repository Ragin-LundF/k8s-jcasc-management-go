package encryption

import (
	"k8s-management-go/app/utils/logger"
	"os"
	"os/exec"
)

// encrypt secrets
func GpgEncryptSecrets(secretsFilePath string, password string) (info string, err error) {
	log := logger.Log()
	log.Info("[GPG Encrypt] Try to encrypt secrets file [" + secretsFilePath + "]")
	cmd := exec.Command("gpg", "--batch", "--yes", "--passphrase", password, "-c", secretsFilePath)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("[GPG Encrypt] Unable to encrypt secrets file [" + secretsFilePath + "]...")
		log.Error(cmdOutput)
		log.Error(err)
		return info, err
	}

	log.Info("[GPG Encrypt] Encrypt secrets file [" + secretsFilePath + "] done.")
	info = "Secrets file [" + secretsFilePath + "] encrypted."
	// after everything was ok -> delete original file
	err = os.Remove(secretsFilePath)
	if err != nil {
		log.Error("[GPG Encrypt] Unable to remove secrets file ["+secretsFilePath+"]... %v\n", err)
	}

	return info, err
}

// decrypt secrets
func GpgDecryptSecrets(secretsFilePath string, password string) (info string, err error) {
	log := logger.Log()
	cmd := exec.Command("gpg", "--batch", "--yes", "--passphrase", password, secretsFilePath)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("[GPG Dencrypt] Unable to decrypt secrets file [" + secretsFilePath + "]...")
		log.Error(cmdOutput)
		log.Error(err)
		return info, err
	}
	log.Info("[GPG Decrypt] Decrypt secrets file [" + secretsFilePath + "] done.")
	info = "Secrets file [" + secretsFilePath + "] decrypted."

	return info, err
}

package secretsactions

import (
	"fmt"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/loggingstate"
	"os"
)

// ActionEncryptSecretsFile encrypts the secrets file
func ActionEncryptSecretsFile(password string) (err error) {
	secretsFilePath := models.GetGlobalSecretsFile()
	err = encryption.GpgEncryptSecrets(secretsFilePath, password)

	if err != nil {
		return err
	}

	// after everything was ok -> delete original file
	err = os.Remove(secretsFilePath)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to delete decrypted secrets file [%s].", secretsFilePath), err.Error())
		return err
	}

	return err
}

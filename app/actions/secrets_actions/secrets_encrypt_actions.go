package secrets_actions

import (
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/encryption"
)

// encrypt secrets file
func ActionEncryptSecretsFile(password string) (err error) {
	secretsFilePath := models.GetGlobalSecretsFile()
	err = encryption.GpgEncryptSecrets(secretsFilePath, password)

	return err
}

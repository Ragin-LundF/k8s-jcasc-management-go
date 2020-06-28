package secrets_actions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/encryption"
)

// decrypt secrets file with password
func ActionDecryptSecretsFile(password string) (err error) {
	secretsFilePath := models.GetGlobalSecretsFile() + constants.SecretsFileEncodedEnding
	err = encryption.GpgDecryptSecrets(secretsFilePath, password)

	return err
}

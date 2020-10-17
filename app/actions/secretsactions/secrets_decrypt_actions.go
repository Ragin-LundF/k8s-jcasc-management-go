package secretsactions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/encryption"
)

// ActionDecryptSecretsFile decrypts secrets file with password
func ActionDecryptSecretsFile(password string, fileName string) (err error) {
	var secretsFilePath = models.GetGlobalSecretsPath() + fileName + constants.SecretsFileEncodedEnding
	err = encryption.GpgDecryptSecrets(secretsFilePath, password)

	return err
}

package secretsactions

import (
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/loggingstate"
	"os"
)

// ActionEncryptSecretsFileByName encrypts the given secrets file
func ActionEncryptSecretsFile(password string, secretsFileName string) (err error) {
	var secretsFilePath = configuration.GetConfiguration().GetGlobalSecretsPath() + secretsFileName
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

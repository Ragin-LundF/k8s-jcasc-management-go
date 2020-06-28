package secrets_actions

import (
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"os/exec"
)

func ActionApplySecretsToNamespace(namespace string) (err error) {
	// apply secret to namespace
	secretsFilePath := models.GetGlobalSecretsFile()
	if err = executingSecretsFile(secretsFilePath, namespace); err != nil {
		// try to delete it and return original error
		_ = removeDecryptedSecretsFile(secretsFilePath)
		return err
	}
	err = removeDecryptedSecretsFile(secretsFilePath)

	return err
}

func ActionApplySecretsToAllNamespaces(callback func()) (err error) {
	secretsFilePath := models.GetGlobalSecretsFile()
	// apply secret to namespaces
	for _, ip := range models.GetIpConfiguration().Ips {
		if err = executingSecretsFile(secretsFilePath, ip.Namespace); err != nil {
			_ = removeDecryptedSecretsFile(secretsFilePath)
			return err
		}
		callback()
	}
	err = removeDecryptedSecretsFile(secretsFilePath)

	return err
}

// execute the secrets file
func executingSecretsFile(secretsFilePath string, namespace string) (err error) {
	// execute decrypted file
	_ = exec.Command("chmod", "755", secretsFilePath).Run()
	cmd := exec.Command(secretsFilePath)
	cmd.Env = append(os.Environ(),
		"NAMESPACE="+namespace,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Executing secrets file failed. See output", string(output))
		loggingstate.AddErrorEntryAndDetails("  -> Executing secrets file failed. See errors", err.Error())
	} else {
		loggingstate.AddInfoEntryAndDetails("  -> Executing secrets file done. See output", string(output))
	}

	return err
}

// remove decrypted secrets file
func removeDecryptedSecretsFile(secretsFilePath string) (err error) {
	// delete decrypted file
	if err := os.Remove(secretsFilePath); err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to remove decrypted secrets file.", err.Error())
	}
	return err
}

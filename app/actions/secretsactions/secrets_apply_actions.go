package secretsactions

import (
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"os/exec"
)

// ActionApplySecretsToNamespace executes the secrets to the namespace
func ActionApplySecretsToNamespace(namespace string, secretsFileName string) (err error) {
	// apply secret to namespace
	var secretsFilePath = configuration.GetConfiguration().GetGlobalSecretsPath() + secretsFileName
	if err = executingSecretsFile(secretsFilePath, namespace); err != nil {
		// try to delete it and return original error
		_ = removeDecryptedSecretsFile(secretsFilePath)
		return err
	}
	err = removeDecryptedSecretsFile(secretsFilePath)

	return err
}

// ActionApplySecretsToAllNamespaces applies secrets to all namespaces
func ActionApplySecretsToAllNamespaces(secretsFileName string, callback func()) (err error) {
	secretsFilePath := configuration.GetConfiguration().GetGlobalSecretsPath() + secretsFileName
	// apply secret to namespaces
	for _, ip := range configuration.GetConfiguration().K8SManagement.IPConfig.Deployments {
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

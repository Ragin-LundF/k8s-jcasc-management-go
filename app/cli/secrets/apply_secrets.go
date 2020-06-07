package secrets

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"os"
	"os/exec"
)

func ApplySecrets() (err error) {
	log := logger.Log()
	// ask for namespace
	log.Info("  -> Ask for namespace to apply secrets...")
	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...")

	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for namespace to apply secrets...failed", err.Error())
		return err
	}

	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...done")
	log.Info("  -> Ask for namespace to apply secrets...done")

	// apply secrets to namespace
	log.Info("  -> Apply secrets to  namespace [%v]...", namespace)
	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...")

	if err = ApplySecretsToNamespace(namespace); err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for namespace to apply secrets...failed", err.Error())
		return err
	}

	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...done")
	log.Info("  -> Apply secrets to  namespace [%v]...done", namespace)

	return nil
}

// apply secrets to one namespace
func ApplySecretsToNamespace(namespace string) (err error) {
	log := logger.Log()
	// Decrypt secrets file
	if err = DecryptSecretsFile(); err != nil {
		return err
	}

	// apply secret to namespace
	secretsFilePath := models.GetGlobalSecretsFile()
	errorsWhileApply := false
	if err = executingSecretsFile(secretsFilePath, namespace); err != nil {
		errorsWhileApply = true
	}

	// delete decrypted file
	if err := os.Remove(secretsFilePath); err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to remove decrypted secrets file.", err.Error())
		log.Error("[ApplySecretsToAllNamespaces] Unable to remove decrypted secrets file.\n%v", err)
		return err
	}

	if errorsWhileApply {
		return errors.New("Unable to execute secret file for namespace [" + namespace + "]. See previous errors. ")
	}
	return nil
}

// apply secrets to all namespaces
func ApplySecretsToAllNamespaces() (err error) {
	log := logger.Log()
	// apply secret to namespaces
	errorsWhileApply := false
	secretsFilePath := models.GetGlobalSecretsFile()
	for _, ip := range models.GetIpConfiguration().Ips {
		if err := executingSecretsFile(secretsFilePath, ip.Namespace); err != nil {
			errorsWhileApply = true
		}
	}

	// delete decrypted file
	if err := os.Remove(secretsFilePath); err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to remove decrypted secrets file.", err.Error())
		log.Error("[ApplySecretsToAllNamespaces] Unable to remove decrypted secrets file.\n%v", err)
		return err
	}

	if errorsWhileApply {
		return errors.New("Unable to execute all secret files. See previous errors. ")
	}
	return nil
}

// execute the secrets file
func executingSecretsFile(secretsFilePath string, namespace string) (err error) {
	log := logger.Log()

	// execute decrypted file
	cmd := exec.Command("sh", "-c", secretsFilePath)
	cmd.Env = append(os.Environ(),
		"NAMESPACE="+namespace,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("[executingSecretsFile] Executing secrets file failed. Output: \n%v", string(output))
		log.Error("[executingSecretsFile] Executing secrets file failed. Errors: \n%v", err)
		loggingstate.AddErrorEntryAndDetails("  -> Executing secrets file failed. See output", string(output))
		loggingstate.AddErrorEntryAndDetails("  -> Executing secrets file failed. See errors", err.Error())
		return err
	} else {
		log.Info("[executingSecretsFile] Executing secrets file done. Output: \n%v", string(output))
		loggingstate.AddInfoEntryAndDetails("  -> Executing secrets file done. See output", string(output))
	}

	return nil
}

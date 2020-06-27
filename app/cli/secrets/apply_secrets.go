package secrets

import (
	"errors"
	"fmt"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"os/exec"
)

func ApplySecrets() (err error) {
	log := logger.Log()
	// ask for namespace
	log.Infof("  -> Ask for namespace to apply secrets...")
	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...")

	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for namespace to apply secrets...failed", err.Error())
		return err
	}

	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...done")
	log.Infof("  -> Ask for namespace to apply secrets...done")

	// apply secrets to namespace
	log.Infof("  -> Apply secrets to  namespace [%s]...", namespace)
	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...")

	if err = ApplySecretsToNamespace(namespace, nil); err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Ask for namespace to apply secrets...failed", err.Error())
		return err
	}

	loggingstate.AddInfoEntry("  -> Ask for namespace to apply secrets...done")
	log.Infof("  -> Apply secrets to  namespace [%s]...done", namespace)

	return nil
}

// apply secrets to one namespace
func ApplySecretsToNamespace(namespace string, password *string) (err error) {
	log := logger.Log()
	// Decrypt secrets file
	if password != nil {
		if err = DecryptSecretsFileWithPass(*password); err != nil {
			return err
		}
	} else {
		if err = DecryptSecretsFile(); err != nil {
			return err
		}
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
		log.Errorf("[ApplySecretsToAllNamespaces] Unable to remove decrypted secrets file.\n%s", err.Error())
		return err
	}

	if errorsWhileApply {
		return errors.New(fmt.Sprintf("Unable to execute secret file for namespace [%s]. See previous errors. ", namespace))
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
		log.Errorf("[ApplySecretsToAllNamespaces] Unable to remove decrypted secrets file.\n%s", err.Error())
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
	_ = exec.Command("chmod", "755", secretsFilePath).Run()
	cmd := exec.Command(secretsFilePath)
	cmd.Env = append(os.Environ(),
		"NAMESPACE="+namespace,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("[executingSecretsFile] Executing secrets file failed. Output: \n%s", string(output))
		log.Errorf("[executingSecretsFile] Executing secrets file failed. Errors: \n%s", err.Error())
		loggingstate.AddErrorEntryAndDetails("  -> Executing secrets file failed. See output", string(output))
		loggingstate.AddErrorEntryAndDetails("  -> Executing secrets file failed. See errors", err.Error())
		return err
	} else {
		log.Infof("[executingSecretsFile] Executing secrets file done. Output: \n%s", string(output))
		loggingstate.AddInfoEntryAndDetails("  -> Executing secrets file done. See output", string(output))
	}

	return nil
}

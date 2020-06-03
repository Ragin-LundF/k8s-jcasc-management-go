package secrets

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/encryption"
	"k8s-management-go/app/utils/logger"
	"os"
	"os/exec"
)

func ApplySecrets() (info string, err error) {
	log := logger.Log()
	// select namespace
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		log.Error(err)
		return info, err
	}
	info, err = ApplySecretsToNamespace(namespace)
	return info, err
}

func ApplySecretsToNamespace(namespace string) (info string, err error) {
	log := logger.Log()
	// get password
	password, err := dialogs.DialogAskForPassword("Password for secrets file", nil)
	if err != nil {
		log.Error(err)
		return info, err
	}

	// decrypt secrets file
	secretsFilePath := config.GetGlobalSecretsFile()
	info, err = encryption.GpgDecryptSecrets(secretsFilePath+constants.SecretsFileEncodedEnding, password)
	if err != nil {
		log.Error(err)
		return info, err
	}

	// apply secret to namespace
	info, nsErr := applySecretToNamespace(secretsFilePath, namespace)

	// delete decrypted file
	rmErr := os.Remove(secretsFilePath)

	// Error handling for apply and remove
	if nsErr != nil {
		err = nsErr
	}
	if rmErr != nil {
		if err != nil {
			err = errors.New(err.Error() + rmErr.Error())
		} else {
			err = rmErr
		}
	}

	return info, err
}

func ApplySecretsToAllNamespaces() (info string, err error) {
	log := logger.Log()
	// get password
	password, err := dialogs.DialogAskForPassword("Password for secrets file", nil)
	if err != nil {
		log.Error(err)
		return info, err
	}

	// decrypt secrets file
	secretsFilePath := config.GetGlobalSecretsFile()
	info, err = encryption.GpgDecryptSecrets(secretsFilePath+constants.SecretsFileEncodedEnding, password)
	if err != nil {
		log.Error(err)
		return info, err
	}

	// apply secret to namespaces
	infos := ""
	nsErrs := ""
	for _, ip := range config.GetIpConfiguration().Ips {
		infoNs, nsErr := applySecretToNamespace(secretsFilePath, ip.Namespace)
		if infoNs != "" {
			infos = infos + "\n" + infoNs
		}
		if nsErr != nil {
			nsErrs = nsErrs + "\n" + nsErr.Error()
		}
	}

	// delete decrypted file
	rmErr := os.Remove(secretsFilePath)

	// Error handling for apply and remove
	if nsErrs != "" {
		err = errors.New(nsErrs)
	}
	if rmErr != nil {
		if err != nil {
			err = errors.New(err.Error() + rmErr.Error())
		} else {
			err = rmErr
		}
	}

	return info, err
}

func applySecretToNamespace(secretsFilePath string, namespace string) (info string, err error) {
	log := logger.Log()
	// execute decrypted file
	cmd := exec.Command("sh", "-c", secretsFilePath)
	cmd.Env = append(os.Environ(),
		"NAMESPACE="+namespace,
	)
	err = cmd.Run()
	if err != nil {
		log.Error(err)
	} else {
		info = "Secrets to namespace [" + namespace + "] applied"
	}

	return info, err
}

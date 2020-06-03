package cli

import (
	"errors"
	"k8s-management-go/constants"
	"k8s-management-go/models/config"
	"k8s-management-go/utils/encryption"
	"log"
	"os"
	"os/exec"
)

func ApplySecrets() (info string, err error) {
	// select namespace
	namespace, err := DialogAskForNamespace()
	if err != nil {
		log.Println(err)
		return info, err
	}
	info, err = ApplySecretsToNamespace(namespace)
	return info, err
}

func ApplySecretsToNamespace(namespace string) (info string, err error) {
	// get password
	password, err := DialogPassword("Password for secrets file", nil)
	if err != nil {
		log.Println(err)
		return info, err
	}

	// decrypt secrets file
	secretsFilePath := config.GetGlobalSecretsFile()
	info, err = encryption.GpgDecryptSecrets(secretsFilePath+constants.SecretsFileEncodedEnding, password)
	if err != nil {
		log.Println(err)
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
	// get password
	password, err := DialogPassword("Password for secrets file", nil)
	if err != nil {
		log.Println(err)
		return info, err
	}

	// decrypt secrets file
	secretsFilePath := config.GetGlobalSecretsFile()
	info, err = encryption.GpgDecryptSecrets(secretsFilePath+constants.SecretsFileEncodedEnding, password)
	if err != nil {
		log.Println(err)
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
	// execute decrypted file
	cmd := exec.Command("sh", "-c", secretsFilePath)
	cmd.Env = append(os.Environ(),
		"NAMESPACE="+namespace,
	)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	} else {
		info = "Secrets to namespace [" + namespace + "] applied"
	}

	return info, err
}

package secrets

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"os"
	"os/exec"
)

func ApplySecrets() (info string, err error) {
	// select namespace
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		return info, err
	}
	info, err = ApplySecretsToNamespace(namespace)
	return info, err
}

// apply secrets to one namespace
func ApplySecretsToNamespace(namespace string) (info string, err error) {
	// Decrypt secrets file
	infoLog, err := DecryptSecretsFile()
	info = info + constants.NewLine + infoLog

	// apply secret to namespace
	secretsFilePath := config.GetGlobalSecretsFile()
	infoLog, nsErr := applySecretsToNamespace(secretsFilePath, namespace)
	info = info + constants.NewLine + infoLog

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

// apply secrets to all namespaces
func ApplySecretsToAllNamespaces() (info string, err error) {
	// apply secret to namespaces
	infos := ""
	nsErrs := ""
	secretsFilePath := config.GetGlobalSecretsFile()
	for _, ip := range config.GetIpConfiguration().Ips {
		infoNs, nsErr := applySecretsToNamespace(secretsFilePath, ip.Namespace)
		if infoNs != "" {
			infos = infos + constants.NewLine + infoNs
		}
		if nsErr != nil {
			nsErrs = nsErrs + constants.NewLine + nsErr.Error()
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

// execute the secrets file
func applySecretsToNamespace(secretsFilePath string, namespace string) (info string, err error) {
	// execute decrypted file
	cmd := exec.Command("sh", "-c", secretsFilePath)
	cmd.Env = append(os.Environ(),
		"NAMESPACE="+namespace,
	)
	output, err := cmd.Output()
	if err != nil {
		info = info + constants.NewLine + "Apply secrets failed! See errors."
		return info, err
	} else {
		info = info + constants.NewLine + "Applied secrets to namespace [" + namespace + "]"
		info = info + constants.NewLine + "============================="
		info = info + constants.NewLine + string(output)
		info = info + constants.NewLine + "============================="
	}

	return info, err
}

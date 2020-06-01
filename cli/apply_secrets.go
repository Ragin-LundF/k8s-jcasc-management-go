package cli

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"k8s-management-go/constants"
	"k8s-management-go/models/config"
	"k8s-management-go/utils/encryption"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ApplySecretsToNamespace() (info string, err error) {
	// select namespace
	namespace, err := selectNamespace()
	if err != nil {
		log.Println(err)
		return info, err
	}

	// get password
	password, err := dialogPassword("Password for secrets file", nil)
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

func selectNamespace() (namespace string, err error) {
	// Template for displaying menu
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A4 {{ .Namespace | green }}",
		Inactive: "  {{ .Namespace | cyan }}",
		Selected: "\U000027A4 {{ .Namespace | red | cyan }}",
		Details: `
--------- Namespace selection ----------
{{ "Namespace: " | faint }}	{{ .Namespace }}
{{ "IP       : " | faint }}	{{ .Ip }}`,
	}

	// searcher (with "/")
	searcher := func(input string, index int) bool {
		namespaceItem := config.GetIpConfiguration().Ips[index]
		name := strings.Replace(strings.ToLower(namespaceItem.Namespace), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Please select the namespace to which the secrets should be applied",
		Items:     config.GetIpConfiguration().Ips,
		Templates: templates,
		Size:      12,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	} else {
		namespace = config.GetIpConfiguration().Ips[i].Namespace
	}

	return namespace, err
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

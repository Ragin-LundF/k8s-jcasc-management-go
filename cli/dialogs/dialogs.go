package dialogs

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"k8s-management-go/models/config"
	"log"
	"strings"
)

type confirm struct {
	Selection string
}

func ClearScreen() {
	fmt.Println("\033[2J")
}

// Configurable confirm dialog
func DialogConfirm(templateLabel string, templateSelector string, templateDetails string, dialogLabel string) bool {
	ClearScreen()

	// Template for displaying confirm dialog
	dialogConfim := []confirm{
		{Selection: "yes"},
		{Selection: "no"},
	}

	// prepare template
	templates := &promptui.SelectTemplates{
		Label:    templateLabel,
		Active:   "\U000027A4 {{ ." + templateSelector + " | green }}",
		Inactive: "  {{ .Selection | cyan }}",
		Selected: "\U000027A4 {{ ." + templateSelector + " | red | cyan }}",
		Details:  templateDetails,
	}

	// dialog prompt
	confirmDialogPrompt := promptui.Select{
		Label:     dialogLabel,
		Items:     dialogConfim,
		Templates: templates,
		Size:      2,
	}

	// execute dialog
	_, resultConfirm, err := confirmDialogPrompt.Run()

	// result processing
	if err != nil || resultConfirm != "{yes}" {
		if err != nil {
			log.Printf("Prompt failed %v\n", err)
		}
		return false
	} else {
		return true
	}
}

// Common password dialog
func DialogAskForPassword(label string, validate promptui.ValidateFunc) (password string, err error) {
	ClearScreen()

	// Prepare prompt
	promptPlainPassword := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Mask:     '*',
	}
	password, err = promptPlainPassword.Run()

	// check if everything was ok
	if err != nil {
		log.Printf("Prompt failed %v\n", err)
	}
	return password, err
}

// Ask for deployment name
func DialogAskForDeploymentName(label string, validate promptui.ValidateFunc) (deploymentName string, err error) {
	ClearScreen()

	// try to read deployment name from configuration
	deploymentName = config.GetConfiguration().Jenkins.Helm.Master.DeploymentName
	// check if something was set
	if deploymentName == "" {
		// No pre-configured deployment name found -> ask for a new one
		// Prepare prompt
		promptDeploymentName := promptui.Prompt{
			Label:    label,
			Validate: validate,
		}
		deploymentName, err = promptDeploymentName.Run()

		// check if everything was ok
		if err != nil {
			log.Printf("Prompt failed %v\n", err)
		}
	}
	return deploymentName, err
}

// dialog to ask for the namespace
func DialogAskForNamespace() (namespace string, err error) {
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
		log.Printf("Prompt failed %v\n", err)
	} else {
		namespace = config.GetIpConfiguration().Ips[i].Namespace
	}

	return namespace, err
}

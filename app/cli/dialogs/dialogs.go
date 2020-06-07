package dialogs

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/arrays"
	"k8s-management-go/app/utils/logger"
	"strings"
)

type confirm struct {
	Selection string
}

type CloudTemplatesDialog struct {
	CloudTemplateFiles     []string
	SelectedCloudTemplates []string
}

func ClearScreen() {
	fmt.Println("\033[2J")
}

// Configurable confirm dialog
func DialogConfirm(templateLabel string, templateSelector string, templateDetails string, dialogLabel string) bool {
	log := logger.Log()
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
			log.Error("[DialogConfirm] Prompt confirm dialog failed %v\n", err)
		}
		return false
	} else {
		return true
	}
}

// Common password dialog
func DialogAskForPassword(label string, validate promptui.ValidateFunc) (password string, err error) {
	log := logger.Log()
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
		log.Error("[DialogAskForPassword] Prompt ask for password failed %v\n", err)
	}
	return password, err
}

// prompt to enter something
func DialogPrompt(label string, validate promptui.ValidateFunc) (answer string, err error) {
	promptEntry := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	answer, err = promptEntry.Run()

	return answer, err
}

// Ask for deployment name
func DialogAskForDeploymentName(label string, validate promptui.ValidateFunc) (deploymentName string, err error) {
	log := logger.Log()
	ClearScreen()

	// try to read deployment name from configuration
	deploymentName = models.GetConfiguration().Jenkins.Helm.Master.DeploymentName
	// check if something was set
	if deploymentName == "" {
		// No pre-configured deployment name found -> ask for a new one
		// Prepare prompt
		deploymentName, err = DialogPrompt(label, validate)
		// check if everything was ok
		if err != nil {
			log.Error("[DialogAskForDeploymentName] Prompt ask for deployment name failed %v\n", err)
		}
	}
	return deploymentName, err
}

// dialog to ask for the namespace
func DialogAskForNamespace() (namespace string, err error) {
	log := logger.Log()
	ClearScreen()

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
		namespaceItem := models.GetIpConfiguration().Ips[index]
		name := strings.Replace(strings.ToLower(namespaceItem.Namespace), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Please select the namespace to which the secrets should be applied",
		Items:     models.GetIpConfiguration().Ips,
		Templates: templates,
		Size:      12,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Error("[DialogAskForNamespace] Prompt ask for namespace failed %v\n", err)
	} else {
		namespace = models.GetIpConfiguration().Ips[i].Namespace
	}

	return namespace, err
}

// dialog to ask for cloud templates
func DialogAskForCloudTemplates(cloudTemplateDialog *CloudTemplatesDialog) (err error) {
	log := logger.Log()
	ClearScreen()

	// Template for displaying menu
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A4 {{ . | green }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "\U000027A4 {{ . | red | cyan }}",
		Details: `
--------- Selected Cloud Templates: ----------
{{ "Selected templates: " | faint }}	` + strings.Join(cloudTemplateDialog.SelectedCloudTemplates, ", "),
	}

	// searcher (with "/")
	searcher := func(input string, index int) bool {
		templateItem := cloudTemplateDialog.CloudTemplateFiles[index]
		name := strings.Replace(strings.ToLower(templateItem), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Please select the cloud templates you want to use for this namespace",
		Items:     cloudTemplateDialog.CloudTemplateFiles,
		Templates: templates,
		Size:      12,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Error("[DialogAskForCloudTemplates] Prompt ask for cloud templates failed %v\n", err)
	} else {
		if i > 0 {
			foundElement := -1
			// first look, if entry already exists and if it was found, remove it
			for idx, selectedElementInCloudTemplates := range cloudTemplateDialog.SelectedCloudTemplates {
				if selectedElementInCloudTemplates == cloudTemplateDialog.CloudTemplateFiles[i] {
					cloudTemplateDialog.SelectedCloudTemplates = arrays.RemoveElementFromStringArr(cloudTemplateDialog.SelectedCloudTemplates, idx)
					foundElement = idx
					break
				}
			}

			// element was not found -> add it
			if foundElement == -1 {
				cloudTemplateDialog.SelectedCloudTemplates = append(cloudTemplateDialog.SelectedCloudTemplates, cloudTemplateDialog.CloudTemplateFiles[i])
			}
			err = DialogAskForCloudTemplates(cloudTemplateDialog)
		}
	}

	return err
}

// Show info and error output as select prompt with search
func DialogShowLogging(loggingStateEntries []loggingstate.LoggingState) {
	log := logger.Log()
	// clear screen
	ClearScreen()

	// if there is something to show, create dialog and show the log
	if cap(loggingStateEntries) > 0 {
		// Template for displaying MenuitemModel
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U000027A4 [{{ .Type | green }}] {{ .Entry | white }}",
			Inactive: "  [{{ .Type | cyan }}] {{ .Entry | red }}",
			Selected: "\U000027A4 [{{ .Type | red | cyan }}] {{ .Entry | red }}",
			Details: `
--------- Log Entry ----------
{{ "Type   :" | faint }}	{{ .Type }}
{{ "Message:" | faint }}	{{ .Entry }}
{{ "Details:" | faint }}
{{.Details}}`,
		}

		// searcher (with "/")
		searcher := func(input string, index int) bool {
			logItem := loggingStateEntries[index]
			logEntry := strings.Replace(strings.ToLower(logItem.Entry), " ", "", -1)
			logType := strings.Replace(strings.ToLower(logItem.Type), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(logEntry, input) || strings.Contains(logType, input)
		}

		prompt := promptui.Select{
			Label:     "Log Output. Press Enter to leave this view",
			Items:     loggingStateEntries,
			Templates: templates,
			Size:      20,
			Searcher:  searcher,
		}

		_, _, err := prompt.Run()

		if err != nil {
			log.Error(err)
		}
	}
}

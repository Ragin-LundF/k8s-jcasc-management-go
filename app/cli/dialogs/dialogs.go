package dialogs

import (
	"github.com/inancgumus/screen"
	"github.com/manifoldco/promptui"
	"github.com/schollz/progressbar/v3"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/arrays"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

type confirm struct {
	Selection string
}

// CloudTemplatesDialog represents cloud template files and their selections
type CloudTemplatesDialog struct {
	CloudTemplateFiles     []string
	SelectedCloudTemplates []string
}

// ClearScreen clears the screen
func ClearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

// DialogConfirm is a configurable confirm dialog
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
		Stdout:    &BellSkipper{},
	}

	// execute dialog
	_, resultConfirm, err := confirmDialogPrompt.Run()

	// result processing
	if err != nil || resultConfirm != "{yes}" {
		if err != nil {
			log.Errorf("[DialogConfirm] Prompt confirm dialog failed %s\n", err.Error())
		}
		return false
	}
	return true
}

// DialogAskForPassword is a common password dialog
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
		log.Errorf("[DialogAskForPassword] Prompt ask for password failed %s\n", err.Error())
	}
	return password, err
}

// DialogPrompt prompts to enter something
func DialogPrompt(label string, validate promptui.ValidateFunc) (answer string, err error) {
	promptEntry := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	answer, err = promptEntry.Run()

	return answer, err
}

// DialogAskForDeploymentName asks for deployment name
func DialogAskForDeploymentName(label string, validate promptui.ValidateFunc) (deploymentName string, err error) {
	log := logger.Log()

	// try to read deployment name from configuration
	deploymentName = configuration.GetConfiguration().Jenkins.Controller.DeploymentName
	// check if something was set
	if deploymentName == "" {
		ClearScreen()
		// No pre-configured deployment name found -> ask for a new one
		// Prepare prompt
		deploymentName, err = DialogPrompt(label, validate)
		// check if everything was ok
		if err != nil {
			log.Errorf("[DialogAskForDeploymentName] Prompt ask for deployment name failed %s\n", err.Error())
		}
	}
	return deploymentName, err
}

// DialogAskForNamespace shows dialog to ask for the namespace
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
{{ "IP       : " | faint }}	{{ .IP }}`,
	}

	// searcher (with "/")
	searcher := func(input string, index int) bool {
		namespaceItem := configuration.GetConfiguration().K8SManagement.IPConfig.Deployments[index]
		name := strings.Replace(strings.ToLower(namespaceItem.Namespace), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Please select the namespace to which the secrets should be applied",
		Items:     configuration.GetConfiguration().K8SManagement.IPConfig.Deployments,
		Templates: templates,
		Size:      12,
		Searcher:  searcher,
		Stdout:    &BellSkipper{},
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Errorf("[DialogAskForNamespace] Prompt ask for namespace failed %s\n", err.Error())
	} else {
		namespace = configuration.GetConfiguration().K8SManagement.IPConfig.Deployments[i].Namespace
	}

	return namespace, err
}

// DialogAskForSecretsFile shows dialog to ask for the secrets file
func DialogAskForSecretsFile() (secretsFile string, err error) {
	log := logger.Log()
	ClearScreen()
	var secretFilesArray = configuration.GetConfiguration().GetSecretsFiles()

	// Template for displaying menu
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A4 {{ . | green }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "\U000027A4 {{ . | red | cyan }}",
		Details: `
--------- Secrets file selection ----------
{{ "SecretsFile: " | faint }}	{{ . }}`,
	}

	// searcher (with "/")
	searcher := func(input string, index int) bool {
		var secretFileItem = secretFilesArray[index]

		return strings.Contains(secretFileItem, input)
	}

	prompt := promptui.Select{
		Label:     "Please select the secrets file which should be applied",
		Items:     secretFilesArray,
		Templates: templates,
		Size:      12,
		Searcher:  searcher,
		Stdout:    &BellSkipper{},
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Errorf("[DialogAskForSecretsFile] Prompt ask for secrets file failed %s\n", err.Error())
	} else {
		secretsFile = secretFilesArray[i]
	}

	return secretsFile, err
}

// DialogAskForCloudTemplates shows dialog to ask for cloud templates
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
		Stdout:    &BellSkipper{},
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Errorf("[DialogAskForCloudTemplates] Prompt ask for cloud templates failed %s\n", err.Error())
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

// DialogShowLogging shows info and error output as select prompt with search
func DialogShowLogging(loggingStateEntries []loggingstate.LoggingState, err error) {
	log := logger.Log()
	// clear screen
	ClearScreen()

	// if there is something to show, create dialog and show the log
	if len(loggingStateEntries) > 0 {
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

		var errorHint = " \U00002705 Successful. Possible errors can be ignored. "
		if err != nil {
			errorHint = " \U0000274C Unsuccessful. Please check errors! "
		}
		prompt := promptui.Select{
			Label:     ".:===" + errorHint + "===:. -> Press <Return> to leave this view <-",
			Items:     loggingStateEntries,
			Templates: templates,
			Size:      20,
			Searcher:  searcher,
			Stdout:    &BellSkipper{},
		}

		_, _, err := prompt.Run()

		if err != nil {
			log.Errorf(err.Error())
		}
	}
}

// CreateProgressBar creates a preconfigured progress bar
func CreateProgressBar(description string, progressMax int) progressbar.ProgressBar {
	bar := progressbar.NewOptions(progressMax,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetDescription(description),
		progressbar.OptionSpinnerType(50),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	return *bar
}

// ProgressBar is a structure which contains the progress bar
type ProgressBar struct {
	Bar *progressbar.ProgressBar
}

// AddCallback is a function to add progress. Will be used as callback
func (progress *ProgressBar) AddCallback() {
	_ = progress.Bar.Add(1)
}

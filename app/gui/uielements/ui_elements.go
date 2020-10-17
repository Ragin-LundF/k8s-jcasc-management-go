package uielements

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/kubernetesactions"
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// ProgressBar configures the progress bar
type ProgressBar struct {
	Bar        *dialog.ProgressDialog
	MaxCount   float64
	CurrentCnt float64
}

// AddCallback is a function to add progress. Will be used as callback
func (progress *ProgressBar) AddCallback() {
	progress.Bar.SetValue(float64(1) / progress.MaxCount * progress.CurrentCnt)
	progress.CurrentCnt = progress.CurrentCnt + 1
}

// CreateNamespaceSelectEntry creates namespace select entry
func CreateNamespaceSelectEntry(namespaceErrorLabel *widget.Label) (namespaceSelectEntry *widget.SelectEntry) {
	// Namespace
	namespaceSelectEntry = widget.NewSelectEntry(namespaceactions.ActionReadNamespaceWithFilter(nil))
	namespaceSelectEntry.PlaceHolder = "Type or select namespace"
	namespaceSelectEntry.OnChanged = func(input string) {
		namespaces := namespaceactions.ActionReadNamespaceWithFilter(&input)
		namespaceSelectEntry.SetOptions(namespaces)
		if strings.TrimSpace(strings.Join(namespaces, "")) == "" {
			namespaceErrorLabel.SetText("No namespace found with these characters.")
		} else {
			namespaceErrorLabel.SetText("")
		}
	}

	return namespaceSelectEntry
}

// CreateKubernetesContextSelectEntry creates kubernetes context select entry
func CreateKubernetesContextSelectEntry(k8sErrorLabel *widget.Label) (k8sContextSelectEntry *widget.SelectEntry) {
	// K8S Context
	k8sContextSelectEntry = widget.NewSelectEntry(kubernetesactions.ActionReadK8SContextWithFilter(nil))
	k8sContextSelectEntry.PlaceHolder = "Type or select context name"
	k8sContextSelectEntry.OnChanged = func(input string) {
		k8sContext := kubernetesactions.ActionReadK8SContextWithFilter(&input)
		k8sContextSelectEntry.SetOptions(k8sContext)
		if strings.TrimSpace(strings.Join(k8sContext, "")) == "" {
			k8sErrorLabel.SetText("No contexts found with these characters.")
		} else {
			k8sErrorLabel.SetText("")
		}
	}

	return k8sContextSelectEntry
}

// CreateDeploymentNameEntry creates deployment name entry
func CreateDeploymentNameEntry() (deploymentNameEntry *widget.Entry) {
	// Deployment name
	deploymentNameEntry = widget.NewEntry()
	deploymentNameEntry.SetPlaceHolder("Deployment name")
	if models.GetConfiguration().Jenkins.Helm.Master.DeploymentName != "" {
		deploymentNameEntry.Text = models.GetConfiguration().Jenkins.Helm.Master.DeploymentName
		deploymentNameEntry.Disable()
	}
	return deploymentNameEntry
}

// CreateSecretsFileEntry creates a dropdown which contains a selection of secret files
func CreateSecretsFileEntry() (secretsFileEntry *widget.Select) {
	var secretFiles []string
	var alternativeSecretFiles = models.GetSecretsFiles()
	if alternativeSecretFiles != nil && len(*alternativeSecretFiles) > 0 {
		secretFiles = append(secretFiles, *alternativeSecretFiles...)
	}
	secretsFileEntry = widget.NewSelect(secretFiles, func(s string) {})
	secretsFileEntry.SetSelected(secretsFileEntry.Options[0])
	return secretsFileEntry
}

// CreateInstallTypeRadio creates radio install type radio
func CreateInstallTypeRadio() (radioInstallType *widget.Radio) {
	// Install or update
	radioInstallType = widget.NewRadio([]string{constants.HelmCommandInstall, constants.HelmCommandUpgrade}, nil)
	radioInstallType.SetSelected(constants.HelmCommandInstall)

	return radioInstallType
}

// CreateDryRunRadio creates radio install type radio
func CreateDryRunRadio() (radioInstallType *widget.Radio) {
	// Execute or dry-run
	radioInstallType = widget.NewRadio([]string{constants.InstallDryRunInactive, constants.InstallDryRunActive}, nil)
	radioInstallType.SetSelected(constants.InstallDryRunInactive)

	return radioInstallType
}

// ShowLogOutput shows output of internal logging
func ShowLogOutput(window fyne.Window) {
	// read the log
	loggingStates := loggingstate.GetLoggingStateEntries()
	// clear the log
	loggingstate.ClearLoggingState()

	// prepare accordion
	logAccordion := widget.NewAccordionContainer()
	var accItem *widget.AccordionItem
	var accLabel *widget.Label

	for _, logState := range loggingStates {
		if logState.Details != "" {
			accLabel = widget.NewLabel(processLogTextForBestView(logState.Details))
			accItem = widget.NewAccordionItem(processLogTextForBestView("["+logState.Type+"] "+logState.Entry), accLabel)
			logAccordion.Append(accItem)
		} else {
			accLabel = widget.NewLabel("No content...")
			accItem = widget.NewAccordionItem(processLogTextForBestView("["+logState.Type+"] "+logState.Entry), accLabel)
			logAccordion.Append(accItem)
		}
	}

	scrollContainer := widget.NewScrollContainer(logAccordion)
	scrollContainer.SetMinSize(fyne.Size{
		Width:  700,
		Height: 400,
	})
	loggingstate.LogLoggingStateEntries()
	dialog.ShowCustom("", "Ok", scrollContainer, window)
}

func processLogTextForBestView(label string) string {
	var resultText string
	var characterCount = 0

	// split
	wordArr := strings.Split(label, " ")
	for _, word := range wordArr {
		characterCount += len(word)
		if characterCount > 100 {
			resultText += "\n"
			characterCount = 0
		}
		resultText += word + " "
	}
	return resultText
}

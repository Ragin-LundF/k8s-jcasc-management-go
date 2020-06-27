package ui_elements

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"image/color"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"sort"
	"strings"
)

// create namespace select entry
func CreateNamespaceSelectEntry(namespaceErrorLabel *widget.Label) (namespaceSelectEntry *widget.SelectEntry) {
	// Namespace
	namespaceSelectEntry = widget.NewSelectEntry(FindNamespacesForSelect(nil))
	namespaceSelectEntry.PlaceHolder = "Type or select namespace"
	namespaceSelectEntry.OnChanged = func(input string) {
		namespaces := FindNamespacesForSelect(&input)
		namespaceSelectEntry.SetOptions(namespaces)
		if strings.TrimSpace(strings.Join(namespaces, "")) == "" {
			namespaceErrorLabel.SetText("No namespace found with these characters.")
		} else {
			namespaceErrorLabel.SetText("")
		}
	}

	return namespaceSelectEntry
}

// create deployment name entry
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

// create radio install_actions type radio
func CreateInstallTypeRadio() (radioInstallType *widget.Radio) {
	// Install or update
	radioInstallType = widget.NewRadio([]string{constants.HelmCommandInstall, constants.HelmCommandUpgrade}, nil)
	radioInstallType.SetSelected(constants.HelmCommandInstall)

	return radioInstallType
}

// create radio install_actions type radio
func CreateDryRunRadio() (radioInstallType *widget.Radio) {
	// Execute or dry-run
	radioInstallType = widget.NewRadio([]string{constants.InstallDryRunInactive, constants.InstallDryRunActive}, nil)
	radioInstallType.SetSelected(constants.InstallDryRunInactive)

	return radioInstallType
}

// namespaces loader and filter
func FindNamespacesForSelect(filter *string) (namespaces []string) {
	ipList := models.GetIpConfiguration().Ips
	for _, ip := range ipList {
		if filter != nil && *filter != "" {
			if strings.Contains(ip.Namespace, *filter) {
				namespaces = append(namespaces, ip.Namespace)
			}
		} else {
			namespaces = append(namespaces, ip.Namespace)
		}
	}
	sort.Strings(namespaces)
	return namespaces
}

func ShowLogOutput(window fyne.Window) {
	// read the log
	loggingStates := loggingstate.GetLoggingStateEntries()
	// clear the log
	loggingstate.ClearLoggingState()

	// create text grid
	grid := widget.NewTextGrid()
	red := &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 128, G: 0, B: 0, A: 255}}

	var currentLine = 0
	var linesWithError []int
	var textContent string
	for _, logState := range loggingStates {
		lineContent := fmt.Sprintf("[%s] %s", logState.Type, logState.Entry)
		textContent = fmt.Sprintf("%s%s\n", textContent, lineContent)
		if logState.Type == "ERROR" {
			linesWithError = append(linesWithError, currentLine)
		}
		currentLine++

		if logState.Details != "" {
			textContent = fmt.Sprintf("%s--- Details start----\n", textContent)
			textContent = fmt.Sprintf("%s%s\n", textContent, logState.Details)
			textContent = fmt.Sprintf("%s--- Details end----\n", textContent)
			if logState.Type == "ERROR" {
				linesWithError = append(linesWithError, currentLine, currentLine+1, currentLine+2, currentLine+3)
			}

			currentLine = currentLine + 3
		}

		textContent = fmt.Sprintf("%s\n", textContent)
		currentLine++
	}
	grid.SetText(textContent)

	for _, errRow := range linesWithError {
		grid.Rows[errRow].Style = red
	}

	grid.ShowLineNumbers = true
	grid.ShowWhitespace = true

	scrollContainer := widget.NewScrollContainer(grid)
	scrollContainer.SetMinSize(fyne.NewSize(700, 400))

	dialog.ShowCustom("", "Ok", scrollContainer, window)
}

// check selected namespace against namespace list
func ValidateNamespace(namespaceToValidate string) bool {
	for _, ip := range models.GetIpConfiguration().Ips {
		if ip.Namespace == namespaceToValidate {
			return true
		}
	}
	return false
}

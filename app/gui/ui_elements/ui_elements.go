package ui_elements

import (
	"fyne.io/fyne/widget"
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

// check selected namespace against namespace list
func ValidateNamespace(namespaceToValidate string) bool {
	for _, ip := range models.GetIpConfiguration().Ips {
		if ip.Namespace == namespaceToValidate {
			return true
		}
	}
	return false
}

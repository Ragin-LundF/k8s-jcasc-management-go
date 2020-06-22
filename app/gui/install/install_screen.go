package install

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"sort"
	"strings"
)

func ScreenInstall(window fyne.Window) fyne.CanvasObject {
	var namespace string
	var deploymentName string
	var installTypeOption string
	var dryRunOption string
	var secretsPasswords string

	// Namespace
	namespaceErrorLabel := widget.NewLabel("")
	namespaceSelectEntry := createNamespaceSelectEntry(namespaceErrorLabel)

	// Deployment name
	deploymentNameEntry := createDeploymentNameEntry()

	// Install or update
	installTypeRadio := createInstallTypeRadio()

	// Dry-run or execute
	dryRunRadio := createDryRunRadio()

	// secrets password
	secretsPasswordEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceSelectEntry},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "Deployment Name", Widget: deploymentNameEntry},
			{Text: "Installation type", Widget: installTypeRadio},
			{Text: "Execute or dry run", Widget: dryRunRadio},
		},
		OnSubmit: func() {
			// get variables
			namespace = namespaceSelectEntry.Text
			deploymentName = deploymentNameEntry.Text
			installTypeOption = installTypeRadio.Selected
			dryRunOption = dryRunRadio.Selected
			if !validateNamespace(namespace) {
				namespaceErrorLabel.SetText("Error: namespace is unknown!")
				namespaceErrorLabel.Show()
				return
			}

			// ask for password
			if dryRunOption == constants.InstallDryRunInactive {
				openSecretsPasswordDialog(window, secretsPasswordEntry, secretsPasswords)
			}
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)

	return box
}

// create namespace select entry
func createNamespaceSelectEntry(namespaceErrorLabel *widget.Label) (namespaceSelectEntry *widget.SelectEntry) {
	// Namespace
	namespaceSelectEntry = widget.NewSelectEntry(findNamespacesForSelect(nil))
	namespaceSelectEntry.PlaceHolder = "Type or select namespace"
	namespaceSelectEntry.OnChanged = func(input string) {
		namespaces := findNamespacesForSelect(&input)
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
func createDeploymentNameEntry() (deploymentNameEntry *widget.Entry) {
	// Deployment name
	deploymentNameEntry = widget.NewEntry()
	deploymentNameEntry.SetPlaceHolder("Deployment name")
	if models.GetConfiguration().Jenkins.Helm.Master.DeploymentName != "" {
		deploymentNameEntry.Text = models.GetConfiguration().Jenkins.Helm.Master.DeploymentName
		deploymentNameEntry.Disable()
	}
	return deploymentNameEntry
}

// create radio install type radio
func createInstallTypeRadio() (radioInstallType *widget.Radio) {
	// Install or update
	radioInstallType = widget.NewRadio([]string{constants.InstallTypeInstall, constants.InstallTypeUpgrade}, nil)
	radioInstallType.SetSelected(constants.InstallTypeInstall)

	return radioInstallType
}

// create radio install type radio
func createDryRunRadio() (radioInstallType *widget.Radio) {
	// Execute or dry-run
	radioInstallType = widget.NewRadio([]string{constants.InstallDryRunInactive, constants.InstallDryRunActive}, nil)
	radioInstallType.SetSelected(constants.InstallDryRunInactive)

	return radioInstallType
}

// Secrets password dialog
func openSecretsPasswordDialog(window fyne.Window, secretsPasswordEntry *widget.Entry, secretsPassword string) {
	secretsPasswordWindow := widget.NewForm(widget.NewFormItem("Password", secretsPasswordEntry))
	secretsPasswordWindow.Resize(fyne.Size{Width: 300, Height: 90})

	dialog.ShowCustomConfirm("Password for Secrets...", "Ok", "Cancel", secretsPasswordWindow, func(confirmed bool) {
		if !confirmed {
			secretsPassword = secretsPasswordEntry.Text
		} else {
			return
		}
	}, window)
}

// namespaces loader and filter
func findNamespacesForSelect(filter *string) (namespaces []string) {
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
func validateNamespace(namespaceToValidate string) bool {
	for _, ip := range models.GetIpConfiguration().Ips {
		if ip.Namespace == namespaceToValidate {
			return true
		}
	}
	return false
}

package install

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/models"
	"sort"
	"strings"
)

func ScreenInstall(window fyne.Window) fyne.CanvasObject {
	// Namespace
	selectNamespaceEntry := createNamespace()

	// Deployment name
	deploymentNameEntry := createDeploymentName()

	// Install or update
	radioInstallType := createRadioInstallType()

	// secrets password
	secretsPassword := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: selectNamespaceEntry},
			{Text: "Deployment Name", Widget: deploymentNameEntry},
			{Text: "Installation type", Widget: radioInstallType},
		},
		OnSubmit: func() {
			openSecretsPasswordDialog(window, secretsPassword)

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   fmt.Sprintf("%s of %s started", radioInstallType.Selected, deploymentNameEntry.Text),
				Content: fmt.Sprintf("Starting on namespace [%s]...", selectNamespaceEntry.Text),
			})
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)

	return box
}

// create namespace select entry
func createNamespace() (namespaceSelectEntry *widget.SelectEntry) {
	// Namespace
	namespaceSelectEntry = widget.NewSelectEntry(findNamespacesForSelect(nil))
	namespaceSelectEntry.PlaceHolder = "Type or select namespace"
	namespaceSelectEntry.OnChanged = func(input string) {
		namespaces := findNamespacesForSelect(&input)
		namespaceSelectEntry.SetOptions(namespaces)
	}
	return namespaceSelectEntry
}

// create deployment name entry
func createDeploymentName() (deploymentNameEntry *widget.Entry) {
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
func createRadioInstallType() (radioInstallType *widget.Radio) {
	// Install or update
	radioInstallType = widget.NewRadio([]string{"install", "upgrade"}, func(s string) { fmt.Println("selected", s) })
	radioInstallType.SetSelected("install")

	return radioInstallType
}

// Secrets password dialog
func openSecretsPasswordDialog(window fyne.Window, secretsPassword *widget.Entry) {
	secretsPasswordWindow := widget.NewForm(widget.NewFormItem("Password", secretsPassword))

	dialog.ShowCustomConfirm("Password for Secrets...", "Ok", "Cancel", secretsPasswordWindow, func(b bool) {
		if !b {
			return
		}
	}, window)
}

func findNamespacesForSelect(filter *string) (namespaces []string) {
	ipList := models.GetIpConfiguration().Ips
	for _, ip := range ipList {
		if filter != nil {
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

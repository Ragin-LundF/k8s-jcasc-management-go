package install

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/models"
	"sort"
	"strings"
)

func ScreenInstall() fyne.CanvasObject {
	// Namespace
	selectNamespaceEntry := widget.NewSelectEntry(findNamespacesForSelect(nil))
	selectNamespaceEntry.PlaceHolder = "Type or select namespace"
	selectNamespaceEntry.OnChanged = func(input string) {
		namespaces := findNamespacesForSelect(&input)
		selectNamespaceEntry.SetOptions(namespaces)
	}

	// Deployment name
	deploymentNameEntry := widget.NewEntry()
	deploymentNameEntry.SetPlaceHolder("Deployment name")
	if models.GetConfiguration().Jenkins.Helm.Master.DeploymentName != "" {
		deploymentNameEntry.Text = models.GetConfiguration().Jenkins.Helm.Master.DeploymentName
		deploymentNameEntry.Disable()
	}

	// Install or update
	radioInstallType := widget.NewRadio([]string{"install", "upgrade"}, func(s string) { fmt.Println("selected", s) })
	radioInstallType.SetSelected("install")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: selectNamespaceEntry},
			{Text: "Deployment Name", Widget: deploymentNameEntry},
			{Text: "Installation type", Widget: radioInstallType},
		},
		OnSubmit: func() {
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

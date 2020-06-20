package install

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/models"
	"sort"
	"strings"
)

func ScreenInstall() fyne.CanvasObject {
	selectNamespaceEntry := widget.NewSelectEntry(findNamespacesForSelect(nil))
	selectNamespaceEntry.PlaceHolder = "Type or select namespace"
	selectNamespaceEntry.OnChanged = func(input string) {
		namespaces := findNamespacesForSelect(&input)
		selectNamespaceEntry.SetOptions(namespaces)
	}

	deploymentNameEntry := widget.NewEntry()
	deploymentNameEntry.SetPlaceHolder("Deployment name")

	box := widget.NewVBox(
		&*selectNamespaceEntry,
		deploymentNameEntry,
		widget.NewHBox(layout.NewSpacer()),
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

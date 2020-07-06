package resources

import "fyne.io/fyne"

// K8sJcascMgmtIcon returns the Icon as a StaticResource
func K8sJcascMgmtIcon() *fyne.StaticResource {
	return fyne.NewStaticResource("Icon.png", resourceIconPng.StaticContent)
}

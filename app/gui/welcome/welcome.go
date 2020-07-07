package welcome

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/gui/resources"
)

// ScreenWelcome shows the welcome screen
func ScreenWelcome(info string) fyne.CanvasObject {
	// set label
	var labelInfo *widget.Label

	logo := canvas.NewImageFromResource(resources.K8sJcascMgmtIcon())
	logo.SetMinSize(fyne.NewSize(128, 128))

	if info == "" {
		labelInfo = widget.NewLabelWithStyle("You are on the latest version.", fyne.TextAlignCenter, fyne.TextStyle{
			Italic: true,
		})
	} else {
		labelInfo = widget.NewLabelWithStyle(info, fyne.TextAlignCenter, fyne.TextStyle{
			Bold: true,
		})
	}
	return widget.NewVBox(
		layout.NewSpacer(),
		labelInfo,
		widget.NewLabelWithStyle("Welcome to Kubernetes Jenkins Configuration as Code", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		layout.NewSpacer(),
	)
}

package k8s_screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func WelcomeScreen() fyne.CanvasObject {
	logo := canvas.NewImageFromFile("docs/images/k8s-mgmt-workflow.png")
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(1064/4, 1145/4))
	} else {
		logo.SetMinSize(fyne.NewSize(1064/2, 1145/2))
	}

	return widget.NewVBox(
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Welcome to Kubernetes Jenkins Configuration as Code", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
	)
}

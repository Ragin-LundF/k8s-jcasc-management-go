package welcome

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func ScreenWelcome(info string) fyne.CanvasObject {
	logo := canvas.NewImageFromFile("docs/images/k8s-mgmt-workflow.png")
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(1064/4, 1145/4))
	} else {
		logo.SetMinSize(fyne.NewSize(1064/2, 1145/2))
	}

	// set label
	var labelInfo *widget.Label

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
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Welcome to Kubernetes Jenkins Configuration as Code", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
	)
}

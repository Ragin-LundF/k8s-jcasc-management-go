package createproject

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/createprojectactions"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/validator"
)

// ScreenCreateFullProject shows the full project setup screen
func ScreenCreateFullProject(window fyne.Window) fyne.CanvasObject {
	var projectConfig models.ProjectConfig
	projectConfig.CreateDeploymentOnlyProject = false

	// Namespace
	namespaceErrorLabel := widget.NewLabel("")
	namespaceEntry := widget.NewEntry()
	namespaceEntry.PlaceHolder = "my-namespace"

	// IP address
	ipAddressErrorLabel := widget.NewLabel("")
	ipAddressEntry := widget.NewEntry()
	ipAddressEntry.PlaceHolder = "0.0.0.0"

	// Jenkins system message
	jenkinsSysMsgErrorLabel := widget.NewLabel("")
	jenkinsSysMsgEntry := widget.NewEntry()
	jenkinsSysMsgEntry.PlaceHolder = constants.CommonJenkinsSystemMessage

	// Jenkins jobs config repository
	jenkinsJobsCfgErrorLabel := widget.NewLabel("")
	jenkinsJobsCfgEntry := widget.NewEntry()
	jenkinsJobsCfgEntry.PlaceHolder = "http://vcs.domain.tld/project/repo/jenkins-jobs.git"

	// todo cloud templates

	// Existing PVC
	jenkinsExistingPvcErrorLabel := widget.NewLabel("")
	jenkinsExistingPvcEntry := widget.NewEntry()
	jenkinsExistingPvcEntry.PlaceHolder = "pvc-jenkins"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceEntry},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "IP address", Widget: ipAddressEntry},
			{Text: "", Widget: ipAddressErrorLabel},
			{Text: "Jenkins system message", Widget: jenkinsSysMsgEntry},
			{Text: "", Widget: jenkinsSysMsgErrorLabel},
			{Text: "Jenkins Jobs Repository", Widget: jenkinsJobsCfgEntry},
			{Text: "", Widget: jenkinsJobsCfgErrorLabel},
			{Text: "Jenkins Existing PVC", Widget: jenkinsExistingPvcEntry},
			{Text: "", Widget: jenkinsExistingPvcErrorLabel},
		},
		OnSubmit: func() {
			// get variables
			projectConfig.Namespace = namespaceEntry.Text
			projectConfig.IPAddress = ipAddressEntry.Text
			projectConfig.JenkinsSystemMsg = jenkinsSysMsgEntry.Text
			projectConfig.JobsCfgRepo = jenkinsJobsCfgEntry.Text
			projectConfig.ExistingPvc = jenkinsExistingPvcEntry.Text
			hasErrors := false

			// validate namespace
			err := validator.ValidateNewNamespace(projectConfig.Namespace)
			if err != nil {
				namespaceErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// validate IP address
			err = validator.ValidateIP(projectConfig.IPAddress)
			if err != nil {
				ipAddressErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// validate jenkins system message
			err = validator.ValidateJenkinsSystemMessage(projectConfig.JenkinsSystemMsg)
			if err != nil {
				jenkinsSysMsgErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// validate jobs config
			err = validator.ValidateJenkinsJobConfig(projectConfig.JobsCfgRepo)
			if err != nil {
				jenkinsJobsCfgErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// validate PVC
			err = validator.ValidatePersistentVolumeClaim(projectConfig.ExistingPvc)
			if err != nil {
				jenkinsExistingPvcErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// process project creation if no error was found
			if !hasErrors {
				bar := uielements.ProgressBar{
					Bar:        dialog.NewProgress("Create project...", "Progress", window),
					CurrentCnt: 0,
					MaxCount:   createprojectactions.CountCreateProjectWorkflow,
				}
				bar.Bar.Show()
				_ = createprojectactions.ActionProcessProjectCreate(projectConfig, bar.AddCallback)
				bar.Bar.Hide()

				// show output
				uielements.ShowLogOutput(window)
			}
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)

	return box
}

// ScreenCreateDeployOnlyProject shows the screen for deployment only project without Jenkins
func ScreenCreateDeployOnlyProject(window fyne.Window) fyne.CanvasObject {
	var projectConfig models.ProjectConfig
	projectConfig.CreateDeploymentOnlyProject = true

	// Namespace
	namespaceErrorLabel := widget.NewLabel("")
	namespaceEntry := widget.NewEntry()
	namespaceEntry.PlaceHolder = "my-namespace"

	// IP address
	ipAddressErrorLabel := widget.NewLabel("")
	ipAddressEntry := widget.NewEntry()
	ipAddressEntry.PlaceHolder = "0.0.0.0"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceEntry},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "IP address", Widget: ipAddressEntry},
			{Text: "", Widget: ipAddressErrorLabel},
		},
		OnSubmit: func() {
			// get variables
			projectConfig.Namespace = namespaceEntry.Text
			projectConfig.IPAddress = ipAddressEntry.Text
			hasError := false

			// validate namespace
			err := validator.ValidateNewNamespace(projectConfig.Namespace)
			if err != nil {
				namespaceErrorLabel.SetText(err.Error())
				hasError = true
			}

			// validate IP address
			err = validator.ValidateIP(projectConfig.IPAddress)
			if err != nil {
				ipAddressErrorLabel.SetText(err.Error())
				hasError = true
			}

			if !hasError {
				// process project creation
				bar := uielements.ProgressBar{
					Bar:        dialog.NewProgress("Create project...", "Progress", window),
					CurrentCnt: 0,
					MaxCount:   createprojectactions.CountCreateProjectWorkflow,
				}
				bar.Bar.Show()
				_ = createprojectactions.ActionProcessProjectCreate(projectConfig, bar.AddCallback)
				bar.Bar.Hide()

				// show output
				uielements.ShowLogOutput(window)
			}
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer()),
		form,
	)

	return box
}

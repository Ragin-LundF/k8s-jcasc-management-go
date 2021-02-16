package createproject

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"k8s-management-go/app/actions/createprojectactions"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/validator"
	"strings"
)

// ScreenCreateFullProject shows the full project setup screen form
func ScreenCreateFullProject(window fyne.Window) *widget.Form {
	var projectConfig models.ProjectConfig
	projectConfig.CreateDeploymentOnlyProject = false

	// Namespace
	namespaceErrorLabel := widget.NewLabel("")
	namespaceEntry := widget.NewEntry()
	namespaceEntry.PlaceHolder = "my-namespace"

	// IP address
	ipAddressErrorLabel := widget.NewLabel("")
	ipAddressEntry := widget.NewEntry()
	ipAddressEntry.PlaceHolder = "0.0.0.0 or mydomain.tld"

	// Domain for Jenkins
	jenkinsUrlErrorLabel := widget.NewLabel("")
	jenkinsUrlEntry := widget.NewEntry()
	jenkinsUrlEntry.PlaceHolder = "domain.tld (or leave empty to use <namespace>.<configured nginx domain>)"

	// Jenkins system message
	jenkinsSysMsgErrorLabel := widget.NewLabel("")
	jenkinsSysMsgEntry := widget.NewEntry()
	jenkinsSysMsgEntry.PlaceHolder = constants.CommonJenkinsSystemMessage

	// Jenkins jobs config repository
	jenkinsJobsCfgErrorLabel := widget.NewLabel("")
	jenkinsJobsCfgEntry := widget.NewEntry()
	jenkinsJobsCfgEntry.PlaceHolder = "http://vcs.domain.tld/project/repo/jenkins-jobs.git"

	// Cloud templates
	cloudTemplatesCheckBoxes := createCloudTemplates()
	cloudBox := createCloudTemplatesCheckboxes(cloudTemplatesCheckBoxes)

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
			{Text: "Domain name", Widget: jenkinsUrlEntry},
			{Text: "", Widget: jenkinsUrlErrorLabel},
			{Text: "Jenkins system message", Widget: jenkinsSysMsgEntry},
			{Text: "", Widget: jenkinsSysMsgErrorLabel},
			{Text: "Jenkins Jobs Repository", Widget: jenkinsJobsCfgEntry},
			{Text: "", Widget: jenkinsJobsCfgErrorLabel},
			{Text: "Jenkins Existing PVC", Widget: jenkinsExistingPvcEntry},
			{Text: "", Widget: jenkinsExistingPvcErrorLabel},
			{Text: "Cloud Templates", Widget: cloudBox},
		},
		OnSubmit: func() {
			// get variables
			projectConfig.Namespace = namespaceEntry.Text
			projectConfig.IPAddress = ipAddressEntry.Text
			projectConfig.JenkinsDomain = jenkinsUrlEntry.Text
			projectConfig.JenkinsSystemMsg = jenkinsSysMsgEntry.Text
			projectConfig.JobsCfgRepo = jenkinsJobsCfgEntry.Text
			projectConfig.ExistingPvc = jenkinsExistingPvcEntry.Text
			projectConfig.SelectedCloudTemplates = checkCloudboxes(cloudTemplatesCheckBoxes)
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

			// validate Jenkins domain
			if projectConfig.JenkinsDomain == "" && models.GetConfiguration().LoadBalancer.Annotations.ExtDNS.Hostname != "" {
				projectConfig.JenkinsDomain = projectConfig.Namespace + models.GetConfiguration().LoadBalancer.Annotations.ExtDNS.Hostname
			} else {
				err = validator.ValidateIP(projectConfig.JenkinsDomain)
				if err != nil {
					jenkinsUrlErrorLabel.SetText(err.Error())
					hasErrors = true
				} else {
					if strings.HasSuffix(projectConfig.JenkinsDomain, "/") {
						jenkinsUrlErrorLabel.SetText("URL should not end with a slash")
						hasErrors = true
					}
				}
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
					Bar:        widget.NewProgressBar(), //("Create project...", "Progress", window),
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

	return form
}

// ScreenCreateDeployOnlyProject creates the screen form for deployment only project without Jenkins
func ScreenCreateDeployOnlyProject(window fyne.Window) *widget.Form {
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

	// Domain for Jenkins
	jenkinsUrlErrorLabel := widget.NewLabel("")
	jenkinsUrlEntry := widget.NewEntry()
	jenkinsUrlEntry.PlaceHolder = "domain.tld (or leave empty to use <namespace>.<configured nginx domain>)"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Namespace", Widget: namespaceEntry},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "IP address", Widget: ipAddressEntry},
			{Text: "", Widget: ipAddressErrorLabel},
			{Text: "Domain name", Widget: jenkinsUrlEntry},
			{Text: "", Widget: jenkinsUrlErrorLabel},
		},
		OnSubmit: func() {
			// get variables
			projectConfig.Namespace = namespaceEntry.Text
			projectConfig.IPAddress = ipAddressEntry.Text
			projectConfig.JenkinsDomain = jenkinsUrlEntry.Text
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

			// validate Jenkins domain
			if projectConfig.JenkinsDomain == "" && models.GetConfiguration().LoadBalancer.Annotations.ExtDNS.Hostname != "" {
				projectConfig.JenkinsDomain = projectConfig.Namespace + models.GetConfiguration().LoadBalancer.Annotations.ExtDNS.Hostname
			} else {
				err = validator.ValidateIP(projectConfig.JenkinsDomain)
				if err != nil {
					jenkinsUrlErrorLabel.SetText(err.Error())
					hasError = true
				} else {
					if strings.HasSuffix(projectConfig.JenkinsDomain, "/") {
						jenkinsUrlErrorLabel.SetText("URL should not end with a slash")
						hasError = true
					}
				}
			}

			if !hasError {
				// process project creation
				bar := uielements.ProgressBar{
					Bar:        widget.NewProgressBar(), // ("Create project...", "Progress", window),
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

	return form
}

func createCloudTemplates() []*widget.Check {
	var cloudTemplates = createprojectactions.ActionReadCloudTemplates()
	var checkboxes []*widget.Check
	for _, cloudTemplate := range cloudTemplates {
		checkboxes = append(checkboxes, widget.NewCheck(cloudTemplate, func(set bool) {
			// not needed, because it ready it later from the options
		}))
	}

	return checkboxes
}

func createCloudTemplatesCheckboxes(boxes []*widget.Check) *container.Scroll {
	var box = container.NewVBox()
	// append boxes to VBox
	for _, checkbox := range boxes {
		box.Add(checkbox)
	}
	// pack them into a new VScrollContainer
	var content = container.NewVScroll(box)
	// set a min size, that it is possible to see more than 1
	content.SetMinSize(fyne.NewSize(-1, 150))

	return content
}

func checkCloudboxes(checkboxes []*widget.Check) []string {
	var checkedBoxes []string
	for _, checkbox := range checkboxes {
		if checkbox.Checked {
			checkedBoxes = append(checkedBoxes, checkbox.Text)
		}
	}

	return checkedBoxes
}

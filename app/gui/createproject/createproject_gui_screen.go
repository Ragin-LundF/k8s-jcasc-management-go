package createproject

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/utils/validator"
	"strings"
)

// ScreenCreateFullProject shows the full project setup screen form
func ScreenCreateFullProject(window fyne.Window) *widget.Form {
	var prj = project.NewProject()
	prj.Base.DeploymentOnly = false

	// Namespace
	var namespaceErrorLabel = widget.NewLabel("")
	var namespaceEntry = widget.NewEntry()
	namespaceEntry.PlaceHolder = "my-namespace"

	// IP address
	var ipAddressErrorLabel = widget.NewLabel("")
	var ipAddressEntry = widget.NewEntry()
	ipAddressEntry.PlaceHolder = "0.0.0.0 or mydomain.tld"

	// Domain for Jenkins
	var jenkinsUrlErrorLabel = widget.NewLabel("")
	var jenkinsUrlEntry = widget.NewEntry()
	jenkinsUrlEntry.PlaceHolder = "domain.tld (or leave empty to use <namespace>.<configured nginx domain>)"

	// Jenkins system message
	var jenkinsSysMsgErrorLabel = widget.NewLabel("")
	var jenkinsSysMsgEntry = widget.NewEntry()
	jenkinsSysMsgEntry.PlaceHolder = constants.CommonJenkinsSystemMessage

	// Jenkins jobs config repository
	var jenkinsJobsCfgErrorLabel = widget.NewLabel("")
	var jenkinsJobsCfgEntry = widget.NewEntry()
	jenkinsJobsCfgEntry.PlaceHolder = "http://vcs.domain.tld/project/repo/jenkins-jobs.git"

	// Cloud templates
	var cloudTemplatesCheckBoxes = createCloudTemplates()
	var cloudBox = createCloudTemplatesCheckboxes(cloudTemplatesCheckBoxes)

	// Existing PVC
	var jenkinsExistingPvcErrorLabel = widget.NewLabel("")
	var jenkinsExistingPvcEntry = widget.NewEntry()
	jenkinsExistingPvcEntry.PlaceHolder = "pvc-jenkins"

	// checkbox store config only
	var checkboxStoreTemplates = widget.NewCheck("", func(bool) {
		// the result is computed with checkboxStoreTemplates.Checked
	})
	checkboxStoreTemplates.SetChecked(true)

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Store only config", Widget: checkboxStoreTemplates},
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
			prj.SetNamespace(namespaceEntry.Text)
			prj.SetIPAddress(ipAddressEntry.Text)
			prj.SetDomain(jenkinsUrlEntry.Text)
			prj.SetJenkinsSystemMessage(jenkinsSysMsgEntry.Text)
			prj.SetJobsDefinitionRepository(jenkinsJobsCfgEntry.Text)
			prj.SetPersistentVolumeClaimExistingName(jenkinsExistingPvcEntry.Text)
			prj.SetCloudKubernetesAdditionalTemplateFiles(checkCloudboxes(cloudTemplatesCheckBoxes))
			prj.StoreConfigOnly = checkboxStoreTemplates.Checked
			hasErrors := false

			// validate namespace
			var err = validator.ValidateNewNamespace(prj.Base.Namespace)
			if err != nil {
				namespaceErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// validate IP address
			err = validator.ValidateIP(prj.Base.IPAddress)
			if err != nil {
				ipAddressErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// validate Jenkins domain
			if prj.Base.Domain == "" && configuration.GetConfiguration().Nginx.Loadbalancer.ExternalDNS.HostName != "" {
				prj.SetDomain(prj.Base.Namespace + configuration.GetConfiguration().Nginx.Loadbalancer.ExternalDNS.HostName)
			} else {
				err = validator.ValidateIP(prj.Base.Domain)
				if err != nil {
					jenkinsUrlErrorLabel.SetText(err.Error())
					hasErrors = true
				} else {
					if strings.HasSuffix(prj.Base.Domain, "/") {
						jenkinsUrlErrorLabel.SetText("URL should not end with a slash")
						hasErrors = true
					}
				}
			}

			// without annotations enabled an IP address is required
			if !configuration.GetConfiguration().Nginx.Loadbalancer.Annotations.Enabled && prj.Base.IPAddress == "" {
				ipAddressErrorLabel.SetText("If NGINX_LOADBALANCER_ANNOTATIONS_ENABLED is set to false, an IP address is required ")
				hasErrors = true
			}

			// validate jenkins system message
			err = validator.ValidateJenkinsSystemMessage(prj.JCasc.SystemMessage)
			if err != nil {
				jenkinsSysMsgErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// validate jobs config
			err = validator.ValidateJenkinsJobConfig(prj.JCasc.JobsConfig.JobsDefinitionRepository)
			if err != nil {
				jenkinsJobsCfgErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// validate PVC
			err = validator.ValidatePersistentVolumeClaim(prj.Base.ExistingVolumeClaim)
			if err != nil {
				jenkinsExistingPvcErrorLabel.SetText(err.Error())
				hasErrors = true
			}

			// process project creation if no error was found
			if !hasErrors {
				var bar = uielements.ProgressBar{
					Bar:        widget.NewProgressBar(), //("Create project...", "Progress", window),
					CurrentCnt: 0,
					MaxCount:   project.CountCreateProjectWorkflow,
				}
				bar.Bar.Show()
				_ = prj.ActionProcessProjectCreate(bar.AddCallback)
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
	var prj = project.NewProject()
	prj.Base.DeploymentOnly = true

	// Namespace
	var namespaceErrorLabel = widget.NewLabel("")
	var namespaceEntry = widget.NewEntry()
	namespaceEntry.PlaceHolder = "my-namespace"

	// IP address
	var ipAddressErrorLabel = widget.NewLabel("")
	var ipAddressEntry = widget.NewEntry()
	ipAddressEntry.PlaceHolder = "0.0.0.0"

	// Domain for Jenkins
	var jenkinsUrlErrorLabel = widget.NewLabel("")
	var jenkinsUrlEntry = widget.NewEntry()
	jenkinsUrlEntry.PlaceHolder = "domain.tld (or leave empty to use <namespace>.<configured nginx domain>)"

	// checkbox store config only
	var checkboxStoreTemplates = widget.NewCheck("", func(bool) {
		// the result is computed with checkboxStoreTemplates.Checked
	})
	checkboxStoreTemplates.SetChecked(true)

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Store only config", Widget: checkboxStoreTemplates},
			{Text: "Namespace", Widget: namespaceEntry},
			{Text: "", Widget: namespaceErrorLabel},
			{Text: "IP address", Widget: ipAddressEntry},
			{Text: "", Widget: ipAddressErrorLabel},
			{Text: "Domain name", Widget: jenkinsUrlEntry},
			{Text: "", Widget: jenkinsUrlErrorLabel},
		},
		OnSubmit: func() {
			// get variables
			prj.SetNamespace(namespaceEntry.Text)
			prj.SetIPAddress(ipAddressEntry.Text)
			prj.SetDomain(jenkinsUrlEntry.Text)
			prj.StoreConfigOnly = checkboxStoreTemplates.Checked
			hasError := false

			// validate namespace
			var err = validator.ValidateNewNamespace(prj.Base.Namespace)
			if err != nil {
				namespaceErrorLabel.SetText(err.Error())
				hasError = true
			}

			// validate IP address
			err = validator.ValidateIP(prj.Base.IPAddress)
			if err != nil {
				ipAddressErrorLabel.SetText(err.Error())
				hasError = true
			}

			// validate Jenkins domain
			if prj.Base.Domain == "" && configuration.GetConfiguration().Nginx.Loadbalancer.ExternalDNS.HostName != "" {
				prj.SetDomain(prj.Base.Namespace + configuration.GetConfiguration().Nginx.Loadbalancer.ExternalDNS.HostName)
			} else {
				err = validator.ValidateIP(prj.Base.Domain)
				if err != nil {
					jenkinsUrlErrorLabel.SetText(err.Error())
					hasError = true
				} else {
					if strings.HasSuffix(prj.Base.Domain, "/") {
						jenkinsUrlErrorLabel.SetText("URL should not end with a slash")
						hasError = true
					}
				}
			}

			if !hasError {
				// process project creation
				var bar = uielements.ProgressBar{
					Bar:        widget.NewProgressBar(), // ("Create project...", "Progress", window),
					CurrentCnt: 0,
					MaxCount:   project.CountCreateProjectWorkflow,
				}
				bar.Bar.Show()
				_ = prj.ActionProcessProjectCreate(bar.AddCallback)
				bar.Bar.Hide()

				// show output
				uielements.ShowLogOutput(window)
			}
		},
	}

	return form
}

func createCloudTemplates() []*widget.Check {
	var cloudTemplates = project.ActionReadCloudTemplates()
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

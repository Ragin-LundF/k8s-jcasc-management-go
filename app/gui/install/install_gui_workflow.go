package install

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/gui/uielements"
	"k8s-management-go/app/utils/loggingstate"
	"time"
)

// ExecuteInstallWorkflow executes the install  workflow
func ExecuteInstallWorkflow(window fyne.Window, projectConfig install.ProjectConfig) (err error) {
	// Progress Bar
	var progressCnt = 1
	var progressMaxCnt = projectConfig.CalculateBarCounter()
	var bar = dialog.NewProgress(
		projectConfig.HelmCommand,
		"Installing on namespace "+projectConfig.Project.Base.Namespace,
		window)
	bar.Show()

	// it is not a dry-run -> install required stuff
	if !configuration.GetConfiguration().K8SManagement.DryRunOnly {
		// check if namespace is available or create a new one if not
		err = namespaceactions.ProcessNamespaceCreation(projectConfig)
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++
		if err != nil {
			bar.Hide()
			uielements.ShowLogOutput(window)
			return err
		}

		// check if PVC was specified and install it if needed
		err = projectConfig.ProcessCheckAndCreatePvc()
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++
		if err != nil {
			bar.Hide()
			uielements.ShowLogOutput(window)
			return err
		}

		// Jenkins exists and it is not a dry-run install secrets
		if projectConfig.JenkinsHelmValuesExist {
			// apply secrets
			err = projectConfig.ProcessCreateSecrets()
			bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
			progressCnt++
			if err != nil {
				bar.Hide()
				uielements.ShowLogOutput(window)
				return err
			}
		}
	} else {
		loggingstate.AddInfoEntry("-> Dry run. Skipping namespace creation, pvc installation and secrets apply...")
	}

	// install Jenkins
	err = projectConfig.ProcessInstallJenkins()
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		uielements.ShowLogOutput(window)
		return err
	}

	// install Nginx ingress controller
	err = projectConfig.ProcessNginxController()
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		uielements.ShowLogOutput(window)
		return err
	}

	// last but not least execute install scripts if it is not dry-run only
	err = projectConfig.ProcessScripts()
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	time.Sleep(time.Duration(1) * time.Second)
	bar.Hide()

	uielements.ShowLogOutput(window)

	return err
}

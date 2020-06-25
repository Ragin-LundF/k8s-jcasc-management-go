package install

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"time"
)

// execute the workflow
func ExecuteInstallWorkflow(window fyne.Window, state install.StateData) (err error) {
	log := logger.Log()

	// Progress Bar
	progressCnt := 1
	progressMaxCnt := install.CalculateBarCounter(state)
	bar := dialog.NewProgress(state.HelmCommand, "Installing on namespace "+state.Namespace, window)
	bar.Show()
	//	bar.Max = float64(install.CalculateBarCounter(state))
	// createProgressBar(window, bar, &*textDialogProgress)

	// it is not a dry-run -> install required stuff
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		// check if namespace is available or create a new one if not
		err = install.ProgressNamespace(state)
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++
		if err != nil {
			bar.Hide()
			return err
		}

		// check if PVC was specified and install it if needed
		err = install.ProgressPvc(state)
		bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
		progressCnt++
		if err != nil {
			bar.Hide()
			return err
		}

		// Jenkins exists and it is not a dry-run install secrets
		if state.JenkinsHelmValuesExist {
			// apply secrets
			err = install.ProgressSecrets(state)
			bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
			progressCnt++
			if err != nil {
				bar.Hide()
				return err
			}
		}
	} else {
		loggingstate.AddInfoEntry("-> Dry run. Skipping namespace creation, pvc installation and secrets apply...")
		log.Infof("[DoUpgradeOrInstall] Dry run only, skipping namespace [%s] creation, pvc installation and secrets apply...", state.Namespace)
	}

	// install Jenkins
	err = install.ProgressJenkins(state.HelmCommand, state)
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		return err
	}

	// install Nginx ingress controller
	err = install.ProgressNginxController(state.HelmCommand, state)
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	progressCnt++
	if err != nil {
		bar.Hide()
		log.Errorf("[DoUpgradeOrInstall] Unable to install nginx-ingress-controller.\n%s", err.Error())
		return err
	}

	// last but not least execute install scripts if it is not dry-run only
	err = install.ProgressScripts(state)
	bar.SetValue(float64(1) / float64(progressMaxCnt) * float64(progressCnt))
	time.Sleep(1000)
	bar.Hide()

	return err
}

type progressDescription struct {
	Text string
}

func createProgressBar(window fyne.Window, bar *widget.ProgressBar, text *string) {
	widget.NewVBox(
		widget.NewLabel(*text),
		bar,
	)
	info := dialog.NewCustom("Progress", "Ok", bar, window)
	info.Show()
}

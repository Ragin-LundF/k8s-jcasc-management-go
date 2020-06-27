package install

import (
	"fmt"
	"k8s-management-go/app/actions/install_actions"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

/*
Workflow for Jenkins installation

if dry-run only:
- Install Jenkins
- Install Nginx Ingress Controller

if ! dry-run only && ! jenkins installation:
- Namespace check & creation
- Install PVC
- Install Nginx Ingress Controller
- Install Scripts

if ! dry-run only && jenkins installation
- Namespace check & creation
- Install PVC
- Apply Secrets
- Install Jenkins
- Install Nginx Ingress Controller
- Install Scripts
*/
func DoUpgradeOrInstall(helmCommand string) (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf("Starting Jenkins [%s]...", helmCommand))
	// show all needed ui_elements and collect data
	state, err := ShowDialogs()
	if err != nil {
		return err
	}

	// set command to the state if ui_elements was successful
	state.HelmCommand = helmCommand

	// execute workflow
	err = executeWorkflow(state)

	loggingstate.AddInfoEntry(fmt.Sprintf("Starting Jenkins [%s]...done", helmCommand))
	return err
}

// execute the workflow
func executeWorkflow(state models.StateData) (err error) {
	log := logger.Log()

	// Progress Bar
	bar := dialogs.CreateProgressBar("Installing...", install_actions.CalculateBarCounter(state))

	// it is not a dry-run -> install_actions required stuff
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		// check if namespace is available or create a new one if not
		bar.Describe("Check and create namespace if necessary...")
		err = install_actions.ProgressNamespace(state)
		_ = bar.Add(1)
		if err != nil {
			return err
		}

		// check if PVC was specified and install_actions it if needed
		bar.Describe("Check and create PVC if necessary...")
		err = install_actions.ProgressPvc(state)
		_ = bar.Add(1)
		if err != nil {
			return err
		}

		// Jenkins exists and it is not a dry-run install_actions secrets
		if state.JenkinsHelmValuesExist {
			// apply secrets
			bar.Describe("Applying secrets...")
			err = install_actions.ProgressSecrets(state)
			_ = bar.Add(1)
			if err != nil {
				return err
			}
		}
	} else {
		loggingstate.AddInfoEntry("-> Dry run. Skipping namespace creation, pvc installation and secrets apply...")
		log.Infof("[DoUpgradeOrInstall] Dry run only, skipping namespace [%s] creation, pvc installation and secrets apply...", state.Namespace)
	}

	// install_actions Jenkins
	bar.Describe("Installing Jenkins...")
	err = install_actions.ProgressJenkins(state.HelmCommand, state)
	_ = bar.Add(1)
	if err != nil {
		return err
	}

	// install_actions Nginx ingress controller
	bar.Describe("Installing nginx-ingress-controller...")
	err = install_actions.ProgressNginxController(state.HelmCommand, state)
	_ = bar.Add(1)
	if err != nil {
		log.Errorf("[DoUpgradeOrInstall] Unable to install_actions nginx-ingress-controller.\n%s", err.Error())
		return err
	}

	// last but not least execute install_actions scripts if it is not dry-run only
	bar.Describe("Check and execute additional scripts...")
	err = install_actions.ProgressScripts(state)
	_ = bar.Add(1)

	return err
}

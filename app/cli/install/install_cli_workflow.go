package install

import (
	"fmt"
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/actions/namespaceactions"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

/*
DoUpgradeOrInstall is the workflow for Jenkins installation

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
	// show all needed uielements and collect data
	projectConfig, err := ShowInstallDialogs()
	if err != nil {
		return err
	}

	// set command to the state if uielements was successful
	projectConfig.HelmCommand = helmCommand

	// execute workflow
	err = executeWorkflow(projectConfig)

	loggingstate.AddInfoEntry(fmt.Sprintf("Starting Jenkins [%s]...done", helmCommand))
	return err
}

// execute the workflow
func executeWorkflow(projectConfig install.ProjectConfig) (err error) {
	var log = logger.Log()

	// Progress Bar
	var bar = dialogs.CreateProgressBar("Installing...", projectConfig.CalculateBarCounter())

	// it is not a dry-run -> install required stuff
	if !configuration.GetConfiguration().K8SManagement.DryRunOnly {
		// check if namespace is available or create a new one if not
		bar.Describe("Check and create namespace if necessary...")
		err = namespaceactions.ProcessNamespaceCreation(projectConfig)
		_ = bar.Add(1)
		if err != nil {
			return err
		}

		// check if PVC was specified and install it if needed
		bar.Describe("Check and create PVC if necessary...")
		err = projectConfig.ProcessCheckAndCreatePvc()
		_ = bar.Add(1)
		if err != nil {
			return err
		}

		// Jenkins exists and it is not a dry-run install secrets
		if projectConfig.Project.CalculateIfDeploymentFileIsRequired(constants.FilenameJenkinsHelmValues) {
			// apply secrets
			bar.Describe("Applying secrets...")
			err = projectConfig.ProcessCreateSecrets()
			_ = bar.Add(1)
			if err != nil {
				return err
			}
		}
	} else {
		loggingstate.AddInfoEntry("-> Dry run. Skipping namespace creation, pvc installation and secrets apply...")
		log.Infof(
			"[DoUpgradeOrInstall] Dry run only, skipping namespace [%s] creation, pvc installation and secrets apply...",
			projectConfig.Project.Base.Namespace)
	}

	// install Jenkins
	bar.Describe("Installing Jenkins...")
	err = projectConfig.ProcessInstallJenkins()
	_ = bar.Add(1)
	if err != nil {
		return err
	}

	// install Nginx ingress controller
	bar.Describe("Installing nginx-ingress-controller...")
	err = projectConfig.ProcessNginxController()
	_ = bar.Add(1)
	if err != nil {
		log.Errorf("[DoUpgradeOrInstall] Unable to install nginx-ingress-controller.\n%s", err.Error())
		return err
	}

	// last but not least execute install of the scripts if it is not dry-run only
	bar.Describe("Check and execute additional scripts...")
	err = projectConfig.ProcessScripts()
	_ = bar.Add(1)

	return err
}

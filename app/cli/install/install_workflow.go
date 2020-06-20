package install

import (
	"errors"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
)

type stateData struct {
	ProjectPath            string
	Namespace              string
	DeploymentName         string
	JenkinsHelmValuesFile  string
	JenkinsHelmValuesExist bool
	SecretsPassword        *string
}

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
	log := logger.Log()
	loggingstate.AddInfoEntry(fmt.Sprintf("Starting Jenkins [%s]...", helmCommand))
	// show all needed dialogs and collect data
	state, err := showDialogs()
	if err != nil {
		return err
	}

	// Progress Bar
	bar := dialogs.CreateProgressBar("Installing...", calculateBarCounter(state))

	// it is not a dry-run -> install required stuff
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		// check if namespace is available or create a new one if not
		err = progressNamespace(state, &bar)
		if err != nil {
			return err
		}

		// check if PVC was specified and install it if needed
		err = progressPvc(state, &bar)
		if err != nil {
			return err
		}

		// Jenkins exists and it is not a dry-run install secrets
		if state.JenkinsHelmValuesExist {
			// apply secrets
			err = progressSecrets(state, &bar)
			if err != nil {
				return err
			}
		}
	} else {
		loggingstate.AddInfoEntry("-> Dry run. Skipping namespace creation, pvc installation and secrets apply...")
		log.Infof("[DoUpgradeOrInstall] Dry run only, skipping namespace [%s] creation, pvc installation and secrets apply...", state.Namespace)
	}

	// install Jenkins
	err = progressJenkins(helmCommand, state, &bar)
	if err != nil {
		return err
	}

	// install Nginx ingress controller
	err = progressNginxController(helmCommand, state, &bar)
	if err != nil {
		log.Errorf("[DoUpgradeOrInstall] Unable to install nginx-ingress-controller.\n%s", err.Error())
		return err
	}

	// last but not least execute install scripts if it is not dry-run only
	err = progressScripts(state, &bar)

	loggingstate.AddInfoEntry(fmt.Sprintf("Starting Jenkins [%s]...done", helmCommand))
	return err
}

func showDialogs() (state stateData, err error) {
	// ask for namespace
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	state.Namespace, err = dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Ask for namespace...failed.", err.Error())
		return state, err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// first check if namespace directory exists
	loggingstate.AddInfoEntry("-> Checking existing directories...")
	state.ProjectPath = files.AppendPath(
		models.GetProjectBaseDirectory(),
		state.Namespace,
	)
	// validate that project is existing
	if !files.FileOrDirectoryExists(state.ProjectPath) {
		err = errors.New(fmt.Sprintf("Project directory not found: [%s]", state.ProjectPath))
		loggingstate.AddErrorEntryAndDetails("-> Checking existing directories...failed", err.Error())
		return state, err
	}
	loggingstate.AddInfoEntry("-> Checking existing directories...done")

	// check if project configuration contains Jenkins Helm values file
	state.JenkinsHelmValuesFile = files.AppendPath(
		state.ProjectPath,
		constants.FilenameJenkinsHelmValues,
	)
	state.JenkinsHelmValuesExist = files.FileOrDirectoryExists(state.JenkinsHelmValuesFile)

	// if it is Jenkins installation ask more things
	if state.JenkinsHelmValuesExist {
		// if it is no dry-run, ask for secrets password
		if !models.GetConfiguration().K8sManagement.DryRunOnly {
			secretsPassword, err := secrets.AskForSecretsPassword("Password for secrets file")
			state.SecretsPassword = &secretsPassword
			if err != nil {
				return state, err
			}
		}

		// ask for deployment name if necessary
		state.DeploymentName, err = dialogs.DialogAskForDeploymentName("Deployment name", nil)
		if err != nil {
			log := logger.Log()
			loggingstate.AddErrorEntryAndDetails("  -> Unable to get deployment name.", err.Error())
			log.Errorf("[DoUpgradeOrInstall] Unable to get deployment name.\n%s", err.Error())
			return state, err
		}
	}
	return state, err
}

func progressNamespace(state stateData, bar *progressbar.ProgressBar) (err error) {
	bar.Describe("Check and create namespace if necessary...")

	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...")
	if err = CheckAndCreateNamespace(state.Namespace); err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Check and create namespace if necessary...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...done")
	bar.Add(1)

	return nil
}

func progressPvc(state stateData, bar *progressbar.ProgressBar) (err error) {
	bar.Describe("Check and create PVC if necessary...")

	loggingstate.AddInfoEntry("-> Check and create pvc if necessary...")
	if err = PersistenceVolumeClaimInstall(state.Namespace); err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Check and create pvc if necessary...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Check and create pvc if necessary...done")
	bar.Add(1)

	return nil
}

func progressSecrets(state stateData, bar *progressbar.ProgressBar) (err error) {
	log := logger.Log()
	// apply secrets
	bar.Describe("Applying secrets...")
	log.Infof("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...", state.Namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...", state.Namespace))

	if err = secrets.ApplySecretsToNamespace(state.Namespace, state.SecretsPassword); err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...failed", state.Namespace), err.Error())
		log.Errorf("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...failed\n%s", err.Error())
		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...done", state.Namespace))
	log.Infof("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...done", state.Namespace)
	bar.Add(1)

	return nil
}

func progressJenkins(helmCommand string, state stateData, bar *progressbar.ProgressBar) (err error) {
	log := logger.Log()
	// install Jenkins
	if state.JenkinsHelmValuesExist {
		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...")
		log.Infof("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...")

		bar.Describe("Installing Jenkins...")
		err = HelmInstallJenkins(helmCommand, state.Namespace, state.DeploymentName)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Jenkins Helm values.yaml found. Installing Jenkins...failed", err.Error())
			log.Errorf("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...failed\n%s", err.Error())
			return err
		}

		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...done")
		log.Infof("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...done")
		bar.Add(1)
	} else {
		loggingstate.AddInfoEntry(fmt.Sprintf("-> No Jenkins Helm chart found in path [%s]. Skipping installation...", state.JenkinsHelmValuesFile))
		log.Infof("No Jenkins Helm chart found in path [%s]. Skipping Jenkins installation.", state.JenkinsHelmValuesFile)
	}

	return nil
}

func progressNginxController(helmCommand string, state stateData, bar *progressbar.ProgressBar) (err error) {
	// install Nginx ingress controller
	bar.Describe("Installing nginx-ingress-controller...")
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Installing nginx-ingress-controller on namespace [%s]...", state.Namespace))
	err = HelmInstallNginxIngressController(helmCommand, state.Namespace, state.JenkinsHelmValuesExist)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to install nginx-ingress-controller.", err.Error())
		return err
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Installing nginx-ingress-controller on namespace [%s]...done", state.Namespace))
	bar.Add(1)

	return nil
}

func progressScripts(state stateData, bar *progressbar.ProgressBar) (err error) {
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		bar.Describe("Check and execute additional scripts...")
		// install scripts
		// try to install scripts
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute install scripts on [%s]...", state.Namespace))
		// we ignore errors. They will be logged, but we keep on doing the install for the scripts
		_ = ShellScriptsInstall(state.Namespace)
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute install scripts on [%s]...done", state.Namespace))
	}
	bar.Add(1)

	return nil
}

func calculateBarCounter(state stateData) int {
	var dryRunOnly = 0
	var notDryRunOnly = 0
	var jenkinsInstallation = 0
	if models.GetConfiguration().K8sManagement.DryRunOnly {
		// only dry-run
		dryRunOnly = 2
	} else {
		notDryRunOnly = 4
		if state.JenkinsHelmValuesExist {
			jenkinsInstallation = 2
		}
	}
	return dryRunOnly + notDryRunOnly + jenkinsInstallation
}

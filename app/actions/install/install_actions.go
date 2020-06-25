package install

import (
	"errors"
	"fmt"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
)

type StateData struct {
	ProjectPath            string
	Namespace              string
	DeploymentName         string
	JenkinsHelmValuesFile  string
	JenkinsHelmValuesExist bool
	SecretsPassword        *string
	HelmCommand            string
}

func ProgressNamespace(state StateData) (err error) {
	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...")
	if err = CheckAndCreateNamespace(state.Namespace); err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Check and create namespace if necessary...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...done")

	return nil
}

func ProgressPvc(state StateData) (err error) {
	loggingstate.AddInfoEntry("-> Check and create pvc if necessary...")
	if err = PersistenceVolumeClaimInstall(state.Namespace); err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Check and create pvc if necessary...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Check and create pvc if necessary...done")

	return nil
}

func ProgressSecrets(state StateData) (err error) {
	log := logger.Log()
	// apply secrets
	log.Infof("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...", state.Namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...", state.Namespace))

	if err = secrets.ApplySecretsToNamespace(state.Namespace, state.SecretsPassword); err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...failed", state.Namespace), err.Error())
		log.Errorf("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...failed\n%s", err.Error())
		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...done", state.Namespace))
	log.Infof("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...done", state.Namespace)

	return nil
}

func ProgressJenkins(helmCommand string, state StateData) (err error) {
	log := logger.Log()
	// install Jenkins
	if state.JenkinsHelmValuesExist {
		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...")
		log.Infof("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...")

		err = HelmInstallJenkins(helmCommand, state.Namespace, state.DeploymentName)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Jenkins Helm values.yaml found. Installing Jenkins...failed", err.Error())
			log.Errorf("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...failed\n%s", err.Error())
			return err
		}

		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...done")
		log.Infof("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...done")
	} else {
		loggingstate.AddInfoEntry(fmt.Sprintf("-> No Jenkins Helm chart found in path [%s]. Skipping installation...", state.JenkinsHelmValuesFile))
		log.Infof("No Jenkins Helm chart found in path [%s]. Skipping Jenkins installation.", state.JenkinsHelmValuesFile)
	}

	return nil
}

func ProgressNginxController(helmCommand string, state StateData) (err error) {
	// install Nginx ingress controller
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Installing nginx-ingress-controller on namespace [%s]...", state.Namespace))
	err = HelmInstallNginxIngressController(helmCommand, state.Namespace, state.JenkinsHelmValuesExist)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to install nginx-ingress-controller.", err.Error())
		return err
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Installing nginx-ingress-controller on namespace [%s]...done", state.Namespace))

	return nil
}

func ProgressScripts(state StateData) (err error) {
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		// install scripts
		// try to install scripts
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute install scripts on [%s]...", state.Namespace))
		// we ignore errors. They will be logged, but we keep on doing the install for the scripts
		_ = ShellScriptsInstall(state.Namespace)
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute install scripts on [%s]...done", state.Namespace))
	}

	return nil
}

// calculate bar counter
func CalculateBarCounter(state StateData) int {
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

func CalculateDirectoriesForInstall(state StateData, namespace string) (err error, stateResult StateData) {
	// first check if namespace directory exists
	loggingstate.AddInfoEntry("-> Checking existing directories...")
	state.Namespace = namespace
	state.ProjectPath = files.AppendPath(
		models.GetProjectBaseDirectory(),
		namespace,
	)
	// validate that project is existing
	if !files.FileOrDirectoryExists(state.ProjectPath) {
		err = errors.New(fmt.Sprintf("Project directory not found: [%s]", state.ProjectPath))
		loggingstate.AddErrorEntryAndDetails("-> Checking existing directories...failed", err.Error())
		return err, state
	}
	loggingstate.AddInfoEntry("-> Checking existing directories...done")
	return err, state
}

func CheckJenkinsDirectories(state StateData) StateData {
	// check if project configuration contains Jenkins Helm values file
	state.JenkinsHelmValuesFile = files.AppendPath(
		state.ProjectPath,
		constants.FilenameJenkinsHelmValues,
	)
	state.JenkinsHelmValuesExist = files.FileOrDirectoryExists(state.JenkinsHelmValuesFile)

	return state
}

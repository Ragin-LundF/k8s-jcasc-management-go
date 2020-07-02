package installactions

import (
	"fmt"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
)

// ProcessCheckAndCreatePvc checks for existing PVC and creates new one if it does not exist
func ProcessCheckAndCreatePvc(state models.StateData) (err error) {
	loggingstate.AddInfoEntry("-> Check and create pvc if necessary...")
	if err = ActionPersistenceVolumeClaimInstall(state.Namespace); err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Check and create pvc if necessary...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Check and create pvc if necessary...done")

	return nil
}

// ProcessCreateSecrets executes the secrets script
func ProcessCreateSecrets(state models.StateData) (err error) {
	// apply secrets
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...", state.Namespace))

	if err = secrets.ApplySecretsToNamespace(state.Namespace, state.SecretsPassword); err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...failed", state.Namespace), err.Error())
		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...done", state.Namespace))

	return nil
}

// ProcessInstallJenkins processes the Jenkins master installation
func ProcessInstallJenkins(helmCommand string, state models.StateData) (err error) {
	// install Jenkins
	if state.JenkinsHelmValuesExist {
		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...")

		err = ActionHelmInstallJenkins(helmCommand, state.Namespace, state.DeploymentName)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Jenkins Helm values.yaml found. Installing Jenkins...failed", err.Error())
			return err
		}

		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...done")
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("-> No Jenkins Helm chart found in path [%s]. Skipping installation...", state.JenkinsHelmValuesFile))

	return nil
}

// ProcessNginxController installs the Nginx Ingress controller
func ProcessNginxController(helmCommand string, state models.StateData) (err error) {
	// install Nginx ingress controller
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Installing nginx-ingress-controller on namespace [%s]...", state.Namespace))
	err = ActionHelmInstallNginxIngressController(helmCommand, state.Namespace, state.JenkinsHelmValuesExist)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to install nginx-ingress-controller.", err.Error())
		return err
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Installing nginx-ingress-controller on namespace [%s]...done", state.Namespace))

	return nil
}

// ProcessScripts processes the scripts execution
func ProcessScripts(state models.StateData) (err error) {
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		// install scripts
		// try to install scripts
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute install scripts on [%s]...", state.Namespace))
		// we ignore errors. They will be logged, but we keep on doing the installation of the scripts
		_ = ActionShellScriptsInstall(state.Namespace)
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute install scripts on [%s]...done", state.Namespace))
	}

	return nil
}

// CalculateBarCounter calculates bar counter
func CalculateBarCounter(state models.StateData) int {
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

// CalculateDirectoriesForInstall calculates the directory for installation
func CalculateDirectoriesForInstall(state models.StateData, namespace string) (stateResult models.StateData, err error) {
	// first check if namespace directory exists
	loggingstate.AddInfoEntry("-> Checking existing directories...")
	state.Namespace = namespace
	state.ProjectPath = files.AppendPath(
		models.GetProjectBaseDirectory(),
		namespace,
	)
	// validate that project is existing
	if !files.FileOrDirectoryExists(state.ProjectPath) {
		err = fmt.Errorf("Project directory not found: [%s] ", state.ProjectPath)
		loggingstate.AddErrorEntryAndDetails("-> Checking existing directories...failed", err.Error())
	}
	loggingstate.AddInfoEntry("-> Checking existing directories...done")
	return state, err
}

// CheckJenkinsDirectories checks if Jenkins Helm values file exists
func CheckJenkinsDirectories(state models.StateData) models.StateData {
	// check if project configuration contains Jenkins Helm values file
	state.JenkinsHelmValuesFile = files.AppendPath(
		state.ProjectPath,
		constants.FilenameJenkinsHelmValues,
	)
	state.JenkinsHelmValuesExist = files.FileOrDirectoryExists(state.JenkinsHelmValuesFile)

	return state
}

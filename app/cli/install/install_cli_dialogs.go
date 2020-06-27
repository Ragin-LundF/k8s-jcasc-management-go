package install

import (
	"k8s-management-go/app/actions/install_actions"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

// show CLI ui_elements
func ShowDialogs() (state models.StateData, err error) {
	// ask for namespace
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	state.Namespace, err = dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Ask for namespace...failed.", err.Error())
		return state, err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// Directories
	err, state = install_actions.CalculateDirectoriesForInstall(state, state.Namespace)
	if err != nil {
		return state, err
	}

	// check if project configuration contains Jenkins Helm values file
	state = install_actions.CheckJenkinsDirectories(state)

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

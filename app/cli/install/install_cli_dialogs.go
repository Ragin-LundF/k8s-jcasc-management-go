package install

import (
	"k8s-management-go/app/actions/installactions"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

// ShowInstallDialogs shows CLI ui elements
func ShowInstallDialogs() (state models.StateData, err error) {
	// ask for namespace
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	state.Namespace, err = dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Ask for namespace...failed.", err.Error())
		return state, err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// Directories
	state, err = installactions.CalculateDirectoriesForInstall(state, state.Namespace)
	if err != nil {
		return state, err
	}

	// check if project configuration contains Jenkins Helm values file
	state = installactions.CheckJenkinsDirectories(state)

	// if it is Jenkins installation ask more things
	if state.JenkinsHelmValuesExist {
		// if it is no dry-run, ask for secrets password
		if !models.GetConfiguration().K8sManagement.DryRunOnly {
			secretsFileName, secretsPassword, err := secrets.AskForSecretsPassword("Password for secrets file", true)
			state.SecretsPassword = &secretsPassword
			state.SecretsFileName = secretsFileName
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

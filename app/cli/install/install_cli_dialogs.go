package install

import (
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

// ShowInstallDialogs shows CLI ui elements
func ShowInstallDialogs() (projectConfig install.ProjectConfig, err error) {
	// ask for namespace
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Ask for namespace...failed.", err.Error())
		return projectConfig, err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// instantiate a new project config
	projectConfig = install.NewInstallProjectConfig()
	err = projectConfig.LoadProjectConfigIfExists(namespace)
	if err != nil {
		return projectConfig, err
	}
	projectConfig.Project.SetNamespace(namespace)

	// Directories
	err = projectConfig.CalculateDirectoriesForInstall()
	if err != nil {
		return projectConfig, err
	}

	// if it is Jenkins installation ask more things
	if projectConfig.Project.CalculateIfDeploymentFileIsRequired(constants.FilenameJenkinsHelmValues) {
		// if it is no dry-run, ask for secrets password
		if !configuration.GetConfiguration().K8SManagement.DryRunOnly {
			secretsFileName, secretsPassword, err := secrets.AskForSecretsPassword("Password for secrets file", true)
			projectConfig.SecretsPassword = &secretsPassword
			projectConfig.SecretsFileName = secretsFileName
			if err != nil {
				return projectConfig, err
			}
		}

		// ask for deployment name if necessary
		deploymentName, err := dialogs.DialogAskForDeploymentName("Deployment name", nil)
		if err != nil {
			log := logger.Log()
			loggingstate.AddErrorEntryAndDetails("  -> Unable to get deployment name.", err.Error())
			log.Errorf("[DoUpgradeOrInstall] Unable to get deployment name.\n%s", err.Error())
			return projectConfig, err
		}
		projectConfig.Project.Base.DeploymentName = deploymentName
	}
	return projectConfig, err
}

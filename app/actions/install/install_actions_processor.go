package install

import (
	"fmt"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
)

// ProcessCheckAndCreatePvc checks for existing PVC and creates new one if it does not exist
func (projectConfig *ProjectConfig) ProcessCheckAndCreatePvc() (err error) {
	if err = projectConfig.ActionPersistenceVolumeClaimInstall(); err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Check and create pvc if necessary...failed", err.Error())
		return err
	}

	return nil
}

// ProcessCreateSecrets executes the secrets script
func (projectConfig *ProjectConfig) ProcessCreateSecrets() (err error) {
	// apply secrets
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Starting apply secrets to namespace [%s]...",
		projectConfig.Project.Base.Namespace))

	if err = secrets.ApplySecretsToNamespace(projectConfig.Project.Base.Namespace, projectConfig.SecretsFileName, projectConfig.SecretsPassword); err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
			"  -> Starting apply secrets to namespace [%s]...failed",
			projectConfig.Project.Base.Namespace), err.Error())
		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Starting apply secrets to namespace [%s]...done",
		projectConfig.Project.Base.Namespace))

	return nil
}

// ProcessInstallJenkins processes the Jenkins master installation
func (projectConfig *ProjectConfig) ProcessInstallJenkins() (err error) {
	// install Jenkins
	if projectConfig.Project.CalculateIfDeploymentFileIsRequired(constants.FilenameJenkinsHelmValues) {
		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...")

		err = projectConfig.ActionHelmInstallJenkins()
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Jenkins Helm values.yaml found. Installing Jenkins...failed", err.Error())
			return err
		}

		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...done")
	}
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"-> No Jenkins Helm chart found in path [%s]. Skipping installation...",
		projectConfig.ProjectPath))

	return nil
}

// ProcessNginxController installs the Nginx Ingress controller
func (projectConfig *ProjectConfig) ProcessNginxController() (err error) {
	// install Nginx ingress controller
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"-> Installing nginx-ingress-controller on namespace [%s]...",
		projectConfig.Project.Base.Namespace))
	err = projectConfig.ActionHelmInstallNginxIngressController()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to install nginx-ingress-controller.", err.Error())
		return err
	}
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"-> Installing nginx-ingress-controller on namespace [%s]...done",
		projectConfig.Project.Base.Namespace))

	return nil
}

// ProcessScripts processes the scripts execution
func (projectConfig *ProjectConfig) ProcessScripts() (err error) {
	if !configuration.GetConfiguration().K8SManagement.DryRunOnly {
		// install scripts
		// try to install scripts
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"-> Try to execute install scripts on [%s]...",
			projectConfig.Project.Base.Namespace))
		// we ignore errors. They will be logged, but we keep on doing the installation of the scripts
		_ = projectConfig.ActionShellScriptsInstall()
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"-> Try to execute install scripts on [%s]...done",
			projectConfig.Project.Base.Namespace))
	}

	return nil
}

// CalculateBarCounter calculates bar counter
func (projectConfig *ProjectConfig) CalculateBarCounter() int {
	var dryRunOnly = 0
	var notDryRunOnly = 0
	var jenkinsInstallation = 0
	if configuration.GetConfiguration().K8SManagement.DryRunOnly {
		// only dry-run
		dryRunOnly = 2
	} else {
		notDryRunOnly = 4
		if projectConfig.Project.CalculateIfDeploymentFileIsRequired(constants.FilenameJenkinsHelmValues) {
			jenkinsInstallation = 2
		}
	}
	return dryRunOnly + notDryRunOnly + jenkinsInstallation
}

// CalculateDirectoriesForInstall calculates the directory for installation
func (projectConfig *ProjectConfig) CalculateDirectoriesForInstall() (err error) {
	// first check if namespace directory exists
	loggingstate.AddInfoEntry("-> Checking existing directories...")
	projectConfig.ProjectPath = files.AppendPath(
		configuration.GetConfiguration().GetProjectBaseDirectory(),
		projectConfig.Project.Base.Namespace,
	)
	// validate that project is existing
	if !files.FileOrDirectoryExists(projectConfig.ProjectPath) {
		err = fmt.Errorf("Project directory not found: [%s] ", projectConfig.ProjectPath)
		loggingstate.AddErrorEntryAndDetails("-> Checking existing directories...failed", err.Error())
	}
	loggingstate.AddInfoEntry("-> Checking existing directories...done")
	return err
}

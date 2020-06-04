package install

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/files"
)

// workflow for Jenkins installation
func DoUpgradeOrInstall(helmCommand string) (info string, err error) {
	// ask for namespace
	var infoLog string
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		return info, err
	}

	// first check if namespace directory exists
	projectPath := files.AppendPath(
		config.GetProjectBaseDirectory(),
		namespace,
	)
	if !files.FileOrDirectoryExists(projectPath) {
		return info, errors.New("Project directory not found: [" + projectPath + "]")
	}

	// create namespace and pvc only, if it is not dry-run only
	if !config.GetConfiguration().K8sManagement.DryRunOnly {
		// check if namespace is available or create a new one if not
		infoLog, err = CheckAndCreateNamespace(namespace)
		info = info + constants.NewLine + infoLog
		if err != nil {
			return info, err
		}

		// check if PVC was specified and install it if needed
		infoLog, err = PersistenceVolumeClaimInstall(namespace)
		info = info + constants.NewLine + infoLog
		if err != nil {
			return info, err
		}
	}

	// check if project configuration contains Jenkins Helm values file
	jenkinsHelmValuesFile := files.AppendPath(
		projectPath,
		constants.FilenameJenkinsHelmValues,
	)
	jenkinsHelmValuesExist := files.FileOrDirectoryExists(jenkinsHelmValuesFile)

	// apply secrets only if Jenkins Helm values are existing
	if jenkinsHelmValuesExist {
		// apply secrets only, if it is not dry-run only
		if !config.GetConfiguration().K8sManagement.DryRunOnly {
			// apply secrets
			infoLog, err = secrets.ApplySecretsToNamespace(namespace)
			info = info + constants.NewLine + infoLog
			if err != nil {
				return info, err
			}
		}

		// ask for deployment name
		deploymentName, err := dialogs.DialogAskForDeploymentName("Deployment name", nil)
		if err != nil {
			return info, err
		}

		// install Jenkins
		infoLog, err := HelmInstallJenkins(helmCommand, namespace, deploymentName)
		info = info + constants.NewLine + infoLog
		if err != nil {
			return info, err
		}
	} else {
		info = info + constants.NewLine + "No Jenkins Helm chart found in path [" + jenkinsHelmValuesFile + "]."
	}

	// install Nginx ingress controller
	infoLog, err = HelmInstallNginxIngressController(helmCommand, namespace, jenkinsHelmValuesExist)
	info = info + constants.NewLine + infoLog
	if err != nil {
		return info, err
	}

	// install scripts only, if it is not dry-run only
	if !config.GetConfiguration().K8sManagement.DryRunOnly {
		// install scripts
		infoLog, err = ShellScriptsInstall(namespace)
		info = info + constants.NewLine + infoLog
		if err != nil {
			return info, err
		}
	}

	return info, err
}

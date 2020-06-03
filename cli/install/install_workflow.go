package install

import (
	"k8s-management-go/cli/dialogs"
	"k8s-management-go/cli/secrets"
	"k8s-management-go/constants"
	"k8s-management-go/models/config"
	"k8s-management-go/utils/files"
)

// workflow for Jenkins installation
func JenkinsInstallOrUpgrade(helmCommand string) (info string, err error) {
	// ask for namespace
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		return info, err
	}

	// check if project configuration contains Jenkins Helm values file
	jenkinsHelmValuesExist := files.FileOrDirectoryExists(
		files.AppendPath(
			files.AppendPath(
				files.AppendPath(
					config.GetConfiguration().BasePath,
					config.GetConfiguration().Directories.ProjectsBaseDirectory,
				),
				namespace,
			),
			constants.FilenameJenkinsHelmValues,
		),
	)

	// check if namespace is available or create a new one if not
	infoLog, err := CheckAndCreateNamespace(namespace)
	info = info + infoLog
	if err != nil {
		return info, err
	}

	// check if PVC was specified and install it if needed
	infoLog, err = PersistenceVolumeClaimInstall(namespace)
	info = info + infoLog
	if err != nil {
		return info, err
	}

	// apply secrets only if Jenkins Helm values are existing
	if jenkinsHelmValuesExist {
		// apply secrets
		infoLog, err = secrets.ApplySecretsToNamespace(namespace)
		info = info + infoLog
		if err != nil {
			return info, err
		}

		// ask for deployment name
		deploymentName, err := dialogs.DialogAskForDeploymentName("Deployment name", nil)
		if err != nil {
			return info, err
		}

		// install Jenkins
		infoLog, err := HelmInstallJenkins(helmCommand, deploymentName, namespace)
		info = info + infoLog
		if err != nil {
			return info, err
		}
	}

	// install scripts
	infoLog, err = ShellScriptsInstall(namespace)
	info = info + infoLog
	if err != nil {
		return info, err
	}

	// TODO install ingress

	return info, err
}

package uninstall

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/files"
)

// workflow for uninstall
func DoUninstall() (info string, err error) {
	// ask for namespace
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		return info, err
	}

	// first check if namespace directory exists
	projectPath := files.AppendPath(
		config.GetProjectBaseDirectory(),
		namespace,
	)

	// ask for deployment name
	deploymentName, err := dialogs.DialogAskForDeploymentName("Deployment name", nil)
	if err != nil {
		return info, err
	}

	// start uninstalling Jenkins
	info = info + constants.NewLine + "Uninstalling deployment [" + deploymentName + "]"
	infoLog, err := HelmUninstallJenkins(namespace, deploymentName)
	info = info + constants.NewLine + infoLog
	if err != nil {
		return info, err
	}

	// uninstall nginx ingress controller
	infoLog, err = HelmUninstallNginxIngressController(namespace)
	info = info + constants.NewLine + infoLog
	if err != nil {
		return info, err
	}

	// try to uninstall scripts
	infoLog, err = ShellScriptsUninstall(namespace)
	info = info + constants.NewLine + infoLog
	if err != nil {
		return info, err
	}

	return info, err
}

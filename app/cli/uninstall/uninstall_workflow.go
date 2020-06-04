package uninstall

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
)

// workflow for uninstall
func DoUninstall() (info string, err error) {
	// ask for namespace
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		return info, err
	}

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

	// in dry-run we do not want to uninstall the scripts
	if !config.GetConfiguration().K8sManagement.DryRunOnly {
		// try to uninstall scripts
		infoLog, err = ShellScriptsUninstall(namespace)
		info = info + constants.NewLine + infoLog

		// nginx-ingress-controller cleanup
		infoLog, _ := CleanupK8sNginxIngressController(namespace)
		info = info + constants.NewLine + infoLog
	}

	return info, err
}

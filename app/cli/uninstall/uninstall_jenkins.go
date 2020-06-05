package uninstall

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/logger"
)

// uninstall Jenkins with Helm
func HelmUninstallJenkins(namespace string, deploymentName string) (info string, err error) {
	log := logger.Log()
	log.Info("[Uninstall Jenkins] Try to uninstall Jenkins on namespace [" + namespace + "] with deployment name [" + deploymentName + "]...")

	// execute Helm command
	helmCmdArgs := []string{
		deploymentName,
		"-n", namespace,
	}
	// add dry-run flags if necessary
	if models.GetConfiguration().K8sManagement.DryRunOnly {
		helmCmdArgs = append(helmCmdArgs, "--dry-run", "--debug")
	}
	helmCmdOutput, infoLog, err := helm.ExecutorHelm("uninstall", helmCmdArgs)
	info = info + constants.NewLine + infoLog
	if err != nil {
		log.Error("[Uninstall Jenkins] Unable to uninstall Jenkins on namespace [" + namespace + "] with deployment name [" + deploymentName + "].")
		info = info + constants.NewLine + "Jenkins uninstall aborted. See errors."
		err = errors.New(helmCmdOutput + constants.NewLine + err.Error())
		return info, err
	}
	log.Info("[Uninstall Jenkins] Uninstall Jenkins on namespace [" + namespace + "] with deployment name [" + deploymentName + "] done.")

	return info, err
}

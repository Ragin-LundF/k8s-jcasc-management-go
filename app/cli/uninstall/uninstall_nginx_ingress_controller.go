package uninstall

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/logger"
)

// uninstall Jenkins with Helm
func HelmUninstallNginxIngressController(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[Uninstall NginxIngressCtrl] Try to uninstall nginx-ingress-controller in namespace [" + namespace + "]...")

	// execute Helm command
	helmCmdArgs := []string{
		models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName,
		"-n", namespace,
	}
	// add dry-run flags if necessary
	if models.GetConfiguration().K8sManagement.DryRunOnly {
		helmCmdArgs = append(helmCmdArgs, "--dry-run", "--debug")
	}
	helmCmdOutput, infoLog, err := helm.ExecutorHelm("uninstall", helmCmdArgs)
	info = info + constants.NewLine + infoLog

	if err != nil {
		log.Error("[Uninstall NginxIngressCtrl] Unable to uninstall nginx-ingress-controller on namespace [" + namespace + "].")
		info = info + constants.NewLine + "Nginx Ingress Controller uninstall aborted. See errors."
		err = errors.New(helmCmdOutput + constants.NewLine + err.Error())
		return info, err
	}
	log.Info("[Uninstall NginxIngressCtrl] Uninstall of nginx-ingress-controller in namespace [" + namespace + "] done...")

	return info, err
}

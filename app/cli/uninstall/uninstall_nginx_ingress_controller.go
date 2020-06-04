package uninstall

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/logger"
	"os/exec"
)

// uninstall Jenkins with Helm
func HelmUninstallNginxIngressController(namespace string) (info string, err error) {
	log := logger.Log()

	// execute Helm command
	cmd := exec.Command("helm", "uninstall", config.GetConfiguration().Nginx.Ingress.Controller.DeploymentName, "-n", namespace)
	outputCmd, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("Failed to execute: " + cmd.String())
		info = info + constants.NewLine + "Nginx Ingress Controller uninstall aborted. See errors."
		err = errors.New(string(outputCmd) + constants.NewLine + err.Error())
		return info, err
	}

	info = info + constants.NewLine + "Helm Nginx Ingress Controller uninstall output:"
	info = info + constants.NewLine + "==============="
	info = info + string(outputCmd)
	info = info + constants.NewLine + "==============="

	return info, err
}

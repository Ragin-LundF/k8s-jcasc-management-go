package uninstall_actions

import (
	"fmt"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/loggingstate"
)

// uninstall Jenkins with Helm
func ActionHelmUninstallNginxIngressController(namespace string) (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf("[ActionHelmUninstallNginxIngressController] Try to uninstall nginx-ingress-controller in namespace [%s]...", namespace))

	// prepare Helm command
	helmCmdArgs := []string{
		models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName,
		"-n", namespace,
	}
	// add dry-run flags if necessary
	if models.GetConfiguration().K8sManagement.DryRunOnly {
		helmCmdArgs = append(helmCmdArgs, "--dry-run", "--debug")
	}
	// execute helm command
	if err = helm.ExecutorHelm("uninstall", helmCmdArgs); err != nil {
		return err
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("[ActionHelmUninstallNginxIngressController] Uninstall of nginx-ingress-controller in namespace [%s] done...", namespace))

	return nil
}

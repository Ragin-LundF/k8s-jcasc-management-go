package uninstallactions

import (
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/loggingstate"
)

// ActionHelmUninstallNginxIngressController will uninstall Jenkins with Helm
func ActionHelmUninstallNginxIngressController(namespace string) (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf("[ActionHelmUninstallNginxIngressController] Try to uninstall nginx-ingress-controller in namespace [%s]...", namespace))

	// prepare Helm command
	var helmCmdArgs = []string{
		configuration.GetConfiguration().Nginx.Ingress.Deployment.DeploymentName,
		"-n", namespace,
	}
	// add dry-run flags if necessary
	if configuration.GetConfiguration().K8SManagement.DryRunOnly {
		helmCmdArgs = append(helmCmdArgs, "--dry-run", "--debug")
	}
	// execute helm command
	if err = helm.ExecutorHelm("uninstall", helmCmdArgs); err != nil {
		return err
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("[ActionHelmUninstallNginxIngressController] Uninstall of nginx-ingress-controller in namespace [%s] done...", namespace))

	return nil
}

package uninstall_actions

import (
	"fmt"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/logger"
)

// uninstall Jenkins with Helm
func HelmUninstallJenkins(namespace string, deploymentName string) (err error) {
	log := logger.Log()
	log.Infof("[Uninstall Jenkins] Try to uninstall Jenkins on namespace [%s] with deployment name [%s]...", namespace, deploymentName)

	// prepare Helm command
	helmCmdArgs := []string{
		deploymentName,
		"-n", namespace,
	}
	// add dry-run flags if necessary
	if models.GetConfiguration().K8sManagement.DryRunOnly {
		helmCmdArgs = append(helmCmdArgs, "--dry-run", "--debug")
	}
	// execute Helm command
	if err = helm.ExecutorHelm("uninstall", helmCmdArgs); err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Uninstall Jenkins on namespace [%s] with deployment name [%s] done.", namespace, deploymentName), err.Error())
		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Uninstall Jenkins on namespace [%s] with deployment name [%s] done.", namespace, deploymentName))
	log.Infof("[Uninstall Jenkins] Uninstall Jenkins on namespace [%s] with deployment name [%s] done.", namespace, deploymentName)

	return nil
}

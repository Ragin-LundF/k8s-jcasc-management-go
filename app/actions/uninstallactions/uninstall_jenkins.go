package uninstallactions

import (
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/loggingstate"
)

// ActionHelmUninstallJenkins executes the actions to uninstall Jenkins with Helm
func ActionHelmUninstallJenkins(namespace string, deploymentName string) (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Uninstall Jenkins on namespace [%s] with deployment name [%s] start....", namespace, deploymentName))
	// prepare Helm command
	helmCmdArgs := []string{
		deploymentName,
		"-n", namespace,
	}
	// add dry-run flags if necessary
	if configuration.GetConfiguration().K8SManagement.DryRunOnly {
		helmCmdArgs = append(helmCmdArgs, "--dry-run", "--debug")
	}
	// execute Helm command
	if err = helm.ExecutorHelm(constants.HelmCommandUninstall, helmCmdArgs); err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Uninstall Jenkins on namespace [%s] with deployment name [%s] done.", namespace, deploymentName), err.Error())
		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Uninstall Jenkins on namespace [%s] with deployment name [%s] done.", namespace, deploymentName))

	return nil
}

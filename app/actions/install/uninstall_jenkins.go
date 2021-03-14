package install

import (
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/helm"
	"k8s-management-go/app/utils/loggingstate"
)

// ActionHelmUninstallJenkins executes the actions to uninstall Jenkins with Helm
func (projectConfig *ProjectConfig) ActionHelmUninstallJenkins() (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Uninstall Jenkins on namespace [%s] with deployment name [%s] start....",
		projectConfig.Project.Base.Namespace,
		projectConfig.Project.Base.DeploymentName))
	// prepare Helm command
	var helmCmdArgs = []string{
		projectConfig.Project.Base.DeploymentName,
		"-n", projectConfig.Project.Base.Namespace,
	}
	// add dry-run flags if necessary
	if configuration.GetConfiguration().K8SManagement.DryRunOnly {
		helmCmdArgs = append(helmCmdArgs, "--dry-run", "--debug")
	}
	// execute Helm command
	if err = helm.ExecutorHelm(constants.HelmCommandUninstall, helmCmdArgs); err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
			"  -> Uninstall Jenkins on namespace [%s] with deployment name [%s] done.",
			projectConfig.Project.Base.Namespace,
			projectConfig.Project.Base.DeploymentName), err.Error())
		return err
	}

	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Uninstall Jenkins on namespace [%s] with deployment name [%s] done.",
		projectConfig.Project.Base.Namespace,
		projectConfig.Project.Base.DeploymentName))

	return nil
}

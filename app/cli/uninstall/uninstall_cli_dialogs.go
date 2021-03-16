package uninstall

import (
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/utils/loggingstate"
)

// ShowUninstallDialogs shows the uninstall dialog
func ShowUninstallDialogs() (projectConfig install.ProjectConfig, err error) {
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	projectConfig = install.NewInstallProjectConfig()
	namespace, err := dialogs.DialogAskForNamespace()
	projectConfig.Project.SetNamespace(namespace)

	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get namespace.", err.Error())
		return projectConfig, err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// ask for deployment name
	loggingstate.AddInfoEntry("-> Ask for deployment name...")
	projectConfig.Project.Base.DeploymentName, err = dialogs.DialogAskForDeploymentName("Deployment name", nil)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get deployment name.", err.Error())
		return projectConfig, err
	}
	loggingstate.AddInfoEntry("-> Ask for deployment name...done")

	return projectConfig, err
}

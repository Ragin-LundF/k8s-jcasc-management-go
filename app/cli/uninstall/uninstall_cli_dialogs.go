package uninstall

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
)

func ShowUninstallDialogs() (state models.StateData, err error) {
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	state.Namespace, err = dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get namespace.", err.Error())
		return state, err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// ask for deployment name
	loggingstate.AddInfoEntry("-> Ask for deployment name...")
	state.DeploymentName, err = dialogs.DialogAskForDeploymentName("Deployment name", nil)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get deployment name.", err.Error())
		return state, err
	}
	loggingstate.AddInfoEntry("-> Ask for deployment name...done")

	// start uninstalling Jenkins
	jenkinsHelmValuesFile := files.AppendPath(
		files.AppendPath(
			models.GetProjectBaseDirectory(),
			state.Namespace,
		),
		constants.FilenameJenkinsHelmValues,
	)
	state.JenkinsHelmValuesExist = files.FileOrDirectoryExists(jenkinsHelmValuesFile)
	return state, err
}

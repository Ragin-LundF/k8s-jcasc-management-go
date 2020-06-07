package uninstall

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
)

// workflow for uninstall
func DoUninstall() (err error) {
	bar := dialogs.CreateProgressBar("Uninstalling Jenkins", 10)
	// ask for namespace
	loggingstate.AddInfoEntry("Starting Uninstall...")
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get namespace.", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// ask for deployment name
	loggingstate.AddInfoEntry("-> Ask for deployment name...")
	deploymentName, err := dialogs.DialogAskForDeploymentName("Deployment name", nil)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to get deployment name.", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Ask for deployment name...done")

	// start uninstalling Jenkins
	bar.Describe("Check for Jenkins installation...")
	jenkinsHelmValuesFile := files.AppendPath(
		files.AppendPath(
			models.GetProjectBaseDirectory(),
			namespace,
		),
		constants.FilenameJenkinsHelmValues,
	)
	jenkinsHelmValuesExist := files.FileOrDirectoryExists(jenkinsHelmValuesFile)
	bar.Add(1)
	if jenkinsHelmValuesExist {
		bar.Describe("Uninstalling Jenkins deployment...")
		loggingstate.AddInfoEntry("-> Uninstalling deployment [" + deploymentName + "] on namespace [" + namespace + "]...")
		if err = HelmUninstallJenkins(namespace, deploymentName); err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Unable to uninstall deployment ["+deploymentName+"] on namespace ["+namespace+"]...failed", err.Error())
			return err
		}
		loggingstate.AddInfoEntry("-> Uninstalling deployment [" + deploymentName + "] on namespace [" + namespace + "]...done")
	}

	// uninstall nginx ingress controller
	bar.Describe("Check for Nginx installation...")
	nginxHelmValueFile := files.AppendPath(
		files.AppendPath(
			models.GetProjectBaseDirectory(),
			namespace,
		),
		constants.FilenameNginxIngressControllerHelmValues,
	)
	nginxHelmValuesExist := files.FileOrDirectoryExists(nginxHelmValueFile)
	bar.Add(1)
	if nginxHelmValuesExist {
		bar.Describe("Nginx-ingress-controller found...Uninstalling...")
		loggingstate.AddInfoEntry("-> Uninstalling nginx-ingress-controller on namespace [" + namespace + "]...")
		if err = HelmUninstallNginxIngressController(namespace); err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Uninstalling nginx-ingress-controller on namespace [" + namespace + "]...done")
	}
	bar.Add(1)

	// in dry-run we do not want to uninstall the scripts
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		bar.Describe("Try to execute uninstall scripts...")
		// try to uninstall scripts
		loggingstate.AddInfoEntry("-> Try to execute uninstall scripts on [" + namespace + "]...")
		// we ignore errors. They will be logged, but we keep on doing the uninstall for the scripts
		_ = ShellScriptsUninstall(namespace)
		loggingstate.AddInfoEntry("-> Try to execute uninstall scripts on [" + namespace + "]...done")
		bar.Add(1)

		// nginx-ingress-controller cleanup
		bar.Describe("Cleanup configuration...")
		loggingstate.AddInfoEntry("-> Try to cleanup configuration of [" + namespace + "]...")
		CleanupK8sNginxIngressController(namespace, bar)
		loggingstate.AddInfoEntry("-> Try to cleanup configuration of [" + namespace + "]...done")
	} else {
		bar.Add(7) // skipping previous steps
	}

	loggingstate.AddInfoEntry("Starting Uninstall...done")
	return nil
}

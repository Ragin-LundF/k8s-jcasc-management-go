package uninstall

import (
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/models"
)

// workflow for uninstall
func DoUninstall() (err error) {
	loggingstate.AddInfoEntry("Starting Uninstall...")
	// ask for namespace
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
	loggingstate.AddInfoEntry("-> Uninstalling deployment [" + deploymentName + "] on namespace [" + namespace + "]...")
	if err = HelmUninstallJenkins(namespace, deploymentName); err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to uninstall deployment ["+deploymentName+"] on namespace ["+namespace+"]...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Uninstalling deployment [" + deploymentName + "] on namespace [" + namespace + "]...done")

	// uninstall nginx ingress controller
	loggingstate.AddInfoEntry("-> Uninstalling nginx-ingress-controller on namespace [" + namespace + "]...")
	if err = HelmUninstallNginxIngressController(namespace); err != nil {
		return err
	}
	loggingstate.AddInfoEntry("-> Uninstalling nginx-ingress-controller on namespace [" + namespace + "]...done")

	// in dry-run we do not want to uninstall the scripts
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		// try to uninstall scripts
		loggingstate.AddInfoEntry("-> Try to execute uninstall scripts on [" + namespace + "]...")
		// we ignore errors. They will be logged, but we keep on doing the uninstall for the scripts
		_ = ShellScriptsUninstall(namespace)
		loggingstate.AddInfoEntry("-> Try to execute uninstall scripts on [" + namespace + "]...done")

		// nginx-ingress-controller cleanup
		loggingstate.AddInfoEntry("-> Try to cleanup configuration of [" + namespace + "]...")
		CleanupK8sNginxIngressController(namespace)
		loggingstate.AddInfoEntry("-> Try to cleanup configuration of [" + namespace + "]...done")
	}

	loggingstate.AddInfoEntry("Starting Uninstall...done")
	return nil
}

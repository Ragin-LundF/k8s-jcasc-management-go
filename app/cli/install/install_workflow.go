package install

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/cli/secrets"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
)

// workflow for Jenkins installation
func DoUpgradeOrInstall(helmCommand string) (err error) {
	log := logger.Log()
	// ask for namespace
	loggingstate.AddInfoEntry("Starting Jenkins [" + helmCommand + "]...")
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Ask for namespace...failed.", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// first check if namespace directory exists
	loggingstate.AddInfoEntry("-> Checking existing directories...")
	projectPath := files.AppendPath(
		models.GetProjectBaseDirectory(),
		namespace,
	)
	if !files.FileOrDirectoryExists(projectPath) {
		err = errors.New("Project directory not found: [" + projectPath + "]")
		loggingstate.AddErrorEntryAndDetails("-> Checking existing directories...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Checking existing directories...done")

	// create namespace and pvc only, if it is not dry-run only
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		// check if namespace is available or create a new one if not
		loggingstate.AddInfoEntry("-> Check and create namespace if necessary...")
		if err = CheckAndCreateNamespace(namespace); err != nil {
			loggingstate.AddErrorEntryAndDetails("-> Check and create namespace if necessary...failed", err.Error())
			return err
		}
		loggingstate.AddInfoEntry("-> Check and create namespace if necessary...done")

		// check if PVC was specified and install it if needed
		loggingstate.AddInfoEntry("-> Check and create pvc if necessary...")
		if err = PersistenceVolumeClaimInstall(namespace); err != nil {
			loggingstate.AddErrorEntryAndDetails("-> Check and create pvc if necessary...failed", err.Error())
			return err
		}
		loggingstate.AddInfoEntry("-> Check and create pvc if necessary...done")
	} else {
		loggingstate.AddInfoEntry("-> Dry run. Skipping namespace creation and pvc install...")
	}

	// check if project configuration contains Jenkins Helm values file
	jenkinsHelmValuesFile := files.AppendPath(
		projectPath,
		constants.FilenameJenkinsHelmValues,
	)
	jenkinsHelmValuesExist := files.FileOrDirectoryExists(jenkinsHelmValuesFile)

	// apply secrets only if Jenkins Helm values are existing
	if jenkinsHelmValuesExist {
		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...")
		log.Infof("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...")
		// apply secrets only, if it is not dry-run only
		if !models.GetConfiguration().K8sManagement.DryRunOnly {
			// apply secrets
			log.Infof("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...", namespace)
			loggingstate.AddInfoEntry("  -> Starting apply secrets to namespace [" + namespace + "]...")

			if err = secrets.ApplySecretsToNamespace(namespace); err != nil {
				loggingstate.AddErrorEntryAndDetails("  -> Starting apply secrets to namespace ["+namespace+"]...failed", err.Error())
				log.Errorf("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...failed\n%s", err.Error())
				return err
			}

			loggingstate.AddInfoEntry("  -> Starting apply secrets to namespace [" + namespace + "]...done")
			log.Infof("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...done", namespace)
		} else {
			loggingstate.AddInfoEntry("  -> Dry run only, skipping apply secrets to namespace [" + namespace + "]...")
			log.Infof("[DoUpgradeOrInstall] Dry run only, skipping apply secrets to namespace [%s]...", namespace)
		}

		// ask for deployment name
		deploymentName, err := dialogs.DialogAskForDeploymentName("Deployment name", nil)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Unable to get deployment name.", err.Error())
			log.Errorf("[DoUpgradeOrInstall] Unable to get deployment name.\n%s", err.Error())
			return err
		}

		// install Jenkins
		err = HelmInstallJenkins(helmCommand, namespace, deploymentName)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Unable to install Jenkins.", err.Error())
			log.Errorf("[DoUpgradeOrInstall] Unable to install Jenkins.\n%s", err.Error())
			return err
		}

		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...done")
		log.Infof("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...done")
	} else {
		loggingstate.AddInfoEntry("No Jenkins Helm chart found in path [" + jenkinsHelmValuesFile + "].")
		log.Infof("No Jenkins Helm chart found in path [%s]. Skipping Jenkins installation.", jenkinsHelmValuesFile)
	}

	// install Nginx ingress controller
	loggingstate.AddInfoEntry("-> Installing nginx-ingress-controller on namespace [" + namespace + "]...")
	err = HelmInstallNginxIngressController(helmCommand, namespace, jenkinsHelmValuesExist)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to install nginx-ingress-controller.", err.Error())
		log.Errorf("[DoUpgradeOrInstall] Unable to install nginx-ingress-controller.\n%s", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Installing nginx-ingress-controller on namespace [" + namespace + "]...done")

	// install scripts only, if it is not dry-run only
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		// install scripts
		// try to install scripts
		loggingstate.AddInfoEntry("-> Try to execute install scripts on [" + namespace + "]...")
		// we ignore errors. They will be logged, but we keep on doing the install for the scripts
		_ = ShellScriptsInstall(namespace)
		loggingstate.AddInfoEntry("-> Try to execute install scripts on [" + namespace + "]...done")
	}

	loggingstate.AddInfoEntry("Starting Jenkins [" + helmCommand + "]...done")
	return err
}

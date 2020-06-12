package install

import (
	"errors"
	"fmt"
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
	loggingstate.AddInfoEntry(fmt.Sprintf("Starting Jenkins [%s]...", helmCommand))
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	namespace, err := dialogs.DialogAskForNamespace()
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Ask for namespace...failed.", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// Progress Bar
	bar := dialogs.CreateProgressBar("Installing...", 10)

	// first check if namespace directory exists
	bar.Describe("Checking directories...")
	loggingstate.AddInfoEntry("-> Checking existing directories...")
	projectPath := files.AppendPath(
		models.GetProjectBaseDirectory(),
		namespace,
	)
	if !files.FileOrDirectoryExists(projectPath) {
		err = errors.New(fmt.Sprintf("Project directory not found: [%s]", projectPath))
		loggingstate.AddErrorEntryAndDetails("-> Checking existing directories...failed", err.Error())
		return err
	}
	loggingstate.AddInfoEntry("-> Checking existing directories...done")

	// progress
	bar.Add(1)

	// create namespace and pvc only, if it is not dry-run only
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		// check if namespace is available or create a new one if not
		bar.Describe("Check and create namespace if necessary...")

		loggingstate.AddInfoEntry("-> Check and create namespace if necessary...")
		if err = CheckAndCreateNamespace(namespace); err != nil {
			loggingstate.AddErrorEntryAndDetails("-> Check and create namespace if necessary...failed", err.Error())
			return err
		}
		loggingstate.AddInfoEntry("-> Check and create namespace if necessary...done")

		// progress
		bar.Add(1)
		bar.Describe("Check and create PVC if necessary...")

		// check if PVC was specified and install it if needed
		loggingstate.AddInfoEntry("-> Check and create pvc if necessary...")
		if err = PersistenceVolumeClaimInstall(namespace); err != nil {
			loggingstate.AddErrorEntryAndDetails("-> Check and create pvc if necessary...failed", err.Error())
			return err
		}
		loggingstate.AddInfoEntry("-> Check and create pvc if necessary...done")
		bar.Add(1)
	} else {
		loggingstate.AddInfoEntry("-> Dry run. Skipping namespace creation and pvc install...")
		bar.Add(3) // skipping steps above
	}

	// check if project configuration contains Jenkins Helm values file
	bar.Describe("Checking Jenkins...")
	jenkinsHelmValuesFile := files.AppendPath(
		projectPath,
		constants.FilenameJenkinsHelmValues,
	)
	jenkinsHelmValuesExist := files.FileOrDirectoryExists(jenkinsHelmValuesFile)
	bar.Add(1)

	// apply secrets only if Jenkins Helm values are existing
	if jenkinsHelmValuesExist {
		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...")
		log.Infof("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...")
		// apply secrets only, if it is not dry-run only
		if !models.GetConfiguration().K8sManagement.DryRunOnly {
			// apply secrets
			bar.Describe("Applying secrets...")
			bar.Add(1)
			log.Infof("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...", namespace)
			loggingstate.AddInfoEntry(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...", namespace))

			if err = secrets.ApplySecretsToNamespace(namespace, &bar); err != nil {
				loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...failed", namespace), err.Error())
				log.Errorf("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...failed\n%s", err.Error())
				return err
			}

			loggingstate.AddInfoEntry(fmt.Sprintf("  -> Starting apply secrets to namespace [%s]...done", namespace))
			log.Infof("[DoUpgradeOrInstall] Starting apply secrets to namespace [%s]...done", namespace)
		} else {
			loggingstate.AddInfoEntry(fmt.Sprintf("  -> Dry run only, skipping apply secrets to namespace [%s]...", namespace))
			log.Infof("[DoUpgradeOrInstall] Dry run only, skipping apply secrets to namespace [%s]...", namespace)
		}

		// ask for deployment name
		bar.Describe("Check for deployment name...") // if dialog came up
		bar.Add(1)
		deploymentName, err := dialogs.DialogAskForDeploymentName("Deployment name", nil)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Unable to get deployment name.", err.Error())
			log.Errorf("[DoUpgradeOrInstall] Unable to get deployment name.\n%s", err.Error())
			return err
		}

		// install Jenkins
		bar.Describe("Installing Jenkins...")
		err = HelmInstallJenkins(helmCommand, namespace, deploymentName)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Unable to install Jenkins.", err.Error())
			log.Errorf("[DoUpgradeOrInstall] Unable to install Jenkins.\n%s", err.Error())
			return err
		}

		bar.Add(1)
		loggingstate.AddInfoEntry("-> Jenkins Helm values.yaml found. Installing Jenkins...done")
		log.Infof("[DoUpgradeOrInstall] Jenkins Helm values.yaml found. Installing Jenkins...done")
	} else {
		bar.Add(2) // skipping steps above
		loggingstate.AddInfoEntry(fmt.Sprintf("No Jenkins Helm chart found in path [%s].", jenkinsHelmValuesFile))
		log.Infof("No Jenkins Helm chart found in path [%s]. Skipping Jenkins installation.", jenkinsHelmValuesFile)
	}

	// install Nginx ingress controller
	bar.Describe("Installing nginx-ingress-controller...")
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Installing nginx-ingress-controller on namespace [%s]...", namespace))
	err = HelmInstallNginxIngressController(helmCommand, namespace, jenkinsHelmValuesExist)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to install nginx-ingress-controller.", err.Error())
		log.Errorf("[DoUpgradeOrInstall] Unable to install nginx-ingress-controller.\n%s", err.Error())
		return err
	}
	bar.Add(1)
	loggingstate.AddInfoEntry(fmt.Sprintf("-> Installing nginx-ingress-controller on namespace [%s]...done", namespace))

	// install scripts only, if it is not dry-run only
	if !models.GetConfiguration().K8sManagement.DryRunOnly {
		bar.Describe("Check and execute additional scripts...")
		// install scripts
		// try to install scripts
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute install scripts on [%s]...", namespace))
		// we ignore errors. They will be logged, but we keep on doing the install for the scripts
		_ = ShellScriptsInstall(namespace)
		loggingstate.AddInfoEntry(fmt.Sprintf("-> Try to execute install scripts on [%s]...done", namespace))
	}
	bar.Add(1)

	loggingstate.AddInfoEntry(fmt.Sprintf("Starting Jenkins [%s]...done", helmCommand))
	return err
}

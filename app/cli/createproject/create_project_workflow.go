package createproject

import (
	"k8s-management-go/app/utils/loggingstate"
)

func ProjectWizardWorkflow(deploymentOnly bool) (err error) {
	// Start project wizard
	loggingstate.AddInfoEntry("Starting Project Wizard: Dialogs...")

	// Ask for namespace
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	namespace, err := ProjectWizardAskForNamespace()
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// Ask for IP address
	loggingstate.AddInfoEntry("-> Ask for IP address...")
	ipAddress, err := ProjectWizardAskForIpAddress()
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry("-> Ask for IP address...done")

	// declare vars for next if statement
	var jenkinsSystemMsg string
	var jobsCfgRepo string
	var selectedCloudTemplates []string
	var existingPvc string

	// if it is not only a deployment project ask for other Jenkins related vars
	if !deploymentOnly {
		// Select cloud templates
		loggingstate.AddInfoEntry("-> Ask for cloud templates...")
		selectedCloudTemplates, err = ProjectWizardAskForCloudTemplates()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Ask for cloud templates...done")

		// Ask for existing persistent volume claim (PVC)
		loggingstate.AddInfoEntry("-> Ask for persistent volume claim...")
		existingPvc, err = ProjectWizardAskForExistingPersistentVolumeClaim()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Ask for persistent volume claim...done")

		// Ask for Jenkins system message
		loggingstate.AddInfoEntry("-> Ask for Jenkins system message...")
		jenkinsSystemMsg, err = ProjectWizardAskForJenkinsSystemMessage(namespace)
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Ask for Jenkins system message...done")

		// Ask for Jobs Configuration repository
		loggingstate.AddInfoEntry("-> Ask for jobs configuration repository...")
		jobsCfgRepo, err = ProjectWizardAskForJobsConfigurationRepository()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Ask for jobs configuration repository...done")
	}
	loggingstate.AddInfoEntry("Starting Project Wizard: Dialogs...done")

	// Process data and create project
	loggingstate.AddInfoEntry("Starting Project Wizard: Template processing...")
	err = ProcessProjectCreate(namespace, ipAddress, jenkinsSystemMsg, jobsCfgRepo, existingPvc, selectedCloudTemplates, deploymentOnly)
	if err != nil {
		loggingstate.AddInfoEntry("Starting Project Wizard: Template processing...failed")
	}
	loggingstate.AddInfoEntry("Starting Project Wizard: Template processing...done")

	return nil
}

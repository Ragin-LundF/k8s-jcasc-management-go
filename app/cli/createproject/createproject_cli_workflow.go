package createproject

import (
	"k8s-management-go/app/actions/createproject"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

func ProjectWizardWorkflow(deploymentOnly bool) (err error) {
	var projectConfig models.ProjectConfig
	projectConfig.CreateDeploymentOnlyProject = deploymentOnly

	// Start project wizard
	loggingstate.AddInfoEntry("Starting Project Wizard: Dialogs...")

	// Ask for namespace
	loggingstate.AddInfoEntry("-> Ask for namespace...")
	projectConfig.Namespace, err = NamespaceWorkflow()
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry("-> Ask for namespace...done")

	// Ask for IP address
	loggingstate.AddInfoEntry("-> Ask for IP address...")
	projectConfig.IpAddress, err = IpAddressWorkflow()
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry("-> Ask for IP address...done")

	// if it is not only a deployment project ask for other Jenkins related vars
	if !projectConfig.CreateDeploymentOnlyProject {
		// Select cloud templates
		loggingstate.AddInfoEntry("-> Ask for cloud templates...")
		projectConfig.SelectedCloudTemplates, err = CloudTemplatesWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Ask for cloud templates...done")

		// Ask for existing persistent volume claim (PVC)
		loggingstate.AddInfoEntry("-> Ask for persistent volume claim...")
		projectConfig.ExistingPvc, err = PersistentVolumeClaimWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Ask for persistent volume claim...done")

		// Ask for Jenkins system message
		loggingstate.AddInfoEntry("-> Ask for Jenkins system message...")
		projectConfig.JenkinsSystemMsg, err = JenkinsSystemMessageWorkflow(projectConfig.Namespace)
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Ask for Jenkins system message...done")

		// Ask for Jobs Configuration repository
		loggingstate.AddInfoEntry("-> Ask for jobs configuration repository...")
		projectConfig.JobsCfgRepo, err = JenkinsJobsConfigRepositoryWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry("-> Ask for jobs configuration repository...done")
	}
	loggingstate.AddInfoEntry("Starting Project Wizard: Dialogs...done")

	// Process data and create project
	loggingstate.AddInfoEntry("Starting Project Wizard: Template processing...")

	// prepare progressbar
	var maxProgressCnt = createproject.CountCreateProjectWorkflow
	bar := dialogs.CreateProgressBar("Create project...", maxProgressCnt)
	progress := dialogs.ProgressBar{
		Bar: &bar,
	}

	// Create project
	err = createproject.ActionProcessProjectCreate(projectConfig, progress.AddCallback)
	if err != nil {
		loggingstate.AddInfoEntry("Starting Project Wizard: Template processing...failed")
	}
	loggingstate.AddInfoEntry("Starting Project Wizard: Template processing...done")

	return nil
}

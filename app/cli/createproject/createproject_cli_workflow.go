package createproject

import (
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/loggingstate"
)

// ProjectWizardWorkflow represents the project wizard workflow
func ProjectWizardWorkflow(deploymentOnly bool) (err error) {
	var projectConfig models.ProjectConfig
	projectConfig.CreateDeploymentOnlyProject = deploymentOnly

	// Start project wizard
	loggingstate.AddInfoEntry(constants.LogWizardStartProjectWizardDialogs)

	// Ask for namespace
	loggingstate.AddInfoEntry(constants.LogAskForNamespace)
	projectConfig.Namespace, err = NamespaceWorkflow()
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry(constants.LogAskForNamespaceDone)

	// Ask for IP address
	loggingstate.AddInfoEntry(constants.LogAskForIPAddress)
	projectConfig.IPAddress, err = IPAddressWorkflow()
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry(constants.LogAskForIPAddressDone)

	// Ask for Domain
	loggingstate.AddInfoEntry(constants.LogAskForJenkinsUrl)
	projectConfig.JenkinsDomain, err = JenkinsDomainWorkflow()
	if err != nil {
		return err
	}
	if projectConfig.JenkinsDomain == "" && models.GetConfiguration().LoadBalancer.Annotations.ExtDNS.Hostname != "" {
		projectConfig.JenkinsDomain = projectConfig.Namespace + models.GetConfiguration().LoadBalancer.Annotations.ExtDNS.Hostname
	}
	loggingstate.AddInfoEntry(constants.LogAskForJenkinsUrlDone)

	// if it is not only a deployment project ask for other Jenkins related vars
	if !projectConfig.CreateDeploymentOnlyProject {
		// Select cloud templates
		loggingstate.AddInfoEntry(constants.LogAskForCloudTemplates)
		projectConfig.SelectedCloudTemplates, err = CloudTemplatesWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry(constants.LogAskForCloudTemplatesDone)

		// Ask for existing persistent volume claim (PVC)
		loggingstate.AddInfoEntry(constants.LogAskForPvc)
		projectConfig.ExistingPvc, err = PersistentVolumeClaimWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry(constants.LogAskForPvcDone)

		// Ask for Jenkins system message
		loggingstate.AddInfoEntry(constants.LogAskForJenkinsSystemMessage)
		projectConfig.JenkinsSystemMsg, err = JenkinsSystemMessageWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry(constants.LogAskForJenkinsSystemMessageDone)

		// Ask for Jobs Configuration repository
		loggingstate.AddInfoEntry(constants.LogAskForJobsConfigurationRepository)
		projectConfig.JobsCfgRepo, err = JenkinsJobsConfigRepositoryWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry(constants.LogAskForJobsConfigurationRepositoryDone)
	}
	loggingstate.AddInfoEntry(constants.LogWizardStartProjectWizardDialogsDone)

	// Process data and create project
	loggingstate.AddInfoEntry(constants.LogWizardStartProcessingTemplates)

	// prepare progressbar
	var maxProgressCnt = project.CountCreateProjectWorkflow
	bar := dialogs.CreateProgressBar(constants.ActionCreateProject, maxProgressCnt)
	progress := dialogs.ProgressBar{
		Bar: &bar,
	}

	// Create project
	err = project.ActionProcessProjectCreate(projectConfig, progress.AddCallback)
	if err != nil {
		loggingstate.AddInfoEntry(constants.LogWizardStartProcessingTemplatesFailed)
	}
	loggingstate.AddInfoEntry(constants.LogWizardStartProcessingTemplatesDone)

	return nil
}

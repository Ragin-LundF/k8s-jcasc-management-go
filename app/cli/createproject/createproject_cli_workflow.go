package createproject

import (
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
)

// ProjectWizardWorkflow represents the project wizard workflow
func ProjectWizardWorkflow(deploymentOnly bool) (err error) {
	var prj = project.NewProject()
	prj.Base.DeploymentOnly = deploymentOnly

	// Start project wizard
	loggingstate.AddInfoEntry(constants.LogWizardStartProjectWizardDialogs)

	// Ask for namespace
	loggingstate.AddInfoEntry(constants.LogAskForNamespace)
	prj.Base.Namespace, err = NamespaceWorkflow()
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry(constants.LogAskForNamespaceDone)

	// Ask for IP address
	loggingstate.AddInfoEntry(constants.LogAskForIPAddress)
	prj.Base.IPAddress, err = IPAddressWorkflow()
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry(constants.LogAskForIPAddressDone)

	// Ask for Domain
	loggingstate.AddInfoEntry(constants.LogAskForJenkinsUrl)
	prj.Base.Domain, err = JenkinsDomainWorkflow()
	if err != nil {
		return err
	}
	if prj.Base.Domain == "" && configuration.GetConfiguration().Nginx.Loadbalancer.ExternalDNS.HostName != "" {
		prj.SetDomain(prj.Base.Domain + configuration.GetConfiguration().Nginx.Loadbalancer.ExternalDNS.HostName)
	}
	loggingstate.AddInfoEntry(constants.LogAskForJenkinsUrlDone)

	// if it is not only a deployment project ask for other Jenkins related vars
	if !prj.Base.DeploymentOnly {
		// Select cloud templates
		loggingstate.AddInfoEntry(constants.LogAskForCloudTemplates)
		prj.JCasc.Clouds.Kubernetes.Templates.AdditionalCloudTemplateFiles, err = CloudTemplatesWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry(constants.LogAskForCloudTemplatesDone)

		// Ask for existing persistent volume claim (PVC)
		loggingstate.AddInfoEntry(constants.LogAskForPvc)
		prj.Base.ExistingVolumeClaim, err = PersistentVolumeClaimWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry(constants.LogAskForPvcDone)

		// Ask for Jenkins system message
		loggingstate.AddInfoEntry(constants.LogAskForJenkinsSystemMessage)
		prj.JCasc.SystemMessage, err = JenkinsSystemMessageWorkflow()
		if err != nil {
			return err
		}
		loggingstate.AddInfoEntry(constants.LogAskForJenkinsSystemMessageDone)

		// Ask for Jobs Configuration repository
		loggingstate.AddInfoEntry(constants.LogAskForJobsConfigurationRepository)
		prj.JCasc.JobsConfig.JobsDefinitionRepository, err = JenkinsJobsConfigRepositoryWorkflow()
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
	err = prj.ActionProcessProjectCreate(progress.AddCallback)
	if err != nil {
		loggingstate.AddInfoEntry(constants.LogWizardStartProcessingTemplatesFailed)
	}
	loggingstate.AddInfoEntry(constants.LogWizardStartProcessingTemplatesDone)

	return nil
}

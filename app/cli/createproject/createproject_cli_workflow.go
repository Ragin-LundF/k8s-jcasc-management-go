package createproject

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// ProjectWizardWorkflow represents the project wizard workflow
func ProjectWizardWorkflow(deploymentOnly bool) (err error) {
	var prj = project.NewProject()
	prj.Base.DeploymentOnly = deploymentOnly

	// Start project wizard
	loggingstate.AddInfoEntry(constants.LogWizardStartProjectWizardDialogs)

	// Ask for store config only
	prj.StoreConfigOnly, err = StoreConfigOnlyWorkflow()
	if err != nil {
		return err
	}

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

		// Ask for additional namespaces
		loggingstate.AddInfoEntry(constants.LogAskForNamespace)
		additionalNamespaces, err := AdditionalNamespaceWorkflow()
		if err != nil {
			return err
		}
		prj.SetAdditionalNamespaces(additionalNamespaces)
		loggingstate.AddInfoEntry(constants.LogAskForNamespaceDone)
	}
	loggingstate.AddInfoEntry(constants.LogWizardStartProjectWizardDialogsDone)

	// Process data and create project
	loggingstate.AddInfoEntry(constants.LogWizardStartProcessingTemplates)

	// prepare progressbar
	var maxProgressCnt = project.CountCreateProjectWorkflow
	var bar = dialogs.CreateProgressBar(constants.ActionCreateProject, maxProgressCnt)
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

// StoreConfigOnlyWorkflow asks if the project should only store the config
func StoreConfigOnlyWorkflow() (storeConfigOnly bool, err error) {
	// Prepare prompt
	dialogs.ClearScreen()
	var promptResult string
	var storeConfigOnlyPrompt = promptui.Prompt{
		Label:     "Store config only?",
		IsConfirm: true,
		Default:   "y",
	}
	promptResult, err = storeConfigOnlyPrompt.Run()

	// check if everything was ok
	if err != nil && err.Error() != "" {
		loggingstate.AddErrorEntryAndDetails(constants.LogUnableToGetStoreConfigOnly, err.Error())
		return storeConfigOnly, err
	}

	if strings.ToLower(promptResult) == "y" || len(promptResult) == 0 {
		return true, nil
	} else if strings.ToLower(promptResult) == "n" {
		return false, nil
	} else {
		return true, fmt.Errorf("Please confirm with y or n. You typed [%v] ", promptResult)
	}
}

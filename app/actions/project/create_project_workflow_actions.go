package project

import (
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/events"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"time"
)

// CountCreateProjectWorkflow represents the max count for the progress bar
const CountCreateProjectWorkflow = 13

// ActionProcessProjectCreate is processing the project creation. This method controls all required actions
func ActionProcessProjectCreate(projectConfig models.ProjectConfig, callback func()) (err error) {
	// calculate the target directory
	newProjectDir := files.AppendPath(configuration.GetConfiguration().GetProjectBaseDirectory(), projectConfig.Namespace)
	callback()

	// create new project directory
	loggingstate.AddInfoEntryAndDetails("-> Create new project directory...", "Directory: ["+newProjectDir+"]")
	err = ActionCreateNewProjectDirectory(newProjectDir)
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry("-> Create new project directory...done")
	callback()

	// copy necessary files
	loggingstate.AddInfoEntryAndDetails("-> Start to copy templates to new project directory...", "Directory: ["+newProjectDir+"]")
	err = ActionCopyTemplatesToNewDirectory(newProjectDir, len(projectConfig.ExistingPvc) > 0, projectConfig.CreateDeploymentOnlyProject)
	if err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start to copy templates to new project directory...done")
	callback()

	// add IP and namespace to IP configuration
	loggingstate.AddInfoEntry("-> Start adding IP address to ipconfig file...")
	success, err := config.AddToIPConfigFile(projectConfig.Namespace, projectConfig.IPAddress)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start adding IP address to ipconfig file...done")
	callback()

	// setup project
	loggingstate.AddInfoEntry("-> Start template processing...")
	var newProject = NewProject()
	newProject.SetDomain(projectConfig.JenkinsDomain)
	newProject.SetIPAddress(projectConfig.IPAddress)
	newProject.SetJenkinsSystemMessage(projectConfig.JenkinsSystemMsg)
	newProject.SetJobsDefinitionRepository(projectConfig.JobsCfgRepo)
	newProject.SetNamespace(projectConfig.Namespace)
	newProject.SetPersistentVolumeClaimExistingName(projectConfig.ExistingPvc)

	cloudTemplatesString, err := ActionReadCloudTemplatesAsString(projectConfig.SelectedCloudTemplates)
	if err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	newProject.SetCloudKubernetesAdditionalTemplates(cloudTemplatesString)

	err = newProject.ProcessTemplates(newProjectDir)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Start template processing...error", err.Error())
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing...done")
	callback()

	// reload config
	models.ResetIPAndNamespaces()
	config.ReadIPConfig()

	// send event that new namespace was created
	createNamespaceEvent(projectConfig.Namespace)

	return nil
}

func createNamespaceEvent(namespace string) {
	if !configuration.GetConfiguration().K8SManagement.CliOnly {
		events.NamespaceCreated.Trigger(events.NamespaceCreatedPayload{
			Namespace: namespace,
			Time:      time.Now(),
		})
	}
}

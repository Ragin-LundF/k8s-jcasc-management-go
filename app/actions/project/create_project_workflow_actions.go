package project

import (
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/events"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"time"
)

// CountCreateProjectWorkflow represents the max count for the progress bar
const CountCreateProjectWorkflow = 13

// ActionProcessProjectCreate is processing the project creation. This method controls all required actions
func (prj *Project) ActionProcessProjectCreate(callback func()) (err error) {
	// calculate the target directory
	var newProjectDir = files.AppendPath(configuration.GetConfiguration().GetProjectBaseDirectory(), prj.Base.Namespace)
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
	err = prj.ActionCopyTemplatesToNewDirectory(newProjectDir)
	if err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start to copy templates to new project directory...done")
	callback()

	// add IP and namespace to IP configuration
	loggingstate.AddInfoEntry("-> Start adding IP address to ipconfig file...")
	success, err := configuration.GetConfiguration().AddToIPConfigFile(prj.Base.Namespace, prj.Base.IPAddress, prj.Base.Domain)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start adding IP address to ipconfig file...done")
	callback()

	// setup project
	loggingstate.AddInfoEntry("-> Start template processing...")
	cloudTemplatesString, err := ActionReadCloudTemplatesAsString(prj.JCasc.Clouds.Kubernetes.Templates.AdditionalCloudTemplateFiles)
	if err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	prj.SetCloudKubernetesAdditionalTemplates(cloudTemplatesString)

	err = prj.ProcessTemplates(newProjectDir)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Start template processing...error", err.Error())
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing...done")
	callback()

	// send event that new namespace was created
	createNamespaceEvent(prj.Base.Namespace)

	err = prj.SaveProjectConfiguration(newProjectDir)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("-> Error while saving project configuration", err.Error())
		return err
	}

	return nil
}

// create event for namespace update
func createNamespaceEvent(namespace string) {
	if !configuration.GetConfiguration().K8SManagement.CliOnly {
		events.NamespaceCreated.Trigger(events.NamespaceCreatedPayload{
			Namespace: namespace,
			Time:      time.Now(),
		})
	}
}

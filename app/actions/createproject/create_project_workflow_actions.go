package createproject

import (
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
)

const CountCreateProjectWorkflow = 13

// Processing project creation. This method controls all required actions
func ActionProcessProjectCreate(projectConfig models.ProjectConfig, callback func()) (err error) {
	// calculate the target directory
	newProjectDir := files.AppendPath(models.GetProjectBaseDirectory(), projectConfig.Namespace)
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
	success, err := config.AddToIpConfigFile(projectConfig.Namespace, projectConfig.IpAddress)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start adding IP address to ipconfig file...done")
	callback()

	// processing cloud templates
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins cloud templates...")
	success, err = ActionProcessTemplateCloudTemplates(newProjectDir, projectConfig.SelectedCloudTemplates)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins cloud templates...done")
	callback()

	// Replace Jenkins system message
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins system message...")
	templateFiles := []string{files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode)}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplateJenkinsSystemMessage, projectConfig.JenkinsSystemMsg)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins system message...done")
	callback()

	// Replace Jenkins seed job repository
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins Jobs repository...")
	templateFiles = []string{files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode)}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplateJobDefinitionRepository, projectConfig.JobsCfgRepo)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins Jobs repository...done")
	callback()

	// Replace global configuration
	loggingstate.AddInfoEntry("-> Start template processing: Global configuration...")
	success, err = ActionReplaceGlobalConfigDelegation(newProjectDir)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Global configuration...done")
	callback()

	// Replace namespace
	loggingstate.AddInfoEntry("-> Start template processing: Namespaces...")
	templateFiles = []string{
		files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode),
		files.AppendPath(newProjectDir, constants.FilenamePvcClaim),
		files.AppendPath(newProjectDir, constants.FilenameNginxIngressControllerHelmValues),
	}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplateNamespace, projectConfig.Namespace)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Namespaces...done")
	callback()

	// Replace ip address
	loggingstate.AddInfoEntry("-> Start template processing: IP address...")
	templateFiles = []string{
		files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode),
		files.AppendPath(newProjectDir, constants.FilenameNginxIngressControllerHelmValues),
	}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplatePublicIpAddress, projectConfig.IpAddress)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: IP address...done")
	callback()

	// Replace project directory with namespace name
	loggingstate.AddInfoEntry("-> Start template processing: Project directory for JCasC...")
	templateFiles = []string{
		files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode),
		files.AppendPath(newProjectDir, constants.FilenameJenkinsHelmValues),
	}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplateProjectDirectory, projectConfig.Namespace)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Project directory for JCasC...done")
	callback()

	// Replace existing pvc in Jenkins
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins existing persistent volume claim...")
	templateFiles = []string{files.AppendPath(newProjectDir, constants.FilenameJenkinsHelmValues)}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplatePvcExistingVolumeClaim, projectConfig.ExistingPvc)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins existing persistent volume claim...done")
	callback()

	// Replace existing pvc in Jenkins
	loggingstate.AddInfoEntry("-> Start template processing: Persistent volume claim...")
	templateFiles = []string{files.AppendPath(newProjectDir, constants.FilenamePvcClaim)}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplatePvcName, projectConfig.ExistingPvc)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Persistent volume claim...done")
	callback()

	return nil
}

// replace project directory with namespace name
func replacePlaceholderInTemplates(templateFiles []string, placeholder string, newValue string) (success bool, err error) {
	for _, templateFile := range templateFiles {
		if files.FileOrDirectoryExists(templateFile) {
			successful, err := files.ReplaceStringInFile(templateFile, placeholder, newValue)
			if !successful || err != nil {
				loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to replace [%s] in file [%s]", placeholder, templateFile), err.Error())
				return false, err
			}
		}
	}
	return true, nil
}

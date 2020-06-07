package createproject

import (
	"k8s-management-go/app/cli/logoutput"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os"
	"strconv"
)

// Processing project creation
func ProcessProjectCreate(namespace string, ipAddress string, jenkinsSystemMsg string, jobsCfgRepo string, existingPvc string, selectedCloudTemplates []string, createDeploymentOnlyProject bool) (err error) {
	// calculate the target directory
	newProjectDir := files.AppendPath(models.GetProjectBaseDirectory(), namespace)

	// create new project directory
	logoutput.AddInfoEntryAndDetails("-> Create new project directory...", "Directory: ["+newProjectDir+"]")
	err = createNewProjectDirectory(newProjectDir)
	if err != nil {
		return err
	}
	logoutput.AddInfoEntry("-> Create new project directory...done")

	// copy necessary files
	logoutput.AddInfoEntryAndDetails("-> Start to copy templates to new project directory...", "Directory: ["+newProjectDir+"]")
	err = copyTemplatesToNewDirectory(newProjectDir, len(existingPvc) > 0, createDeploymentOnlyProject)
	if err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start to copy templates to new project directory...done")

	// add IP and namespace to IP configuration
	logoutput.AddInfoEntry("-> Start adding IP address to ipconfig file...")
	success, err := config.AddToIpConfigFile(namespace, ipAddress)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start adding IP address to ipconfig file...done")

	// processing cloud templates
	logoutput.AddInfoEntry("-> Start template processing: Jenkins cloud templates...")
	success, err = ProcessTemplateCloudTemplates(newProjectDir, selectedCloudTemplates)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start template processing: Jenkins cloud templates...done")

	// Replace Jenkins system message
	logoutput.AddInfoEntry("-> Start template processing: Jenkins system message...")
	success, err = ProcessTemplateJenkinsSystemMessage(newProjectDir, jenkinsSystemMsg)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start template processing: Jenkins system message...done")

	// Replace Jenkins seed job repository
	logoutput.AddInfoEntry("-> Start template processing: Jenkins Jobs repository...")
	success, err = ProcessTemplateJenkinsJobsRepo(newProjectDir, jobsCfgRepo)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start template processing: Jenkins Jobs repository...done")

	// Replace global configuration
	logoutput.AddInfoEntry("-> Start template processing: Global configuration...")
	success, err = replaceGlobalConfiguration(newProjectDir)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start template processing: Global configuration...done")

	// Replace namespace
	logoutput.AddInfoEntry("-> Start template processing: Namespaces...")
	success, err = ProcessTemplateNamespace(newProjectDir, namespace)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start template processing: Namespaces...done")

	// Replace ip address
	logoutput.AddInfoEntry("-> Start template processing: IP address...")
	success, err = ProcessTemplateIpAddress(newProjectDir, ipAddress)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start template processing: IP address...done")

	// Replace project directory with namespace name
	logoutput.AddInfoEntry("-> Start template processing: Project directory for JCasC...")
	success, err = replaceProjectDirectoryInTemplatesWithNamespace(newProjectDir, namespace)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start template processing: Project directory for JCasC...done")

	// Replace pvc name
	logoutput.AddInfoEntry("-> Start template processing: Persistent volume claim...")
	success, err = ProcessTemplatePvcExistingClaim(newProjectDir, existingPvc)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return err
	}
	logoutput.AddInfoEntry("-> Start template processing: Persistent volume claim...done")

	return err
}

// create new project directory
func createNewProjectDirectory(newProjectDir string) (err error) {
	log := logger.Log()

	// create directory
	err = os.MkdirAll(newProjectDir, os.ModePerm)
	if err != nil {
		logoutput.AddErrorEntryAndDetails("  -> Failed to create a new project directory.", err.Error())
		log.Error("[createNewProjectDirectory] Trying to create a new project directory [%v]...error. \n%v", newProjectDir, err)
		return err
	}

	return err
}

// copy files to new directory
func copyTemplatesToNewDirectory(projectDirectory string, copyPersistentVolume bool, createDeploymentOnlyProject bool) (err error) {
	var fileNamesToCopy []string

	// copy nginx-ingress-controller values.yaml
	fileNamesToCopy = append(fileNamesToCopy, constants.FilenameNginxIngressControllerHelmValues)

	// if it is not a deployment only project, copy more files
	if !createDeploymentOnlyProject {
		// copy Jenkins values.yaml
		fileNamesToCopy = append(fileNamesToCopy, constants.FilenameJenkinsHelmValues)
		// copy Jenkins JCasC config.yaml
		fileNamesToCopy = append(fileNamesToCopy, constants.FilenameJenkinsConfigurationAsCode)
		// copy existing PVC values.yaml
		if copyPersistentVolume {
			fileNamesToCopy = append(fileNamesToCopy, constants.FilenamePvcClaim)
		}
		// copy secrets to project
		if models.GetConfiguration().GlobalSecretsFile == "" {
			fileNamesToCopy = append(fileNamesToCopy, constants.FilenameSecrets)
		}
	}
	err = copyTemplates(fileNamesToCopy, projectDirectory)
	return err
}

// copy the filenames
func copyTemplates(fileNames []string, projectDirectory string) (err error) {
	log := logger.Log()
	for _, fileName := range fileNames {
		logoutput.AddInfoEntry("  -> Copy [" + fileName + "]...")
		_, err = files.CopyFile(
			files.AppendPath(models.GetProjectTemplateDirectory(), fileName),
			files.AppendPath(projectDirectory, fileName),
		)
		if err != nil {
			logoutput.AddErrorEntryAndDetails("  -> Copy ["+fileName+"]...failed. See errors.", err.Error())
			log.Error("Unable to copy [%v] to [%v] \n%v", fileName, projectDirectory, err)
			return err
		}
		logoutput.AddInfoEntry("  -> Copy [" + constants.FilenameNginxIngressControllerHelmValues + "]...done")
	}
	return err
}

func replaceGlobalConfiguration(projectDirectory string) (success bool, err error) {
	success, err = replaceGlobalConfigurationNginxIngressControllerHelmValues(projectDirectory)
	if !success || err != nil {
		return false, err
	}
	success, err = replaceGlobalConfigurationJCasCValues(projectDirectory)
	if !success || err != nil {
		return false, err
	}
	success, err = replaceGlobalConfigurationJenkinsHelmValues(projectDirectory)
	if !success || err != nil {
		return false, err
	}
	success, err = replaceGlobalConfigurationPvcValues(projectDirectory)
	if !success || err != nil {
		return false, err
	}
	return success, err
}

// Replace nginx ingress helm values.yaml
func replaceGlobalConfigurationNginxIngressControllerHelmValues(projectDirectory string) (success bool, err error) {
	var nginxHelmValuesFile = files.AppendPath(projectDirectory, constants.FilenameNginxIngressControllerHelmValues)
	if files.FileOrDirectoryExists(nginxHelmValuesFile) {
		// Replace global vars in nginx file
		// Jenkins related placeholder
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName)
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateJenkinsMasterDefaultUriPrefix, models.GetConfiguration().Jenkins.Helm.Master.DefaultUriPrefix)
		// Nginx ingress controller placeholder
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressDeploymentName, models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName)
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressControllerContainerImage, models.GetConfiguration().Nginx.Ingress.Controller.Container.Name)
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressControllerContainerPullSecrets, models.GetConfiguration().Nginx.Ingress.Controller.Container.PullSecret)
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressControllerContainerForNamespace, strconv.FormatBool(models.GetConfiguration().Nginx.Ingress.Controller.Container.Namespace))
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressAnnotationClass, models.GetConfiguration().Nginx.Ingress.AnnotationClass)
		// Loadbalancer placeholder
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerEnabled, strconv.FormatBool(models.GetConfiguration().LoadBalancer.Enabled))
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHttpPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.Http, 10))
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHttpTargetPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.HttpTarget, 10))
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHttpsPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.Https, 10))
		files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHttpsTargetPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.HttpsTarget, 10))
	}
	return true, err
}

// Replace Jenkins Helm default values
func replaceGlobalConfigurationJenkinsHelmValues(projectDirectory string) (success bool, err error) {
	var jenkinsHelmValues = files.AppendPath(projectDirectory, constants.FilenameJenkinsHelmValues)
	if files.FileOrDirectoryExists(jenkinsHelmValues) {
		// Jenkins
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDefaultUriPrefix, models.GetConfiguration().Jenkins.Helm.Master.DefaultUriPrefix)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDenyAnonymousReadAccess, models.GetConfiguration().Jenkins.Helm.Master.DenyAnonymousReadAccess)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterAdminPassword, models.GetConfiguration().Jenkins.Helm.Master.AdminPassword)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDefaultLabel, models.GetConfiguration().Jenkins.Helm.Master.Label)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterContainerImage, models.GetConfiguration().Jenkins.Helm.Master.Container.Image)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterContainerImageTag, models.GetConfiguration().Jenkins.Helm.Master.Container.ImageTag)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterContainerImagePullSecretName, models.GetConfiguration().Jenkins.Helm.Master.Container.PullSecretName)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterContainerPullPolicy, models.GetConfiguration().Jenkins.Helm.Master.Container.PullPolicy)
		// PVC
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplatePvcStorageClass, models.GetConfiguration().Jenkins.Helm.Master.Persistence.StorageClass)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplatePvcAccessMode, models.GetConfiguration().Jenkins.Helm.Master.Persistence.AccessMode)
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplatePvcStorageSize, models.GetConfiguration().Jenkins.Helm.Master.Persistence.Size)
		// JCasC
		files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsJcascConfigurationUrl, models.GetConfiguration().Jenkins.JCasC.ConfigurationUrl)
	}
	return true, err
}

// Replace Jenkins Configuration as Code default values
func replaceGlobalConfigurationJCasCValues(projectDirectory string) (success bool, err error) {
	var jcascFile = files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode)
	if files.FileOrDirectoryExists(jcascFile) {
		// Jenkins
		files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName)
		files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterDefaultUriPrefix, models.GetConfiguration().Jenkins.Helm.Master.DefaultUriPrefix)
		files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterDefaultLabel, models.GetConfiguration().Jenkins.Helm.Master.Label)
		files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsJobDslSeedJobScriptUrl, models.GetConfiguration().Jenkins.JobDSL.SeedJobScriptUrl)
		files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterAdminPasswordEncrypted, models.GetConfiguration().Jenkins.Helm.Master.AdminPasswordEncrypted)
		files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterUserPasswordEncrypted, models.GetConfiguration().Jenkins.Helm.Master.DefaultProjectUserPasswordEncrypted)
		// Kubernetes configuration
		files.ReplaceStringInFile(jcascFile, constants.TemplateKubernetesServerCertificate, models.GetConfiguration().Kubernetes.ServerCertificate)
		// CredentialIds
		files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIdKubernetesDockerRegistry, models.GetConfiguration().CredentialIds.DefaultDockerRegistry)
		files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIdMaven, models.GetConfiguration().CredentialIds.DefaultMavenRepository)
		files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIdNpm, models.GetConfiguration().CredentialIds.DefaultNpmRepository)
		files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIdVcs, models.GetConfiguration().CredentialIds.DefaultVcsRepository)
	}
	return true, err
}

// Replace PVC default values
func replaceGlobalConfigurationPvcValues(projectDirectory string) (success bool, err error) {
	var pvcFile = files.AppendPath(projectDirectory, constants.FilenamePvcClaim)
	if files.FileOrDirectoryExists(pvcFile) {
		// Jenkins
		files.ReplaceStringInFile(pvcFile, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName)
		// PVC
		files.ReplaceStringInFile(pvcFile, constants.TemplatePvcStorageSize, models.GetConfiguration().Jenkins.Helm.Master.Persistence.Size)
		files.ReplaceStringInFile(pvcFile, constants.TemplatePvcAccessMode, models.GetConfiguration().Jenkins.Helm.Master.Persistence.AccessMode)
		files.ReplaceStringInFile(pvcFile, constants.TemplatePvcStorageClass, models.GetConfiguration().Jenkins.Helm.Master.Persistence.StorageClass)
	}
	return true, err
}

// replace project directory with namespace name
func replaceProjectDirectoryInTemplatesWithNamespace(projectDirectory string, namespace string) (success bool, err error) {
	log := logger.Log()

	templateFiles := []string{
		files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode),
		files.AppendPath(projectDirectory, constants.FilenameJenkinsHelmValues),
	}

	for _, templateFile := range templateFiles {
		if files.FileOrDirectoryExists(templateFile) {
			successful, err := files.ReplaceStringInFile(templateFile, constants.TemplateProjectDirectory, namespace)
			if !successful || err != nil {
				logoutput.AddErrorEntryAndDetails("  -> Unable to replace project directory in file ["+templateFile+"]", err.Error())
				log.Error("[replaceProjectDirectoryInTemplatesWithNamespace] Unable to replace project directory in file [%v], \n%v", templateFile, err)
				return false, err
			}
		}
	}
	return true, err
}

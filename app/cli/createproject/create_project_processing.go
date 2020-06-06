package createproject

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os"
	"strconv"
)

// Processing project creation
func ProcessProjectCreate(namespace string, ipAddress string, jenkinsSystemMsg string, jobsCfgRepo string, existingPvc string, selectedCloudTemplates []string, createDeploymentOnlyProject bool) (info string, err error) {
	// calculate the target directory
	newProjectDir := files.AppendPath(models.GetProjectBaseDirectory(), namespace)

	// create new project directory
	infoLog, err := createNewProjectDirectory(newProjectDir)
	info = info + constants.NewLine + infoLog
	if err != nil {
		return info, err
	}

	// copy necessary files
	infoLog, err = copyTemplatesToNewDirectory(newProjectDir, len(existingPvc) > 0, createDeploymentOnlyProject)
	info = info + constants.NewLine + infoLog
	if err != nil {
		os.RemoveAll(newProjectDir)
		return info, err
	}

	// add IP and namespace to IP configuration
	success, err := config.AddToIpConfigFile(namespace, ipAddress)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return info, err
	}

	// processing cloud templates
	success, err = ProcessCloudTemplates(newProjectDir, selectedCloudTemplates)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return info, err
	}

	// Replace Jenkins system message
	success, err = ProcessJenkinsSystemMessage(newProjectDir, jenkinsSystemMsg)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return info, err
	}

	// Replace Jenkins seed job repository
	success, err = ProcessJenkinsJobsRepo(newProjectDir, jobsCfgRepo)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return info, err
	}

	success, err = replaceGlobalConfiguration(newProjectDir)
	if !success || err != nil {
		os.RemoveAll(newProjectDir)
		return info, err
	}
	return info, err
}

// create new project directory
func createNewProjectDirectory(newProjectDir string) (info string, err error) {
	log := logger.Log()
	log.Info("[createNewProjectDirectory] Trying to create a new project directory...")
	info = "Trying to create a new project directory..."

	// create directory
	err = os.Mkdir(newProjectDir, os.ModePerm)
	if err != nil {
		log.Error("[createNewProjectDirectory] Trying to create a new project directory [%v]...error. \n%v", newProjectDir, err)
		info = info + constants.NewLine + "Error while creating project directory."
		return info, err
	}
	// successful
	log.Info("[createNewProjectDirectory] Trying to create a new project directory...done")
	info = info + constants.NewLine + "Trying to create a new project directory...done"

	return info, err
}

// copy files to new directory
func copyTemplatesToNewDirectory(projectDirectory string, copyPersistentVolume bool, createDeploymentOnlyProject bool) (info string, err error) {
	templateDirectory := models.GetProjectTemplateDirectory()
	// copy nginx-ingress-controller values.yaml
	_, err = files.CopyFile(
		files.AppendPath(templateDirectory, constants.FilenameNginxIngressControllerHelmValues),
		files.AppendPath(projectDirectory, constants.FilenameNginxIngressControllerHelmValues),
	)
	if err != nil {
		return info, err
	}

	// if it is not a deployment only project, copy more files
	if !createDeploymentOnlyProject {
		// copy Jenkins values.yaml
		_, err = files.CopyFile(
			files.AppendPath(templateDirectory, constants.FilenameJenkinsHelmValues),
			files.AppendPath(projectDirectory, constants.FilenameJenkinsHelmValues),
		)
		if err != nil {
			return info, err
		}

		// copy Jenkins JCasC config.yaml
		_, err = files.CopyFile(
			files.AppendPath(templateDirectory, constants.FilenameJenkinsConfigurationAsCode),
			files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode),
		)
		if err != nil {
			return info, err
		}

		// copy existing PVC values.yaml
		if copyPersistentVolume {
			_, err = files.CopyFile(
				files.AppendPath(templateDirectory, constants.FilenamePvcClaim),
				files.AppendPath(projectDirectory, constants.FilenamePvcClaim),
			)
			if err != nil {
				return info, err
			}
		}

		// copy secrets to project
		if models.GetConfiguration().GlobalSecretsFile == "" {
			_, err = files.CopyFile(
				files.AppendPath(templateDirectory, constants.FilenameSecrets),
				files.AppendPath(projectDirectory, constants.FilenameSecrets),
			)
			if err != nil {
				return info, err
			}
		}
	}

	return info, err
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

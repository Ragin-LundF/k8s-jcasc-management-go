package createproject

import (
	"github.com/goware/prefixer"
	"io/ioutil"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os"
	"strconv"
	"strings"
)

// Processing project creation
func ProcessProjectCreate(namespace string, ipAddress string, jenkinsSystemMsg string, jobsCfgRepo string, existingPvc string, selectedCloudTemplates []string, createDeploymentOnlyProject bool) (err error) {
	// calculate the target directory
	newProjectDir := files.AppendPath(models.GetProjectBaseDirectory(), namespace)

	// create new project directory
	loggingstate.AddInfoEntryAndDetails("-> Create new project directory...", "Directory: ["+newProjectDir+"]")
	err = createNewProjectDirectory(newProjectDir)
	if err != nil {
		return err
	}
	loggingstate.AddInfoEntry("-> Create new project directory...done")

	// copy necessary files
	loggingstate.AddInfoEntryAndDetails("-> Start to copy templates to new project directory...", "Directory: ["+newProjectDir+"]")
	err = copyTemplatesToNewDirectory(newProjectDir, len(existingPvc) > 0, createDeploymentOnlyProject)
	if err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start to copy templates to new project directory...done")

	// add IP and namespace to IP configuration
	loggingstate.AddInfoEntry("-> Start adding IP address to ipconfig file...")
	success, err := config.AddToIpConfigFile(namespace, ipAddress)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start adding IP address to ipconfig file...done")

	// processing cloud templates
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins cloud templates...")
	success, err = processTemplateCloudTemplates(newProjectDir, selectedCloudTemplates)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins cloud templates...done")

	// Replace Jenkins system message
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins system message...")
	templateFiles := []string{files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode)}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplateJenkinsSystemMessage, jenkinsSystemMsg)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins system message...done")

	// Replace Jenkins seed job repository
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins Jobs repository...")
	templateFiles = []string{files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode)}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplateJobDefinitionRepository, jobsCfgRepo)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins Jobs repository...done")

	// Replace global configuration
	loggingstate.AddInfoEntry("-> Start template processing: Global configuration...")
	success, err = replaceGlobalConfiguration(newProjectDir)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Global configuration...done")

	// Replace namespace
	loggingstate.AddInfoEntry("-> Start template processing: Namespaces...")
	templateFiles = []string{
		files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode),
		files.AppendPath(newProjectDir, constants.FilenamePvcClaim),
		files.AppendPath(newProjectDir, constants.FilenameNginxIngressControllerHelmValues),
	}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplateNamespace, namespace)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Namespaces...done")

	// Replace ip address
	loggingstate.AddInfoEntry("-> Start template processing: IP address...")
	templateFiles = []string{
		files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode),
		files.AppendPath(newProjectDir, constants.FilenameNginxIngressControllerHelmValues),
	}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplatePublicIpAddress, ipAddress)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: IP address...done")

	// Replace project directory with namespace name
	loggingstate.AddInfoEntry("-> Start template processing: Project directory for JCasC...")

	templateFiles = []string{
		files.AppendPath(newProjectDir, constants.FilenameJenkinsConfigurationAsCode),
		files.AppendPath(newProjectDir, constants.FilenameJenkinsHelmValues),
	}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplateProjectDirectory, namespace)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Project directory for JCasC...done")

	// Replace existing pvc in Jenkins
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins existing persistent volume claim...")
	templateFiles = []string{files.AppendPath(newProjectDir, constants.FilenameJenkinsHelmValues)}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplatePvcExistingVolumeClaim, existingPvc)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Jenkins existing persistent volume claim...done")

	// Replace existing pvc in Jenkins
	loggingstate.AddInfoEntry("-> Start template processing: Persistent volume claim...")
	templateFiles = []string{files.AppendPath(newProjectDir, constants.FilenamePvcClaim)}
	success, err = replacePlaceholderInTemplates(templateFiles, constants.TemplatePvcName, existingPvc)
	if !success || err != nil {
		_ = os.RemoveAll(newProjectDir)
		return err
	}
	loggingstate.AddInfoEntry("-> Start template processing: Persistent volume claim...done")

	return err
}

// create new project directory
func createNewProjectDirectory(newProjectDir string) (err error) {
	log := logger.Log()

	// create directory
	err = os.MkdirAll(newProjectDir, os.ModePerm)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Failed to create a new project directory.", err.Error())
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
		loggingstate.AddInfoEntry("  -> Copy [" + fileName + "]...")
		_, err = files.CopyFile(
			files.AppendPath(models.GetProjectTemplateDirectory(), fileName),
			files.AppendPath(projectDirectory, fileName),
		)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Copy ["+fileName+"]...failed. See errors.", err.Error())
			log.Error("Unable to copy [%v] to [%v] \n%v", fileName, projectDirectory, err)
			return err
		}
		loggingstate.AddInfoEntry("  -> Copy [" + constants.FilenameNginxIngressControllerHelmValues + "]...done")
	}
	return err
}

// add cloud templates to project template
func processTemplateCloudTemplates(projectDirectory string, cloudTemplateFiles []string) (success bool, err error) {
	log := logger.Log()
	targetFile := files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode)
	// if file exists -> try to replace files
	if files.FileOrDirectoryExists(targetFile) {
		// first check if there are templates which should be processed
		if cap(cloudTemplateFiles) > 0 {
			// prepare vars and directory
			var cloudTemplateContent string
			var cloudTemplatePath = files.AppendPath(models.GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)

			// first read every template into a variable
			for _, cloudTemplate := range cloudTemplateFiles {
				cloudTemplateFileWithPath := files.AppendPath(cloudTemplatePath, cloudTemplate)
				read, err := ioutil.ReadFile(cloudTemplateFileWithPath)
				if err != nil {
					loggingstate.AddErrorEntryAndDetails("  -> Unable to read cloud template ["+cloudTemplateFileWithPath+"]", err.Error())
					log.Error("[processTemplateCloudTemplates] Unable to read cloud template [%v] \n%v", cloudTemplateFileWithPath, err)
					return false, err
				}
				cloudTemplateContent = cloudTemplateContent + constants.NewLine + string(read)
			}

			// Prefix the lines with spaces for correct yaml template
			prefixReader := prefixer.New(strings.NewReader(cloudTemplateContent), "          ")
			buffer, _ := ioutil.ReadAll(prefixReader)
			cloudTemplateContent = string(buffer)

			// replace target template
			success, err = files.ReplaceStringInFile(targetFile, constants.TemplateJenkinsCloudTemplates, cloudTemplateContent)
			if !success || err != nil {
				loggingstate.AddErrorEntryAndDetails("  -> Unable to replace ["+constants.TemplateJenkinsCloudTemplates+"] in ["+constants.FilenameJenkinsConfigurationAsCode+"]", err.Error())
				return false, err
			}
		} else {
			// replace placeholder
			success, err = files.ReplaceStringInFile(targetFile, constants.TemplateJenkinsCloudTemplates, "")
			if !success || err != nil {
				loggingstate.AddErrorEntryAndDetails("  -> Unable to replace ["+constants.TemplateJenkinsCloudTemplates+"] in ["+constants.FilenameJenkinsConfigurationAsCode+"]", err.Error())
				return false, err
			}
		}
	}
	return true, err
}

// delegetaion method for replacement of global configuration
func replaceGlobalConfiguration(projectDirectory string) (success bool, err error) {
	log := logger.Log()

	success, err = replaceGlobalConfigurationNginxIngressControllerHelmValues(projectDirectory)
	if !success || err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to replace global nginx-ingress-controller Helm values...abort", err.Error())
		log.Error("Unable to replace global nginx-ingress-controller Helm values...abort \n%v", err.Error())
		return false, err
	}
	success, err = replaceGlobalConfigurationJCasCValues(projectDirectory)
	if !success || err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to replace global JCasc values...abort", err.Error())
		log.Error("Unable to replace global JCasc values...abort \n%v", err.Error())
		return false, err
	}
	success, err = replaceGlobalConfigurationJenkinsHelmValues(projectDirectory)
	if !success || err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to replace global Jenkins Helm values...abort", err.Error())
		log.Error("Unable to replace global Jenkins Helm values...abort \n%v", err.Error())
		return false, err
	}
	success, err = replaceGlobalConfigurationPvcValues(projectDirectory)
	if !success || err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to replace global PVC values...abort", err.Error())
		log.Error("Unable to replace global PVC values...abort \n%v", err.Error())
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
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateJenkinsMasterDefaultUriPrefix, models.GetConfiguration().Jenkins.Helm.Master.DefaultUriPrefix); !success {
			return success, err
		}
		// Nginx ingress controller placeholder
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressDeploymentName, models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressControllerContainerImage, models.GetConfiguration().Nginx.Ingress.Controller.Container.Name); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressControllerContainerPullSecrets, models.GetConfiguration().Nginx.Ingress.Controller.Container.PullSecret); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressControllerContainerForNamespace, strconv.FormatBool(models.GetConfiguration().Nginx.Ingress.Controller.Container.Namespace)); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxIngressAnnotationClass, models.GetConfiguration().Nginx.Ingress.AnnotationClass); !success {
			return success, err
		}
		// Loadbalancer placeholder
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerEnabled, strconv.FormatBool(models.GetConfiguration().LoadBalancer.Enabled)); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHttpPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.Http, 10)); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHttpTargetPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.HttpTarget, 10)); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHttpsPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.Https, 10)); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHttpsTargetPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.HttpsTarget, 10)); !success {
			return success, err
		}
	}
	return true, err
}

// Replace Jenkins Helm default values
func replaceGlobalConfigurationJenkinsHelmValues(projectDirectory string) (success bool, err error) {
	var jenkinsHelmValues = files.AppendPath(projectDirectory, constants.FilenameJenkinsHelmValues)
	if files.FileOrDirectoryExists(jenkinsHelmValues) {
		// Jenkins
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDefaultUriPrefix, models.GetConfiguration().Jenkins.Helm.Master.DefaultUriPrefix); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDenyAnonymousReadAccess, models.GetConfiguration().Jenkins.Helm.Master.DenyAnonymousReadAccess); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterAdminPassword, models.GetConfiguration().Jenkins.Helm.Master.AdminPassword); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDefaultLabel, models.GetConfiguration().Jenkins.Helm.Master.Label); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterContainerImage, models.GetConfiguration().Jenkins.Helm.Master.Container.Image); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterContainerImageTag, models.GetConfiguration().Jenkins.Helm.Master.Container.ImageTag); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterContainerImagePullSecretName, models.GetConfiguration().Jenkins.Helm.Master.Container.PullSecretName); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterContainerPullPolicy, models.GetConfiguration().Jenkins.Helm.Master.Container.PullPolicy); !success {
			return success, err
		}
		// PVC
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplatePvcStorageClass, models.GetConfiguration().Jenkins.Helm.Master.Persistence.StorageClass); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplatePvcAccessMode, models.GetConfiguration().Jenkins.Helm.Master.Persistence.AccessMode); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplatePvcStorageSize, models.GetConfiguration().Jenkins.Helm.Master.Persistence.Size); !success {
			return success, err
		}
		// JCasC
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsJcascConfigurationUrl, models.GetConfiguration().Jenkins.JCasC.ConfigurationUrl); !success {
			return success, err
		}
	}
	return true, err
}

// Replace Jenkins Configuration as Code default values
func replaceGlobalConfigurationJCasCValues(projectDirectory string) (success bool, err error) {
	var jcascFile = files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode)
	if files.FileOrDirectoryExists(jcascFile) {
		// Jenkins
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterDefaultUriPrefix, models.GetConfiguration().Jenkins.Helm.Master.DefaultUriPrefix); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterDefaultLabel, models.GetConfiguration().Jenkins.Helm.Master.Label); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsJobDslSeedJobScriptUrl, models.GetConfiguration().Jenkins.JobDSL.SeedJobScriptUrl); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterAdminPasswordEncrypted, models.GetConfiguration().Jenkins.Helm.Master.AdminPasswordEncrypted); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterUserPasswordEncrypted, models.GetConfiguration().Jenkins.Helm.Master.DefaultProjectUserPasswordEncrypted); !success {
			return success, err
		}
		// Kubernetes configuration
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateKubernetesServerCertificate, models.GetConfiguration().Kubernetes.ServerCertificate); !success {
			return success, err
		}
		// CredentialIds
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIdKubernetesDockerRegistry, models.GetConfiguration().CredentialIds.DefaultDockerRegistry); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIdMaven, models.GetConfiguration().CredentialIds.DefaultMavenRepository); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIdNpm, models.GetConfiguration().CredentialIds.DefaultNpmRepository); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIdVcs, models.GetConfiguration().CredentialIds.DefaultVcsRepository); !success {
			return success, err
		}
	}
	return true, err
}

// Replace PVC default values
func replaceGlobalConfigurationPvcValues(projectDirectory string) (success bool, err error) {
	var pvcFile = files.AppendPath(projectDirectory, constants.FilenamePvcClaim)
	if files.FileOrDirectoryExists(pvcFile) {
		// Jenkins
		if success, err = files.ReplaceStringInFile(pvcFile, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName); !success {
			return success, err
		}
		// PVC
		if success, err = files.ReplaceStringInFile(pvcFile, constants.TemplatePvcStorageSize, models.GetConfiguration().Jenkins.Helm.Master.Persistence.Size); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(pvcFile, constants.TemplatePvcAccessMode, models.GetConfiguration().Jenkins.Helm.Master.Persistence.AccessMode); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(pvcFile, constants.TemplatePvcStorageClass, models.GetConfiguration().Jenkins.Helm.Master.Persistence.StorageClass); !success {
			return success, err
		}
	}
	return true, err
}

// replace project directory with namespace name
func replacePlaceholderInTemplates(templateFiles []string, placeholder string, newValue string) (success bool, err error) {
	log := logger.Log()

	for _, templateFile := range templateFiles {
		if files.FileOrDirectoryExists(templateFile) {
			successful, err := files.ReplaceStringInFile(templateFile, placeholder, newValue)
			if !successful || err != nil {
				loggingstate.AddErrorEntryAndDetails("  -> Unable to replace ["+placeholder+"] in file ["+templateFile+"]", err.Error())
				log.Error("[replacePlaceholderInTemplates] Unable to replace [%v] in file [%v], \n%v", placeholder, templateFile, err)
				return false, err
			}
		}
	}
	return true, err
}

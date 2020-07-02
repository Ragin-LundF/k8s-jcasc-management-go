package createprojectactions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
)

// ActionReplaceGlobalConfigJCasCValues replace Jenkins Configuration as Code default values
func ActionReplaceGlobalConfigJCasCValues(projectDirectory string) (success bool, err error) {
	var jcascFile = files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode)
	if files.FileOrDirectoryExists(jcascFile) {
		// Jenkins
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterDefaultURIPrefix, models.GetConfiguration().Jenkins.Helm.Master.DefaultURIPrefix); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsMasterDefaultLabel, models.GetConfiguration().Jenkins.Helm.Master.Label); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateJenkinsJobDslSeedJobScriptURL, models.GetConfiguration().Jenkins.JobDSL.SeedJobScriptURL); !success {
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
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIDKubernetesDockerRegistry, models.GetConfiguration().CredentialIds.DefaultDockerRegistry); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIDMaven, models.GetConfiguration().CredentialIds.DefaultMavenRepository); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIDNpm, models.GetConfiguration().CredentialIds.DefaultNpmRepository); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jcascFile, constants.TemplateCredentialsIDVcs, models.GetConfiguration().CredentialIds.DefaultVcsRepository); !success {
			return success, err
		}
	}
	return true, nil
}

package createprojectactions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
)

// ActionReplaceGlobalConfigJenkinsHelmValues replaces Jenkins Helm default values
func ActionReplaceGlobalConfigJenkinsHelmValues(projectDirectory string) (success bool, err error) {
	var jenkinsHelmValues = files.AppendPath(projectDirectory, constants.FilenameJenkinsHelmValues)
	if files.FileOrDirectoryExists(jenkinsHelmValues) {
		// Jenkins
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsMasterDefaultURIPrefix, models.GetConfiguration().Jenkins.Helm.Master.DefaultURIPrefix); !success {
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
		if success, err = files.ReplaceStringInFile(jenkinsHelmValues, constants.TemplateJenkinsJcascConfigurationURL, models.GetConfiguration().Jenkins.JCasC.ConfigurationURL); !success {
			return success, err
		}
	}
	return true, nil
}

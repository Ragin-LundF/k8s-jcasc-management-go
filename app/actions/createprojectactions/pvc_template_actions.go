package createprojectactions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
)

// ActionReplaceGlobalConfigPvcValues replaces PVC default values
func ActionReplaceGlobalConfigPvcValues(projectDirectory string) (success bool, err error) {
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
	return true, nil
}

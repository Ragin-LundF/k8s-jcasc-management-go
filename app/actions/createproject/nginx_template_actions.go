package createproject

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"strconv"
)

// Replace nginx ingress helm values.yaml
func ActionReplaceGlobalConfigNginxIngressCtrlHelmValues(projectDirectory string) (success bool, err error) {
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
	return true, nil
}

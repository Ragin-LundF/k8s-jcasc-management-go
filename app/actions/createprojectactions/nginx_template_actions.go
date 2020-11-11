package createprojectactions

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"strconv"
)

// ActionReplaceGlobalConfigNginxIngressCtrlHelmValues replaces nginx ingress helm values.yaml
func ActionReplaceGlobalConfigNginxIngressCtrlHelmValues(projectDirectory string) (success bool, err error) {
	var nginxHelmValuesFile = files.AppendPath(projectDirectory, constants.FilenameNginxIngressControllerHelmValues)
	if files.FileOrDirectoryExists(nginxHelmValuesFile) {
		// Replace global vars in nginx file
		// Jenkins related placeholder
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateJenkinsMasterDeploymentName, models.GetConfiguration().Jenkins.Helm.Master.DeploymentName); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateJenkinsMasterDefaultURIPrefix, models.GetConfiguration().Jenkins.Helm.Master.DefaultURIPrefix); !success {
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
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHTTPPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.HTTP, 10)); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHTTPTargetPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.HTTPTarget, 10)); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHTTPSPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.HTTPS, 10)); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerHTTPSTargetPort, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Port.HTTPSTarget, 10)); !success {
			return success, err
		}
		// Loadbalancer annotations placeholder
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerAnnotationsEnabled, strconv.FormatBool(models.GetConfiguration().LoadBalancer.Annotations.Enabled)); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerAnnotationsExtDnsHostname, models.GetConfiguration().LoadBalancer.Annotations.ExtDNS.Hostname); !success {
			return success, err
		}
		if success, err = files.ReplaceStringInFile(nginxHelmValuesFile, constants.TemplateNginxLoadbalancerAnnotationsExtDnsTtl, strconv.FormatUint(models.GetConfiguration().LoadBalancer.Annotations.ExtDNS.Ttl, 10)); !success {
			return success, err
		}
	}
	return true, nil
}

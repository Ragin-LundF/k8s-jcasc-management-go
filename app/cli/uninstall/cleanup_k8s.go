package uninstall

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/logger"
)

func CleanupK8sNginxIngressController(namespace string) (info string, err error) {
	// Nginx Ingress Roles
	infoLog, err := CleanupNginxIngressRoles(namespace)
	info = info + constants.NewLine + infoLog
	if err != nil {
		return info, err
	}

	return info, err
}

func CleanupNginxIngressRoles(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[CleanupNginxIngressRoles] Start to cleanup nginx-ingress Roles for namespace [" + namespace + "]...")

	// Search for roles with deployment name
	kubectlCmdArgs := []string{
		"role",
		"-n", namespace,
	}
	kubectlCmdOutput, infoLog, err := kubectl.ExecutorKubectl("get", kubectlCmdArgs)
	info = info + constants.NewLine + infoLog
	if err != nil {
		err = errors.New("[CleanupNginxIngressRoles] Unable to get roles for namespace [" + namespace + "]")
		log.Error("[CleanupNginxIngressRoles] Unable to get roles for namespace [" + namespace + "]")
		return info, err
	}

	kubectl.FindFieldValuesInKubectlOutput(kubectlCmdOutput, constants.KubectlFieldName)

	log.Info("[CleanupNginxIngressRoles] Start to cleanup nginx-ingress Roles for namespace [" + namespace + "]...")
	return info, err
}

package uninstall

import "k8s-management-go/app/constants"

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

	return info, err
}

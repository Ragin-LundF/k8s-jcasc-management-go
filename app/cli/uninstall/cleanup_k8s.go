package uninstall

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models/config"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/logger"
	"strings"
)

func CleanupK8sNginxIngressController(namespace string) (info string, err error) {
	// Nginx Ingress Ctrl Roles
	infoLog, _ := CleanupNginxIngressCtrlRoles(namespace)
	info = info + constants.NewLine + infoLog
	// Nginx Ingress Ctrl RoleBindings
	infoLog, _ = CleanupNginxIngressCtrlRoleBindings(namespace)
	info = info + constants.NewLine + infoLog
	// Nginx Ingress Ctrl ServiceAccounts
	infoLog, _ = CleanupNginxIngressCtrlServiceAccounts(namespace)
	info = info + constants.NewLine + infoLog
	// Nginx Ingress Ctrl ClusterRoles
	infoLog, _ = CleanupNginxIngressCtrlClusterRoles(namespace)
	info = info + constants.NewLine + infoLog
	// Nginx Ingress Ctrl ClusterRoleBindings
	infoLog, _ = CleanupNginxIngressCtrlClusterRoleBinding(namespace)
	info = info + constants.NewLine + infoLog
	// Nginx Ingress Ctrl ingress routes
	infoLog, _ = CleanupNginxIngressCtrlIngress(namespace)
	info = info + constants.NewLine + infoLog

	return info, err
}

// uninstall nginx-ingress roles
func CleanupNginxIngressCtrlRoles(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[CleanupNginxIngressCtrlRoles] Start to cleanup nginx-ingress Roles for namespace [" + namespace + "]...")

	searchValue := config.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	info, err = deleteFromKubernetes(namespace, "role", searchValue)

	log.Info("[CleanupNginxIngressCtrlRoles] Cleanup nginx-ingress Roles for namespace [" + namespace + "] done.")
	return info, err
}

// uninstall nginx-ingress role bindings
func CleanupNginxIngressCtrlRoleBindings(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[CleanupNginxIngressCtrlRoleBindings] Start to cleanup nginx-ingress RoleBindings for namespace [" + namespace + "]...")

	searchValue := config.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	info, err = deleteFromKubernetes(namespace, "rolebindings", searchValue)

	log.Info("[CleanupNginxIngressCtrlRoleBindings] Cleanup nginx-ingress RoleBindings for namespace [" + namespace + "] done.")
	return info, err
}

// uninstall nginx-ingress service accounts
func CleanupNginxIngressCtrlServiceAccounts(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[CleanupNginxIngressCtrlServiceAccounts] Start to cleanup nginx-ingress ServiceAccounts for namespace [" + namespace + "]...")

	searchValue := config.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	info, err = deleteFromKubernetes(namespace, "sa", searchValue)

	log.Info("[CleanupNginxIngressCtrlServiceAccounts] Cleanup nginx-ingress ServiceAccounts for namespace [" + namespace + "] done.")
	return info, err
}

// uninstall nginx-ingress clusterroles
func CleanupNginxIngressCtrlClusterRoles(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[CleanupNginxIngressCtrlClusterRoles] Start to cleanup nginx-ingress ClusterRoles for namespace [" + namespace + "]...")

	searchValue := config.GetConfiguration().Nginx.Ingress.Controller.DeploymentName + "-clusterrole-" + namespace
	info, err = deleteFromKubernetes(namespace, "clusterrole", searchValue)

	log.Info("[CleanupNginxIngressCtrlClusterRoles] Cleanup nginx-ingress ClusterRoles for namespace [" + namespace + "] done.")
	return info, err
}

// uninstall nginx-ingress clusterrole binding
func CleanupNginxIngressCtrlClusterRoleBinding(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[CleanupNginxIngressCtrlClusterRoleBinding] Start to cleanup nginx-ingress ClusterRoleBinding for namespace [" + namespace + "]...")

	searchValue := config.GetConfiguration().Nginx.Ingress.Controller.DeploymentName + "-clusterrole-nisa-binding-" + namespace
	info, err = deleteFromKubernetes(namespace, "clusterrolebinding", searchValue)

	log.Info("[CleanupNginxIngressCtrlClusterRoleBinding] Cleanup nginx-ingress ClusterRoleBinding for namespace [" + namespace + "] done.")
	return info, err
}

// uninstall nginx-ingress ingress
func CleanupNginxIngressCtrlIngress(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[CleanupNginxIngressCtrlClusterRoleBinding] Start to cleanup nginx-ingress ingress routes for namespace [" + namespace + "]...")

	searchValue := config.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	info, err = deleteFromKubernetes(namespace, "ingress", searchValue)

	log.Info("[CleanupNginxIngressCtrlClusterRoleBinding] Cleanup nginx-ingress ingress routes for namespace [" + namespace + "] done.")
	return info, err
}

// generic function to delete role, rolebinding, sa... from kubectl
func deleteFromKubernetes(namespace string, kubernetesType string, filterValue string) (info string, err error) {
	log := logger.Log()

	// Search for roles with deployment name
	kubectlCmdArgs := []string{
		kubernetesType,
		"-n", namespace,
	}
	kubectlCmdOutput, infoLog, err := kubectl.ExecutorKubectl("get", kubectlCmdArgs)
	info = info + constants.NewLine + infoLog
	if err != nil {
		err = errors.New("[deleteFromKubernetes] Unable to get [" + kubernetesType + "] for namespace [" + namespace + "]")
		log.Error("[deleteFromKubernetes] Unable to get [" + kubernetesType + "] for namespace [" + namespace + "]")
		return info, err
	}

	// extract NAME values from kubectl output
	fieldValues, err := kubectl.FindFieldValuesInKubectlOutput(kubectlCmdOutput, constants.KubectlFieldName)
	if err != nil {
		log.Error("[deleteFromKubernetes] Unable to find ["+kubernetesType+"] in namespace ["+namespace+"]... %v\n", err)
	}

	// found some roles, try to delete them
	if len(fieldValues) > 0 {
		// first find the relevant roles
		var fieldValuesToDelete []string
		for _, fieldValue := range fieldValues {
			if strings.Contains(fieldValue, filterValue) {
				fieldValuesToDelete = append(fieldValuesToDelete, filterValue)
			}
		}

		// found relevant roles, now uninstall them
		if len(fieldValuesToDelete) > 0 {
			kubectlUninstallCmdArgs := []string{
				"-n", namespace,
				kubernetesType,
			}
			for _, fieldValueToDelete := range fieldValuesToDelete {
				kubectlUninstallCmdArgs = append(kubectlUninstallCmdArgs, fieldValueToDelete)
			}

			// Execute delete command
			_, infoLog, err := kubectl.ExecutorKubectl("delete", kubectlUninstallCmdArgs)
			info = info + constants.NewLine + infoLog
			if err != nil {
				log.Error("[deleteFromKubernetes] Unable to uninstall nginx-ingress-controller [" + kubernetesType + "] from namespace [" + namespace + "]")
			}
		}
	}
	return info, err
}

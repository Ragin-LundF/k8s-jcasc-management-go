package uninstall

import (
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/logger"
	"strings"
)

// This delegate method tries to cleanup everything. It ignores possible errors!
// They will be logged, but it has no impact to the workflow
func CleanupK8sNginxIngressController(namespace string) {
	// Nginx Ingress Ctrl Roles
	CleanupNginxIngressCtrlRoles(namespace)
	// Nginx Ingress Ctrl RoleBindings
	CleanupNginxIngressCtrlRoleBindings(namespace)
	// Nginx Ingress Ctrl ServiceAccounts
	CleanupNginxIngressCtrlServiceAccounts(namespace)
	// Nginx Ingress Ctrl ClusterRoles
	CleanupNginxIngressCtrlClusterRoles(namespace)
	// Nginx Ingress Ctrl ClusterRoleBindings
	CleanupNginxIngressCtrlClusterRoleBinding(namespace)
	// Nginx Ingress Ctrl ingress routes
	CleanupNginxIngressCtrlIngress(namespace)
}

// uninstall nginx-ingress roles
func CleanupNginxIngressCtrlRoles(namespace string) {
	log := logger.Log()
	log.Infof("[CleanupNginxIngressCtrlRoles] Start to cleanup nginx-ingress Roles for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress Roles for namespace [" + namespace + "]...")

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "role", searchValue)

	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress Roles for namespace [" + namespace + "]...done")
	log.Infof("[CleanupNginxIngressCtrlRoles] Cleanup nginx-ingress Roles for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress role bindings
func CleanupNginxIngressCtrlRoleBindings(namespace string) {
	log := logger.Log()
	log.Infof("[CleanupNginxIngressCtrlRoleBindings] Start to cleanup nginx-ingress RoleBindings for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress RoleBindings for namespace [" + namespace + "]...")

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "rolebindings", searchValue)

	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress RoleBindings for namespace [" + namespace + "]...done")
	log.Infof("[CleanupNginxIngressCtrlRoleBindings] Cleanup nginx-ingress RoleBindings for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress service accounts
func CleanupNginxIngressCtrlServiceAccounts(namespace string) {
	log := logger.Log()
	log.Infof("[CleanupNginxIngressCtrlServiceAccounts] Start to cleanup nginx-ingress ServiceAccounts for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress ServiceAccounts for namespace [" + namespace + "]...")

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "sa", searchValue)

	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress ServiceAccounts for namespace [" + namespace + "]...done")
	log.Infof("[CleanupNginxIngressCtrlServiceAccounts] Cleanup nginx-ingress ServiceAccounts for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress clusterroles
func CleanupNginxIngressCtrlClusterRoles(namespace string) {
	log := logger.Log()
	log.Infof("[CleanupNginxIngressCtrlClusterRoles] Start to cleanup nginx-ingress ClusterRoles for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress ClusterRoles for namespace [" + namespace + "]...")

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName + "-clusterrole-" + namespace
	_ = deleteFromKubernetes(namespace, "clusterrole", searchValue)

	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress ClusterRoles for namespace [" + namespace + "]...donen")
	log.Infof("[CleanupNginxIngressCtrlClusterRoles] Cleanup nginx-ingress ClusterRoles for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress clusterrole binding
func CleanupNginxIngressCtrlClusterRoleBinding(namespace string) {
	log := logger.Log()
	log.Infof("[CleanupNginxIngressCtrlClusterRoleBinding] Start to cleanup nginx-ingress ClusterRoleBinding for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress ClusterRoleBinding for namespace [" + namespace + "]...")

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName + "-clusterrole-nisa-binding-" + namespace
	_ = deleteFromKubernetes(namespace, "clusterrolebinding", searchValue)

	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress ClusterRoleBinding for namespace [" + namespace + "]...done")
	log.Infof("[CleanupNginxIngressCtrlClusterRoleBinding] Cleanup nginx-ingress ClusterRoleBinding for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress ingress
func CleanupNginxIngressCtrlIngress(namespace string) {
	log := logger.Log()
	log.Infof("[CleanupNginxIngressCtrlClusterRoleBinding] Start to cleanup nginx-ingress ingress routes for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress ingress routes for namespace [" + namespace + "]...")

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "ingress", searchValue)

	loggingstate.AddInfoEntry("  -> Start to cleanup nginx-ingress ingress routes for namespace [" + namespace + "]...done")
	log.Infof("[CleanupNginxIngressCtrlClusterRoleBinding] Cleanup nginx-ingress ingress routes for namespace [%s] done.", namespace)
}

// generic function to delete role, rolebinding, sa... from kubectl
func deleteFromKubernetes(namespace string, kubernetesType string, filterValue string) (err error) {
	log := logger.Log()

	// Search for roles with deployment name
	kubectlCmdArgs := []string{
		kubernetesType,
		"-n", namespace,
	}
	kubectlCmdOutput, err := kubectl.ExecutorKubectl("get", kubectlCmdArgs)
	if err != nil {
		loggingstate.AddErrorEntry("  -> Unable to get [" + kubernetesType + "] for namespace [" + namespace + "]")
		log.Errorf("[deleteFromKubernetes] Unable to get [%s] for namespace [%s]", kubernetesType, namespace)
		return err
	}

	// extract NAME values from kubectl output
	fieldValues, err := kubectl.FindFieldValuesInKubectlOutput(kubectlCmdOutput, constants.KubectlFieldName)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("Unable to find ["+kubernetesType+"] in namespace ["+namespace+"]...", err.Error())
		log.Errorf("[deleteFromKubernetes] Unable to find ["+kubernetesType+"] in namespace ["+namespace+"]... %s\n", err.Error())
		return err
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
			_, err := kubectl.ExecutorKubectl("delete", kubectlUninstallCmdArgs)
			if err != nil {
				loggingstate.AddErrorEntryAndDetails("  -> Unable to uninstall nginx-ingress-controller ["+kubernetesType+"] from namespace ["+namespace+"]", err.Error())
				log.Errorf("[deleteFromKubernetes] Unable to uninstall nginx-ingress-controller [" + kubernetesType + "] from namespace [" + namespace + "]")
				return err
			}
		}
	}
	return nil
}

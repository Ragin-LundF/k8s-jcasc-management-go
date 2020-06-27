package uninstall_actions

import (
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// This delegate method tries to cleanup everything. It ignores possible errors!
// They will be logged, but it has no impact to the workflow
func ActionCleanupK8sNginxIngressController(namespace string) {
	// Nginx Ingress Ctrl Roles
	ActionCleanupNginxIngressCtrlRoles(namespace)
	// Nginx Ingress Ctrl RoleBindings
	ActionCleanupNginxIngressCtrlRoleBindings(namespace)
	// Nginx Ingress Ctrl ServiceAccounts
	ActionCleanupNginxIngressCtrlServiceAccounts(namespace)
	// Nginx Ingress Ctrl ClusterRoles
	ActionCleanupNginxIngressCtrlClusterRoles(namespace)
	// Nginx Ingress Ctrl ClusterRoleBindings
	ActionCleanupNginxIngressCtrlClusterRoleBinding(namespace)
	// Nginx Ingress Ctrl ingress routes
	ActionCleanupNginxIngressCtrlIngress(namespace)
}

// uninstall nginx-ingress roles
func ActionCleanupNginxIngressCtrlRoles(namespace string) {
	log := logger.Log()
	log.Infof("[ActionCleanupNginxIngressCtrlRoles] Start to cleanup nginx-ingress Roles for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress Roles for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "role", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress Roles for namespace [%s]...done", namespace))
	log.Infof("[ActionCleanupNginxIngressCtrlRoles] Cleanup nginx-ingress Roles for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress role bindings
func ActionCleanupNginxIngressCtrlRoleBindings(namespace string) {
	log := logger.Log()
	log.Infof("[ActionCleanupNginxIngressCtrlRoleBindings] Start to cleanup nginx-ingress RoleBindings for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress RoleBindings for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "rolebindings", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress RoleBindings for namespace [%s]...done", namespace))
	log.Infof("[ActionCleanupNginxIngressCtrlRoleBindings] Cleanup nginx-ingress RoleBindings for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress service accounts
func ActionCleanupNginxIngressCtrlServiceAccounts(namespace string) {
	log := logger.Log()
	log.Infof("[ActionCleanupNginxIngressCtrlServiceAccounts] Start to cleanup nginx-ingress ServiceAccounts for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ServiceAccounts for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "sa", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ServiceAccounts for namespace [%s]...done", namespace))
	log.Infof("[ActionCleanupNginxIngressCtrlServiceAccounts] Cleanup nginx-ingress ServiceAccounts for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress clusterroles
func ActionCleanupNginxIngressCtrlClusterRoles(namespace string) {
	log := logger.Log()
	log.Infof("[ActionCleanupNginxIngressCtrlClusterRoles] Start to cleanup nginx-ingress ClusterRoles for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ClusterRoles for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName + "-clusterrole-" + namespace
	_ = deleteFromKubernetes(namespace, "clusterrole", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ClusterRoles for namespace [%s]...donen", namespace))
	log.Infof("[ActionCleanupNginxIngressCtrlClusterRoles] Cleanup nginx-ingress ClusterRoles for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress clusterrole binding
func ActionCleanupNginxIngressCtrlClusterRoleBinding(namespace string) {
	log := logger.Log()
	log.Infof("[ActionCleanupNginxIngressCtrlClusterRoleBinding] Start to cleanup nginx-ingress ClusterRoleBinding for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ClusterRoleBinding for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName + "-clusterrole-nisa-binding-" + namespace
	_ = deleteFromKubernetes(namespace, "clusterrolebinding", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ClusterRoleBinding for namespace [%s]...done", namespace))
	log.Infof("[ActionCleanupNginxIngressCtrlClusterRoleBinding] Cleanup nginx-ingress ClusterRoleBinding for namespace [%s] done.", namespace)
}

// uninstall nginx-ingress ingress
func ActionCleanupNginxIngressCtrlIngress(namespace string) {
	log := logger.Log()
	log.Infof("[ActionCleanupNginxIngressCtrlClusterRoleBinding] Start to cleanup nginx-ingress ingress routes for namespace [%s]...", namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ingress routes for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "ingress", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ingress routes for namespace [%s]...done", namespace))
	log.Infof("[ActionCleanupNginxIngressCtrlClusterRoleBinding] Cleanup nginx-ingress ingress routes for namespace [%s] done.", namespace)
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
		loggingstate.AddErrorEntry(fmt.Sprintf("  -> Unable to get [%s] for namespace [%s]", kubernetesType, namespace))
		log.Errorf("[deleteFromKubernetes] Unable to get [%s] for namespace [%s]", kubernetesType, namespace)
		return err
	}

	// extract NAME values from kubectl output
	fieldValues, err := kubectl.FindFieldValuesInKubectlOutput(kubectlCmdOutput, constants.KubectlFieldName)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("Unable to find [%s] in namespace [%s]...", kubernetesType, namespace), err.Error())
		log.Errorf("[deleteFromKubernetes] Unable to find [%s] in namespace [%s]... %s\n", kubernetesType, namespace, err.Error())
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
				loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to uninstall nginx-ingress-controller [%s] from namespace [%s]", kubernetesType, namespace), err.Error())
				log.Errorf("[deleteFromKubernetes] Unable to uninstall nginx-ingress-controller [%s] from namespace [%s]", kubernetesType, namespace)
				return err
			}
		}
	}
	return nil
}

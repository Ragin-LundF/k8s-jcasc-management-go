package uninstallactions

import (
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// ActionCleanupK8sNginxIngressController is a delegate method, that tries to cleanup everything.
// It ignores possible errors!
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

// ActionCleanupNginxIngressCtrlRoles will uninstall nginx-ingress roles
func ActionCleanupNginxIngressCtrlRoles(namespace string) {
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress Roles for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "role", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress Roles for namespace [%s]...done", namespace))
}

// ActionCleanupNginxIngressCtrlRoleBindings will uninstall nginx-ingress role bindings
func ActionCleanupNginxIngressCtrlRoleBindings(namespace string) {
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress RoleBindings for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "rolebindings", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress RoleBindings for namespace [%s]...done", namespace))
}

// ActionCleanupNginxIngressCtrlServiceAccounts will uninstall nginx-ingress service accounts
func ActionCleanupNginxIngressCtrlServiceAccounts(namespace string) {
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ServiceAccounts for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "sa", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ServiceAccounts for namespace [%s]...done", namespace))
}

// ActionCleanupNginxIngressCtrlClusterRoles will uninstall nginx-ingress clusterroles
func ActionCleanupNginxIngressCtrlClusterRoles(namespace string) {
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ClusterRoles for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName + "-clusterrole-" + namespace
	_ = deleteFromKubernetes(namespace, "clusterrole", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ClusterRoles for namespace [%s]...donen", namespace))
}

// ActionCleanupNginxIngressCtrlClusterRoleBinding will uninstall nginx-ingress clusterrole binding
func ActionCleanupNginxIngressCtrlClusterRoleBinding(namespace string) {
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ClusterRoleBinding for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName + "-clusterrole-nisa-binding-" + namespace
	_ = deleteFromKubernetes(namespace, "clusterrolebinding", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ClusterRoleBinding for namespace [%s]...done", namespace))
}

// ActionCleanupNginxIngressCtrlIngress will uninstall nginx-ingress ingress
func ActionCleanupNginxIngressCtrlIngress(namespace string) {
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ingress routes for namespace [%s]...", namespace))

	searchValue := models.GetConfiguration().Nginx.Ingress.Controller.DeploymentName
	_ = deleteFromKubernetes(namespace, "ingress", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Start to cleanup nginx-ingress ingress routes for namespace [%s]...done", namespace))
}

// generic function to delete role, rolebinding, sa... from kubectl
func deleteFromKubernetes(namespace string, kubernetesType string, filterValue string) (err error) {
	// Search for roles with deployment name
	kubectlCmdArgs := []string{
		kubernetesType,
		"-n", namespace,
	}
	kubectlCmdOutput, err := kubectl.ExecutorKubectl("get", kubectlCmdArgs)
	if err != nil {
		loggingstate.AddErrorEntry(fmt.Sprintf("  -> Unable to get [%s] for namespace [%s]", kubernetesType, namespace))
		return err
	}

	// extract NAME values from kubectl output
	fieldValues, err := kubectl.FindFieldValuesInKubectlOutput(kubectlCmdOutput, constants.KubectlFieldName)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("Unable to find [%s] in namespace [%s]...", kubernetesType, namespace), err.Error())
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
				return err
			}
		}
	}
	return nil
}

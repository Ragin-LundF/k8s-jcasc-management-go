package install

import (
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

// ActionCleanupK8sNginxIngressController is a delegate method, that tries to cleanup everything.
// It ignores possible errors!
// They will be logged, but it has no impact to the workflow
func (projectConfig *ProjectConfig) ActionCleanupK8sNginxIngressController() {
	// Nginx Ingress Ctrl Roles
	projectConfig.actionCleanupNginxIngressCtrlRoles()
	// Nginx Ingress Ctrl RoleBindings
	projectConfig.actionCleanupNginxIngressCtrlRoleBindings()
	// Nginx Ingress Ctrl ServiceAccounts
	projectConfig.actionCleanupNginxIngressCtrlServiceAccounts()
	// Nginx Ingress Ctrl ClusterRoles
	projectConfig.actionCleanupNginxIngressCtrlClusterRoles()
	// Nginx Ingress Ctrl ClusterRoleBindings
	projectConfig.actionCleanupNginxIngressCtrlClusterRoleBinding()
	// Nginx Ingress Ctrl ingress routes
	projectConfig.actionCleanupNginxIngressCtrlIngress()
}

// actionCleanupNginxIngressCtrlRoles will uninstall nginx-ingress roles
func (projectConfig *ProjectConfig) actionCleanupNginxIngressCtrlRoles() {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress Roles for namespace [%s]...",
		projectConfig.Project.Base.Namespace))

	var searchValue = configuration.GetConfiguration().Nginx.Ingress.Deployment.DeploymentName
	_ = projectConfig.deleteFromKubernetes("role", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress Roles for namespace [%s]...done",
		projectConfig.Project.Base.Namespace))
}

// actionCleanupNginxIngressCtrlRoleBindings will uninstall nginx-ingress role bindings
func (projectConfig *ProjectConfig) actionCleanupNginxIngressCtrlRoleBindings() {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress RoleBindings for namespace [%s]...",
		projectConfig.Project.Base.Namespace))

	var searchValue = configuration.GetConfiguration().Nginx.Ingress.Deployment.DeploymentName
	_ = projectConfig.deleteFromKubernetes("rolebindings", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress RoleBindings for namespace [%s]...done",
		projectConfig.Project.Base.Namespace))
}

// actionCleanupNginxIngressCtrlServiceAccounts will uninstall nginx-ingress service accounts
func (projectConfig *ProjectConfig) actionCleanupNginxIngressCtrlServiceAccounts() {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress ServiceAccounts for namespace [%s]...",
		projectConfig.Project.Base.Namespace))

	var searchValue = configuration.GetConfiguration().Nginx.Ingress.Deployment.DeploymentName
	_ = projectConfig.deleteFromKubernetes("sa", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress ServiceAccounts for namespace [%s]...done",
		projectConfig.Project.Base.Namespace))
}

// actionCleanupNginxIngressCtrlClusterRoles will uninstall nginx-ingress clusterroles
func (projectConfig *ProjectConfig) actionCleanupNginxIngressCtrlClusterRoles() {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress ClusterRoles for namespace [%s]...",
		projectConfig.Project.Base.Namespace))

	var searchValue = configuration.GetConfiguration().Nginx.Ingress.Deployment.DeploymentName + "-clusterrole-" + projectConfig.Project.Base.Namespace
	_ = projectConfig.deleteFromKubernetes("clusterrole", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress ClusterRoles for namespace [%s]...donen",
		projectConfig.Project.Base.Namespace))
}

// actionCleanupNginxIngressCtrlClusterRoleBinding will uninstall nginx-ingress clusterrole binding
func (projectConfig *ProjectConfig) actionCleanupNginxIngressCtrlClusterRoleBinding() {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress ClusterRoleBinding for namespace [%s]...",
		projectConfig.Project.Base.Namespace))

	var searchValue = configuration.GetConfiguration().Nginx.Ingress.Deployment.DeploymentName + "-clusterrole-nisa-binding-" + projectConfig.Project.Base.Namespace
	_ = projectConfig.deleteFromKubernetes("clusterrolebinding", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress ClusterRoleBinding for namespace [%s]...done",
		projectConfig.Project.Base.Namespace))
}

// actionCleanupNginxIngressCtrlIngress will uninstall nginx-ingress ingress
func (projectConfig *ProjectConfig) actionCleanupNginxIngressCtrlIngress() {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress ingress routes for namespace [%s]...",
		projectConfig.Project.Base.Namespace))

	var searchValue = configuration.GetConfiguration().Nginx.Ingress.Deployment.DeploymentName
	_ = projectConfig.deleteFromKubernetes("ingress", searchValue)

	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Start to cleanup nginx-ingress ingress routes for namespace [%s]...done",
		projectConfig.Project.Base.Namespace))
}

// generic function to delete role, rolebinding, sa... from kubectl
func (projectConfig *ProjectConfig) deleteFromKubernetes(kubernetesType string, filterValue string) (err error) {
	// Search for roles with deployment name
	var kubectlCmdArgs = []string{
		kubernetesType,
		"-n", projectConfig.Project.Base.Namespace,
	}
	kubectlCmdOutput, err := kubectl.ExecutorKubectl("get", kubectlCmdArgs)
	if err != nil {
		loggingstate.AddErrorEntry(fmt.Sprintf(
			"  -> Unable to get [%s] for namespace [%s]",
			kubernetesType,
			projectConfig.Project.Base.Namespace))
		return err
	}

	// extract NAME values from kubectl output
	fieldValues, err := kubectl.FindFieldValuesInKubectlOutput(kubectlCmdOutput, constants.KubectlFieldName)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
			"Unable to find [%s] in namespace [%s]...",
			kubernetesType,
			projectConfig.Project.Base.Namespace), err.Error())
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
			var kubectlUninstallCmdArgs = []string{
				"-n", projectConfig.Project.Base.Namespace,
				kubernetesType,
			}

			kubectlUninstallCmdArgs = append(kubectlUninstallCmdArgs, fieldValuesToDelete...)

			// Execute delete command
			_, err := kubectl.ExecutorKubectl("delete", kubectlUninstallCmdArgs)
			if err != nil {
				loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
					"  -> Unable to uninstall nginx-ingress-controller [%s] from namespace [%s]",
					kubernetesType,
					projectConfig.Project.Base.Namespace), err.Error())
				return err
			}
		}
	}
	return nil
}

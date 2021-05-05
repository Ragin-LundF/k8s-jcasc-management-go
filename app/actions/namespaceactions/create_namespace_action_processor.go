package namespaceactions

import (
	"fmt"
	"k8s-management-go/app/actions/install"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/loggingstate"
)

// ProcessNamespaceCreation processes the namespace creation
func ProcessNamespaceCreation(projectConfig install.ProjectConfig) (err error) {
	loggingstate.AddInfoEntry("Start creating namespace...")

	// check if namespace is existing
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Check if namespace [%s] is existing...",
		projectConfig.Project.Base.Namespace))
	nsIsAvailable, err := isNamespaceAvailable(projectConfig.Project.Base.Namespace)
	if err != nil {
		// it is ok, that the namespace is not available
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
			"  -> Check if namespace [%s] is existing...namespace not found with error.",
			projectConfig.Project.Base.Namespace), err.Error())
	}
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Check if namespace [%s] is existing...done",
		projectConfig.Project.Base.Namespace))

	// namespace is not available
	if !nsIsAvailable {
		// namespace does not exist, so create one
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"  -> Namespace [%s] is not available. Trying to create...",
			projectConfig.Project.Base.Namespace))

		kubectlCommandArgs := []string{
			"namespace", projectConfig.Project.Base.Namespace,
		}
		_, err = kubectl.ExecutorKubectl("create", kubectlCommandArgs)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
				"  -> Cannot create namespace [%s]",
				projectConfig.Project.Base.Namespace), err.Error())
			return err
		}

		loggingstate.AddInfoEntry(fmt.Sprintf(
			"  -> Namespace [%s] is not available. Trying to create...done",
			projectConfig.Project.Base.Namespace))
	}
	loggingstate.AddInfoEntry(fmt.Sprintf(
		"  -> Namespace [%s] found.",
		projectConfig.Project.Base.Namespace))
	loggingstate.AddInfoEntry("Start creating namespace...done")

	return nil
}

// check if namespace is available
func isNamespaceAvailable(namespace string) (namespaceIsAvailable bool, err error) {
	var kubectlCmdArgs = []string{
		"namespaces",
	}
	kubectlCmdOutput, err := kubectl.ExecutorKubectl("get", kubectlCmdArgs)
	if err != nil {
		return false, err
	}

	// check if output contains the namespace
	if kubectlCmdOutput != "" {
		namespaceIsAvailable = kubectl.CheckIfKubectlOutputContainsValueForField(kubectlCmdOutput, constants.KubectlFieldName, namespace)
	} else {
		namespaceIsAvailable = false
	}

	return namespaceIsAvailable, err
}

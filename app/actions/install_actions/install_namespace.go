package install_actions

import (
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

// check if namespace is available and create a new one if it does not exist
func CheckAndCreateNamespace(namespace string) (err error) {
	log := logger.Log()
	// check if namespace is existing
	log.Infof("[Install Namespace] Check if namespace [%s] is existing...", namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Check if namespace [%s] is existing...", namespace))
	nsIsAvailable, err := isNamespaceAvailable(namespace)
	if err != nil {
		// it is ok, that the namespace is not available
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Check if namespace [%s] is existing...namespace not found with error.", namespace), err.Error())
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Check if namespace [%s] is existing...done", namespace))

	// namespace is not available
	if !nsIsAvailable {
		// namespace does not exist, so create one
		loggingstate.AddInfoEntry(fmt.Sprintf("  -> Namespace [%s] is not available. Trying to create...", namespace))
		log.Infof("[Install Namespace] Namespace [%s] not found. Trying to create it...", namespace)

		kubectlCommandArgs := []string{
			"namespace", namespace,
		}
		_, err := kubectl.ExecutorKubectl("create", kubectlCommandArgs)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Cannot create namespace [%s]", namespace), err.Error())
			log.Errorf("[Install Namespace] Cannot create namespace [%s]", namespace)
			return err
		}

		loggingstate.AddInfoEntry(fmt.Sprintf("  -> Namespace [%s] is not available. Trying to create...done", namespace))
		log.Infof("[Install Namespace] Finished creating namespace [%s]...", namespace)
	} else {
		loggingstate.AddInfoEntry(fmt.Sprintf("  -> Namespace [%s] found.", namespace))
		log.Infof("[Install Namespace] Namespace [%s] found.", namespace)
	}
	return nil
}

// check if namespace is available
func isNamespaceAvailable(namespace string) (namespaceIsAvailable bool, err error) {
	kubectlCmdArgs := []string{
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

package namespace_actions

import (
	"fmt"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
)

func ProcessNamespaceCreation(state models.StateData) (err error) {
	log := logger.Log()
	loggingstate.AddInfoEntry("Start creating namespace...")
	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...")

	// check if namespace is existing
	log.Infof("[Install Namespace] Check if namespace [%s] is existing...", state.Namespace)
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Check if namespace [%s] is existing...", state.Namespace))
	nsIsAvailable, err := isNamespaceAvailable(state.Namespace)
	if err != nil {
		// it is ok, that the namespace is not available
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Check if namespace [%s] is existing...namespace not found with error.", state.Namespace), err.Error())
	}
	loggingstate.AddInfoEntry(fmt.Sprintf("  -> Check if namespace [%s] is existing...done", state.Namespace))

	// namespace is not available
	if !nsIsAvailable {
		// namespace does not exist, so create one
		loggingstate.AddInfoEntry(fmt.Sprintf("  -> Namespace [%s] is not available. Trying to create...", state.Namespace))
		log.Infof("[Install Namespace] Namespace [%s] not found. Trying to create it...", state.Namespace)

		kubectlCommandArgs := []string{
			"namespace", state.Namespace,
		}
		_, err := kubectl.ExecutorKubectl("create", kubectlCommandArgs)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Cannot create namespace [%s]", state.Namespace), err.Error())
			log.Errorf("[Install Namespace] Cannot create namespace [%s]", state.Namespace)
			return err
		}

		loggingstate.AddInfoEntry(fmt.Sprintf("  -> Namespace [%s] is not available. Trying to create...done", state.Namespace))
		log.Infof("[Install Namespace] Finished creating namespace [%s]...", state.Namespace)
	} else {
		loggingstate.AddInfoEntry(fmt.Sprintf("  -> Namespace [%s] found.", state.Namespace))
		log.Infof("[Install Namespace] Namespace [%s] found.", state.Namespace)
	}

	loggingstate.AddInfoEntry("-> Check and create namespace if necessary...done")
	loggingstate.AddInfoEntry("Start creating namespace...done")
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

package install

import (
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/logger"
)

// check if namespace is available and create a new one if it does not exist
func CheckAndCreateNamespace(namespace string) (err error) {
	log := logger.Log()
	// check if namespace is existing
	log.Infof("[Install Namespace] Check if namespace [%s] is existing...", namespace)
	loggingstate.AddInfoEntry("  -> Check if namespace [" + namespace + "] is existing...")
	nsIsAvailable, err := isNamespaceAvailable(namespace)
	if err != nil {
		// it is ok, that the namespace is not available
		loggingstate.AddErrorEntryAndDetails("  -> Check if namespace ["+namespace+"] is existing...namespace not found with error.", err.Error())
	}
	loggingstate.AddInfoEntry("  -> Check if namespace [" + namespace + "] is existing...done")

	// namespace is not available
	if !nsIsAvailable {
		// namespace does not exist, so create one
		loggingstate.AddInfoEntry("  -> Namespace [" + namespace + "] is not available. Trying to create...")
		log.Infof("[Install Namespace] Namespace [%s] not found. Trying to create it...", namespace)

		kubectlCommandArgs := []string{
			"namespace", namespace,
		}
		_, err := kubectl.ExecutorKubectl("create", kubectlCommandArgs)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("  -> Cannot create namespace ["+namespace+"]", err.Error())
			log.Errorf("[Install Namespace] Cannot create namespace [%s]", namespace)
			return err
		}

		loggingstate.AddInfoEntry("  -> Namespace [" + namespace + "] is not available. Trying to create...done")
		log.Infof("[Install Namespace] Finished creating namespace [%s]...", namespace)
	} else {
		loggingstate.AddInfoEntry("  -> Namespace [" + namespace + "] found.")
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

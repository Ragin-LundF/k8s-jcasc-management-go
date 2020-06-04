package install

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/logger"
)

// check if namespace is available and create a new one if it does not exist
func CheckAndCreateNamespace(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[Install Namespace] Check if namespace [" + namespace + "] is existing...")
	infoLog, err, nsIsAvailable := isNamespaceAvailable(namespace)
	info = info + infoLog
	if err != nil {
		log.Error(err)
		// it is ok, that the namespace is not available
		return info, nil
	}

	// namespace is not available
	if !nsIsAvailable {
		// namespace does not exist, so create one
		log.Info("[Install Namespace] Namespace [" + namespace + "] not found. Trying to create it...")
		info = info + constants.NewLine + "Namespace [" + namespace + "] does not exist! Trying to create it..."

		kubectlCommandArgs := []string{
			"namespace", namespace,
		}
		_, infoLog, err := kubectl.ExecutorKubectl("create", kubectlCommandArgs)
		info = info + constants.NewLine + infoLog
		if err != nil {
			log.Error("[Install Namespace] Cannot create namespace [" + namespace + "]")
			info = "Namespace creation failed:" + constants.NewLine + info
			return info, err
		}
		log.Info("[Install Namespace] Finished creating namespace [" + namespace + "]...")
	} else {
		log.Info("[Install Namespace] Namespace [" + namespace + "] found.")
		info = info + constants.NewLine + "Namespace [" + namespace + "] found."
	}
	return info, nil
}

// check if namespace is available
func isNamespaceAvailable(namespace string) (info string, err error, namespaceIsAvailable bool) {
	kubectlCmdArgs := []string{
		"namespaces",
	}
	kubectlCmdOutput, infoLog, err := kubectl.ExecutorKubectl("get", kubectlCmdArgs)
	info = info + constants.NewLine + infoLog
	if err != nil {
		return info, err, false
	}

	// check if output contains the namespace
	if kubectlCmdOutput != "" {
		namespaceIsAvailable = kubectl.CheckIfKubectlOutputContainsValueForField(kubectlCmdOutput, constants.KubectlFieldName, namespace)
	} else {
		namespaceIsAvailable = false
	}

	return info, err, namespaceIsAvailable
}

package install

import (
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/kubectl"
	"os/exec"
)

// check if namespace is available and create a new one if it does not exist
func CheckAndCreateNamespace(namespace string) (info string, err error) {
	infoLog, err, nsIsAvailable := isNamespaceAvailable(namespace)
	info = info + infoLog
	if err != nil {
		return info, err
	}

	// namespace is not available
	if !nsIsAvailable {
		// namespace does not exist, so create one
		info = info + constants.NewLine + "Namespace [" + namespace + "] does not exist! Try to create it..."
		outputNsCreate, err := exec.Command("kubectl", "create", "namespace", namespace).Output()
		if err != nil {
			info = info + constants.NewLine + "Namespace creation failed."
			return info, err
		}
		// return kubectl output
		info = info + constants.NewLine + "Kubectl Namespace creation output:"
		info = info + constants.NewLine + "==============="
		info = info + string(outputNsCreate)
		info = info + constants.NewLine + "==============="
	} else {
		info = info + constants.NewLine + "Namespace [" + namespace + "] found."
	}
	return info, err
}

// check if namespace is available
func isNamespaceAvailable(namespace string) (info string, err error, namespaceIsAvailable bool) {
	outputCmd, err := exec.Command("kubectl", "get", "namespaces").Output()
	if err != nil {
		return info, err, false
	}

	// check if output contains the namespace
	if outputCmd != nil {
		namespaceIsAvailable = kubectl.CheckIfKubectlOutputContainsValueForField(string(outputCmd), constants.KubectlOutputFieldNamespace, namespace)
	} else {
		namespaceIsAvailable = false
	}

	return info, err, namespaceIsAvailable
}

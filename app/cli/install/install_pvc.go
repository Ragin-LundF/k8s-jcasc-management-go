package install

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/logger"
)

type PvcClaimValuesYaml struct {
	Kind       string
	ApiVersion string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string
	}
	Spec struct {
		AccessModes      []string
		StorageClassName string
		Resources        struct {
			Requests struct {
				Storage string
			}
		}
	}
}

// install PVC is needed
func PersistenceVolumeClaimInstall(namespace string) (info string, err error) {
	log := logger.Log()
	log.Info("[PVC Install] Check if PVC should be installed on namespace [" + namespace + "]")
	// prepare file directories
	projectDir := files.AppendPath(models.GetProjectBaseDirectory(), namespace)
	pvcClaimValuesFilePath := files.AppendPath(projectDir, constants.FilenamePvcClaim)

	// open file
	if files.FileOrDirectoryExists(pvcClaimValuesFilePath) {
		log.Info("[PVC Install] Kubernetes PVC specification found for namespace [" + namespace + "]...")
		// variable to check, if pvc already exists
		infoLog, err, pvcName := readPvcNameFromFile(pvcClaimValuesFilePath)
		info = info + constants.NewLine + infoLog
		if err != nil {
			return info, err
		}
		// if no name was found, something was wrong here...
		if pvcName == nil || *pvcName == "" {
			log.Error("[PVC Install] PVC specification was found for namespace [" + namespace + "], but no name was specified.")
			err = errors.New("PVC specification was found for namespace [" + namespace + "], but no name was specified.")
			return info, err
		}

		// check if pvc is already available in namespace
		log.Info("[PVC Install] Checking if PVC [" + *pvcName + "] is already available for namespace [" + namespace + "].")
		infoLog, err, pvcExists := isPvcAvailableInNamespace(namespace, *pvcName)
		info = info + infoLog

		// no PVC found, so install it
		if !pvcExists {
			log.Info("[PVC Install] PVC [" + *pvcName + "] does not exist in namespace [" + namespace + "]. Trying to install it...")
			info = info + constants.NewLine + "PVC [" + *pvcName + "] does not exist in namespace [" + namespace + "]. Trying to install it..."

			kubectlCmdArgs := []string{
				"-n", namespace,
				"-f", pvcClaimValuesFilePath,
			}
			_, infoLog, err := kubectl.ExecutorKubectl("apply", kubectlCmdArgs)
			info = info + constants.NewLine + infoLog
			if err != nil {
				log.Error("[PVC Install] Cannot create PVC [" + *pvcName + "] for namespace [" + namespace + "]")
				info = "Cannot create PVC [" + *pvcName + "] for namespace [" + namespace + "]" + constants.NewLine + info
				return info, err
			}
			log.Info("[PVC Install] Finished creating PVC [" + *pvcName + "] for namespace [" + namespace + "]...")
		} else {
			log.Info("[PVC Install] PVC [" + *pvcName + "] in namespace [" + namespace + "] found. No need to install it...")
			info = info + constants.NewLine + "PVC [" + *pvcName + "] in namespace [" + namespace + "] found. No need to install it..."
		}
	}

	return info, err
}

// read PVC specification and find name
func readPvcNameFromFile(pvcClaimValuesFilePath string) (info string, err error, pvcName *string) {
	// read PVC claim values.yaml file
	yamlFile, err := ioutil.ReadFile(pvcClaimValuesFilePath)
	if err != nil {
		return info, err, nil
	}

	// parse YAML
	var pvcClaimValues PvcClaimValuesYaml
	err = yaml.Unmarshal(yamlFile, &pvcClaimValues)
	if err != nil {
		return info, err, nil
	}
	pvcMetaName := pvcClaimValues.Metadata.Name

	return info, err, &pvcMetaName
}

// internal function to check if PVC is available in namespace
func isPvcAvailableInNamespace(namespace string, pvcName string) (info string, err error, pvcExists bool) {
	pvcExists = false
	// read all pvc from K8S
	kubectlCmdArgs := []string{
		"-n", namespace,
		"pvc",
	}
	kubectlCmdOutput, infoLog, err := kubectl.ExecutorKubectl("get", kubectlCmdArgs)
	info = info + constants.NewLine + infoLog
	if err != nil {
		return info, err, false
	}

	// check if output contains pvcName
	if kubectlCmdOutput != "" {
		pvcExists = kubectl.CheckIfKubectlOutputContainsValueForField(kubectlCmdOutput, constants.KubectlFieldName, pvcName)
	}
	return info, err, pvcExists
}

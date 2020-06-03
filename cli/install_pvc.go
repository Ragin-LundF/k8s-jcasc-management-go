package cli

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s-management-go/constants"
	"k8s-management-go/models/config"
	"k8s-management-go/utils/files"
	"k8s-management-go/utils/kubectl"
	"log"
	"os/exec"
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
func InstallPersistenceVolumeClaim(namespace string) (info string, err error) {
	// prepare file directories
	projectDir := files.AddFilePath(config.FilePathWithBasePath(config.GetConfiguration().Directories.ProjectsBaseDirectory), namespace)
	pvcClaimValuesFilePath := files.AddFilePath(projectDir, constants.FilenamePvcClaim)

	// open file
	if files.FileExists(pvcClaimValuesFilePath) {
		// variable to check, if pvc already exists
		infoLog, err, pvcName := readPvcNameFromFile(pvcClaimValuesFilePath)
		info = info + infoLog
		if err != nil {
			log.Println(err)
			return info, err
		}
		// if no name was found, something was wrong here...
		if pvcName == nil || *pvcName == "" {
			err = errors.New("PVC specification was found for namespace [" + namespace + "], but no name was specified.")
			return info, err
		}

		// check if pvc is already available in namespace
		infoLog, err, pvcExists := isPvcAvailableInNamespace(namespace, *pvcName)
		info = info + infoLog

		// no PVC found, so install it
		if !pvcExists {
			info = info + "\nPVC specification, but no PVC found in namespace...try to install it."
			outputInstallPvc, err := exec.Command("kubectl", "-n", namespace, "apply", "-f", pvcClaimValuesFilePath).Output()
			if err != nil {
				log.Println(err)
				return info, err
			}
			info = info + "\nKubectl output:"
			info = info + "\n==============="
			info = info + string(outputInstallPvc)
			info = info + "\n==============="
		} else {
			info = info + "\nFound namespace [" + *pvcName + "]...No need to install it."
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
	output, err := exec.Command("kubectl", "-n", namespace, "get", "pvc").Output()
	if err != nil {
		log.Println(err)
		return info, err, false
	}

	// check if output contains pvcName
	if output != nil {
		pvcExists = kubectl.CheckIfKubectlOutputContainsValueForField(string(output), constants.KubectlOutputFieldPvcName, pvcName)
	}
	return info, err, pvcExists
}

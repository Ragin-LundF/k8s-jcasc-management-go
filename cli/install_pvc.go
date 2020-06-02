package cli

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s-management-go/constants"
	"k8s-management-go/models/config"
	"k8s-management-go/utils/files"
	"log"
	"os"
	"os/exec"
	"strings"
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

func InstallPersistenceVolumeClaim() (info string, err error) {
	// ask for namespace
	namespace, err := DialogAskForNamespace()
	if err != nil {
		log.Println(err)
		return info, err
	}

	// prepare file directories
	projectDir := files.AddFilePath(config.FilePathWithBasePath(config.GetConfiguration().Directories.ProjectsBaseDirectory), namespace)
	pvcClaimValuesFilePath := files.AddFilePath(projectDir, constants.FilenamePvcClaim)

	// open file
	if files.FileExists(pvcClaimValuesFilePath) {
		// read PVC claim values.yaml file
		yamlFile, err := ioutil.ReadFile(pvcClaimValuesFilePath)
		if err != nil {
			return info, err
		}

		// parse YAML
		var pvcClaimValues PvcClaimValuesYaml
		err = yaml.Unmarshal(yamlFile, &pvcClaimValues)
		if err != nil {
			return info, err
		}
		pvcName := pvcClaimValues.Metadata.Name

		// read all pvc from K8S
		cmd := exec.Command("kubectl", "-n", namespace, "get", "pvc")
		if err := cmd.Run(); err != nil {
			log.Println(err)
			return info, err
		}

		// Get output of command
		output, err := cmd.Output()
		if err != nil {
			log.Println(err)
			return info, err
		}

		// check if output contains pvcName
		if output != nil {
			// TODO: Parse correctly (awk print $1 with trim)
			if !strings.Contains(string(output), pvcName) {
				// if not found, we install it
			}
		}
	}

	return info, err
}

package install

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/loggingstate"
)

// ActionPersistenceVolumeClaimInstall installs PVC if it is needed
func (projectConfig *ProjectConfig) ActionPersistenceVolumeClaimInstall() (err error) {
	loggingstate.AddInfoEntry(fmt.Sprintf(
		" -> Check if PVC should be installed on namespace [%s]",
		projectConfig.Project.Base.Namespace))

	// prepare file directories
	var pvcClaimValuesFilePath string
	pvcClaimValuesFilePath, err = projectConfig.PrepareInstallYAML(constants.FilenamePvcClaim)

	// open file
	if files.FileOrDirectoryExists(pvcClaimValuesFilePath) {
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"  -> Kubernetes PVC specification found for namespace [%s]...",
			projectConfig.Project.Base.Namespace))

		// variable to check, if pvc already exists
		err = projectConfig.readPvcNameFromFile(pvcClaimValuesFilePath)
		if err != nil {
			project.RemoveTempFile(pvcClaimValuesFilePath)
			return err
		}

		// if no name was found, something was wrong here...
		if len(projectConfig.Project.Base.ExistingVolumeClaim) == 0 {
			loggingstate.AddErrorEntry(fmt.Sprintf(
				"  -> PVC specification was found for namespace [%s], but no name was specified.",
				projectConfig.Project.Base.Namespace))
			err = fmt.Errorf(
				"[PVC Install] PVC specification was found for namespace [%s], but no name was specified. ",
				projectConfig.Project.Base.Namespace)
			project.RemoveTempFile(pvcClaimValuesFilePath)
			return err
		}

		// check if pvc is already available in namespace
		loggingstate.AddInfoEntry(fmt.Sprintf(
			"  -> Checking if PVC [%s] is already available for namespace [%s].",
			projectConfig.Project.Base.ExistingVolumeClaim,
			projectConfig.Project.Base.Namespace))
		var pvcExists, _ = projectConfig.isPvcAvailableInNamespace()

		// no PVC found, so install it
		if !pvcExists {
			loggingstate.AddInfoEntry(fmt.Sprintf(
				"  -> PVC [%s] does not exist in namespace [%s]. Trying to install it...",
				projectConfig.Project.Base.ExistingVolumeClaim,
				projectConfig.Project.Base.Namespace))

			// executing command
			var kubectlCmdArgs = []string{
				"-n", projectConfig.Project.Base.Namespace,
				"-f", pvcClaimValuesFilePath,
			}
			if _, err = kubectl.ExecutorKubectl("apply", kubectlCmdArgs); err != nil {
				loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
					"  -> Cannot create PVC [%s] for namespace [%s]",
					projectConfig.Project.Base.ExistingVolumeClaim,
					projectConfig.Project.Base.Namespace), err.Error())
				project.RemoveTempFile(pvcClaimValuesFilePath)
				return err
			}

			loggingstate.AddInfoEntry(fmt.Sprintf(
				"  -> PVC [%s] does not exist in namespace [%s]. Trying to install it...done",
				projectConfig.Project.Base.ExistingVolumeClaim,
				projectConfig.Project.Base.Namespace))
		} else {
			loggingstate.AddInfoEntry(fmt.Sprintf(
				"  -> PVC [%s] in namespace [%s] found. No need to install it...",
				projectConfig.Project.Base.ExistingVolumeClaim,
				projectConfig.Project.Base.Namespace))
		}
	}
	project.RemoveTempFile(pvcClaimValuesFilePath)

	return err
}

// read PVC specification and find name if not already in configuration
func (projectConfig *ProjectConfig) readPvcNameFromFile(pvcClaimValuesFilePath string) (err error) {
	if len(projectConfig.Project.Base.ExistingVolumeClaim) == 0 {
		// read PVC claim values.yaml file
		var yamlFile []byte
		yamlFile, err = ioutil.ReadFile(pvcClaimValuesFilePath)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
				"  -> Unable to read pvc file [%s]...",
				pvcClaimValuesFilePath), err.Error())
			return err
		}

		// parse YAML
		var pvcClaimValues PvcClaimValuesYaml
		err = yaml.Unmarshal(yamlFile, &pvcClaimValues)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf(
				"  -> Unable to unmarshal pvc yaml file [%s]...",
				pvcClaimValuesFilePath), err.Error())
			return err
		}
		projectConfig.Project.SetPersistentVolumeClaimExistingName(pvcClaimValues.Metadata.Name)

		return nil
	}
	return nil
}

// internal function to check if PVC is available in namespace
func (projectConfig *ProjectConfig) isPvcAvailableInNamespace() (pvcExists bool, err error) {
	pvcExists = false
	// read all pvc from K8S
	var kubectlCmdArgs = []string{
		"-n", projectConfig.Project.Base.Namespace,
		"pvc",
	}
	kubectlCmdOutput, err := kubectl.ExecutorKubectl("get", kubectlCmdArgs)
	if err != nil {
		return false, err
	}

	// check if output contains pvcName
	if kubectlCmdOutput != "" {
		pvcExists = kubectl.CheckIfKubectlOutputContainsValueForField(
			kubectlCmdOutput,
			constants.KubectlFieldName,
			projectConfig.Project.Base.ExistingVolumeClaim)
	}
	return pvcExists, err
}

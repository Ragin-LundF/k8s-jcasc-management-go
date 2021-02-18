package project

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/models"
	"testing"
)

const pvcName = "my-pvc"
const pvcNamespace = "my-Namespace"
const configJenkinsMasterDeploymentName = "jenkins-master"
const configJenkinsMasterPvcAccessMode = "ReadWriteOnce"
const configJenkinsMasterPvcStorageClassName = "hostpath"
const configJenkinsMasterPvcSize = "1Gi"

func TestCreatePersistentVolumeClaim(t *testing.T) {
	defaultPersistentVolumeClaimConfiguration()

	var pvc = NewPersistentVolumeClaim(pvcNamespace, pvcName)
	assert.Equal(t, pvcName, pvc.Metadata.Name)
	assert.Equal(t, pvcNamespace, pvc.Metadata.Namespace)

	assert.Equal(t, configJenkinsMasterDeploymentName, pvc.Metadata.Labels.Component)
	assert.Equal(t, configJenkinsMasterDeploymentName, pvc.Metadata.Labels.Name)

	assert.Equal(t, configJenkinsMasterPvcSize, pvc.Spec.Resources.StorageSize)
	assert.Equal(t, configJenkinsMasterPvcAccessMode, pvc.Spec.AccessMode)
	assert.Equal(t, configJenkinsMasterPvcStorageClassName, pvc.Spec.StorageClassName)
}

func TestCreatePersistentVolumeClaimWithOverwrittenConfig(t *testing.T) {
	const customLabelComponentName = "my-custom-Component"
	const customLabelName = "my-custom-label"
	const customMetadataName = "my-custom-Metadata-Name"
	const customNamespace = "my-custom-Namespace"

	defaultPersistentVolumeClaimConfiguration()

	var pvc = NewPersistentVolumeClaim(pvcNamespace, pvcName)
	pvc.SetMetadataLabelComponentName(customLabelComponentName)
	pvc.SetMetadataLabelName(customLabelName)
	pvc.SetMetadataName(customMetadataName)
	pvc.SetMetadataNamespace(customNamespace)

	assert.Equal(t, customMetadataName, pvc.Metadata.Name)
	assert.Equal(t, customNamespace, pvc.Metadata.Namespace)

	assert.Equal(t, customLabelComponentName, pvc.Metadata.Labels.Component)
	assert.Equal(t, customLabelName, pvc.Metadata.Labels.Name)

	assert.Equal(t, configJenkinsMasterPvcSize, pvc.Spec.Resources.StorageSize)
	assert.Equal(t, configJenkinsMasterPvcAccessMode, pvc.Spec.AccessMode)
	assert.Equal(t, configJenkinsMasterPvcStorageClassName, pvc.Spec.StorageClassName)
}

func defaultPersistentVolumeClaimConfiguration() {
	models.AssignToConfiguration("JENKINS_MASTER_DEPLOYMENT_NAME", configJenkinsMasterDeploymentName)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_ACCESS_MODE", configJenkinsMasterPvcAccessMode)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE", configJenkinsMasterPvcSize)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS", configJenkinsMasterPvcStorageClassName)
}

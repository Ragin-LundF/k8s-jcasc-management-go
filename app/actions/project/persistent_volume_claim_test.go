package project

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePersistentVolumeClaim(t *testing.T) {
	testDefaultProjectConfiguration(t, false)

	var pvc = NewPersistentVolumeClaim(testPvcNamespace)
	assert.Empty(t, pvc.Metadata.Name)
	assert.Equal(t, testPvcNamespace, pvc.Metadata.Namespace)

	assert.Equal(t, testConfigJenkinsMasterDeploymentName, pvc.Metadata.Labels.Component)
	assert.Equal(t, testConfigJenkinsMasterDeploymentName, pvc.Metadata.Labels.Name)

	assertDefaultPvcConfiguration(pvc, t)
}

func TestCreatePersistentVolumeClaimWithOverwrittenConfig(t *testing.T) {
	testDefaultProjectConfiguration(t, false)

	const customLabelComponentName = "my-custom-Component"
	const customLabelName = "my-custom-label"
	const customMetadataName = "my-custom-Metadata-Name"
	const customNamespace = "my-custom-Namespace"

	var pvc = NewPersistentVolumeClaim(testPvcNamespace)
	pvc.SetMetadataLabelComponentName(customLabelComponentName)
	pvc.SetMetadataLabelName(customLabelName)
	pvc.SetMetadataName(customMetadataName)
	pvc.SetMetadataNamespace(customNamespace)

	println(pvc.Metadata.Name)
	assert.Equal(t, customMetadataName, pvc.Metadata.Name)
	assert.Equal(t, customNamespace, pvc.Metadata.Namespace)

	assert.Equal(t, customLabelComponentName, pvc.Metadata.Labels.Component)
	assert.Equal(t, customLabelName, pvc.Metadata.Labels.Name)

	assertDefaultPvcConfiguration(pvc, t)
}

func assertDefaultPvcConfiguration(pvc *persistentVolumeClaim, t *testing.T) {
	assert.Equal(t, testConfigJenkinsMasterPvcSize, pvc.Spec.Resources.StorageSize)
	assert.Equal(t, testConfigJenkinsMasterPvcAccessMode, pvc.Spec.AccessMode)
	assert.Equal(t, testConfigJenkinsMasterPvcStorageClassName, pvc.Spec.StorageClassName)
}

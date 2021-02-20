package project

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePersistentVolumeClaim(t *testing.T) {
	testDefaultProjectConfiguration(t, false)

	var pvc = NewPersistentVolumeClaim()

	assert.Empty(t, pvc.Metadata.Name)
	assertDefaultPvcConfiguration(pvc, t)
}

func TestCreatePersistentVolumeClaimWithOverwrittenConfig(t *testing.T) {
	testDefaultProjectConfiguration(t, false)

	const customLabelComponentName = "my-custom-Component"
	const customLabelName = "my-custom-label"
	const customMetadataName = "my-custom-Metadata-Name"

	var pvc = NewPersistentVolumeClaim()
	pvc.SetMetadataName(customMetadataName)

	assert.Equal(t, customMetadataName, pvc.Metadata.Name)
	assertDefaultPvcConfiguration(pvc, t)
}

func assertDefaultPvcConfiguration(pvc *persistentVolumeClaim, t *testing.T) {
	assert.Equal(t, testConfigJenkinsMasterPvcSize, pvc.Spec.Resources.StorageSize)
	assert.Equal(t, testConfigJenkinsMasterPvcAccessMode, pvc.Spec.AccessMode)
	assert.Equal(t, testConfigJenkinsMasterPvcStorageClassName, pvc.Spec.StorageClassName)
}

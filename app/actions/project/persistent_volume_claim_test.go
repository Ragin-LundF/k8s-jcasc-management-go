package project

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePersistentVolumeClaim(t *testing.T) {
	testDefaultProjectConfiguration(t, false)

	var pvc = NewPersistentVolumeClaim()

	assert.Equal(t, testConfigJenkinsMasterPvcSize, pvc.Spec.Resources.StorageSize)
	assert.Equal(t, testConfigJenkinsMasterPvcAccessMode, pvc.Spec.AccessMode)
	assert.Equal(t, testConfigJenkinsMasterPvcStorageClassName, pvc.Spec.StorageClassName)
}

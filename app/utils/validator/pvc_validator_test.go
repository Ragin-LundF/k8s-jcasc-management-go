package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatePersistentVolumeClaim(t *testing.T) {
	var pvcName = "pvc-jenkins"

	err := ValidatePersistentVolumeClaim(pvcName)

	assert.NoError(t, err)
}

func TestValidatePersistentVolumeClaimError(t *testing.T) {
	var pvcName = "pvc-jenkins-longer-than-253-characters-invalid-string.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum."

	err := ValidatePersistentVolumeClaim(pvcName)

	assert.Error(t, err)
}

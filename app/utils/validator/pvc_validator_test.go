package validator

import (
	"testing"
)

func TestValidatePersistentVolumeClaim(t *testing.T) {
	var pvcName = "pvc-jenkins"

	err := ValidatePersistentVolumeClaim(pvcName)
	if err != nil {
		t.Error("Failed. Valid PVC was rejected.")
	} else {
		t.Log("Success. Valid PVC was accepted.")
	}
}

func TestValidatePersistentVolumeClaimError(t *testing.T) {
	var pvcName = "pvc-jenkins-longer-than-253-characters-invalid-string.Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum."

	err := ValidatePersistentVolumeClaim(pvcName)
	if err != nil {
		t.Log("Success. Invalid PVC was rejected.")
	} else {
		t.Error("Failed. Invalid PVC was accepted.")
	}
}

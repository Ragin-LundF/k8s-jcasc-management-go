package project

import (
	"k8s-management-go/app/models"
)

// ----- Structures
// persistentVolumeClaim : Model which describes the persistent volume claim (PVC)
type persistentVolumeClaim struct {
	Spec pvcSpec
}

// pvcSpec : PVC specification
type pvcSpec struct {
	AccessMode       string
	StorageClassName string
	Resources        pvcSpecResources
}

// pvcSpecResources : PVC Spec Resources
type pvcSpecResources struct {
	StorageSize string
}

// NewPersistentVolumeClaim : creates a new instance of PersistentVolumeClaim
func NewPersistentVolumeClaim() *persistentVolumeClaim {
	var pvc = &persistentVolumeClaim{
		Spec: newDefaultSpec(),
	}

	return pvc
}

// newDefaultSpec : create new default spec for PVC
func newDefaultSpec() pvcSpec {
	var configuration = models.GetConfiguration()
	return pvcSpec{
		AccessMode:       configuration.Jenkins.Helm.Master.Persistence.AccessMode,
		StorageClassName: configuration.Jenkins.Helm.Master.Persistence.StorageClass,
		Resources:        newDefaultSpecResources(),
	}
}

// newDefaultSpecResources : create new default spec resources for PVC
func newDefaultSpecResources() pvcSpecResources {
	var configuration = models.GetConfiguration()
	return pvcSpecResources{
		StorageSize: configuration.Jenkins.Helm.Master.Persistence.Size,
	}
}

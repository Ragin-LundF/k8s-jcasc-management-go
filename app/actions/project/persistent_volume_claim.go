package project

import (
	"k8s-management-go/app/configuration"
)

// ----- Structures
// persistentVolumeClaim : Model which describes the persistent volume claim (PVC)
type persistentVolumeClaim struct {
	Spec pvcSpec `yaml:"pvcSpec,omitempty"`
}

// pvcSpec : PVC specification
type pvcSpec struct {
	AccessMode       string            `yaml:"accessMode,omitempty"`
	StorageClassName string            `yaml:"storageClassName,omitempty"`
	Resources        *pvcSpecResources `yaml:"resources,omitempty"`
}

// pvcSpecResources : PVC Spec Resources
type pvcSpecResources struct {
	StorageSize string `yaml:"storageSize,omitempty"`
}

// newPersistentVolumeClaim : creates a new instance of PersistentVolumeClaim
func newPersistentVolumeClaim() *persistentVolumeClaim {
	var pvc = &persistentVolumeClaim{
		Spec: newDefaultSpec(),
	}

	return pvc
}

// newDefaultSpec : create new default spec for PVC
func newDefaultSpec() pvcSpec {
	return pvcSpec{
		AccessMode:       configuration.GetConfiguration().Jenkins.Persistence.AccessMode,
		StorageClassName: configuration.GetConfiguration().Jenkins.Persistence.StorageClass,
		Resources:        newDefaultSpecResources(),
	}
}

// newDefaultSpecResources : create new default spec resources for PVC
func newDefaultSpecResources() *pvcSpecResources {
	return &pvcSpecResources{
		StorageSize: configuration.GetConfiguration().Jenkins.Persistence.StorageSize,
	}
}

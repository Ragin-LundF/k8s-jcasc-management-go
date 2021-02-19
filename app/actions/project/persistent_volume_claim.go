package project

import (
	"k8s-management-go/app/models"
)

// ----- Structures
// persistentVolumeClaim : Model which describes the persistent volume claim (PVC)
type persistentVolumeClaim struct {
	Metadata pvcMetadata
	Spec     pvcSpec
}

// pvcMetadata : PVC Metadata
type pvcMetadata struct {
	Namespace string
	Name      string
	Labels    pvcMetadataLabels
}

// pvcMetadataLabels : PVC Metadata Labels
type pvcMetadataLabels struct {
	Name      string
	Component string
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
func NewPersistentVolumeClaim(namespace string) *persistentVolumeClaim {
	var pvc = &persistentVolumeClaim{
		Metadata: newDefaultMetadata(namespace),
		Spec:     newDefaultSpec(),
	}

	return pvc
}

// ----- Setter to manipulate the default object
// SetMetadataName : Set PVC Name to Metadata
func (pvc *persistentVolumeClaim) SetMetadataName(pvcName string) {
	pvc.Metadata.Name = pvcName
}

// SetMetadataNamespace : Set PVC Namespace to Metadata
func (pvc *persistentVolumeClaim) SetMetadataNamespace(pvcNamespace string) {
	pvc.Metadata.Namespace = pvcNamespace
}

// SetMetadataLabelComponentName : Set PVC Component Name to the Metadata Labels
func (pvc *persistentVolumeClaim) SetMetadataLabelComponentName(componentName string) {
	pvc.Metadata.Labels.Component = componentName
}

// SetMetadataLabelName : Set PVC label Name to Metadata Labels
func (pvc *persistentVolumeClaim) SetMetadataLabelName(labelName string) {
	pvc.Metadata.Labels.Name = labelName
}

// ----- internal methods
// newDefaultMetadata : create a new Metadata object
func newDefaultMetadata(namespace string) pvcMetadata {
	return pvcMetadata{
		Namespace: namespace,
		Name:      "",
		Labels:    newDefaultMetadataLabels(),
	}
}

// newDefaultMetadataLabels : create new default label Metadata with configuration
func newDefaultMetadataLabels() pvcMetadataLabels {
	var configuration = models.GetConfiguration()
	return pvcMetadataLabels{
		Name:      configuration.Jenkins.Helm.Master.DeploymentName,
		Component: configuration.Jenkins.Helm.Master.DeploymentName,
	}
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

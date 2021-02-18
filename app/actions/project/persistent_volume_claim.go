package project

import (
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
)

// ----- Structures
// PersistentVolumeClaim : Model which describes the persistent volume claim (PVC)
type PersistentVolumeClaim struct {
	Metadata PvcMetadata
	Spec     PvcSpec
}

// PvcMetadata : PVC Metadata
type PvcMetadata struct {
	Namespace string
	Name      string
	Labels    PvcMetadataLabels
}

// PvcMetadataLabels : PVC Metadata Labels
type PvcMetadataLabels struct {
	Name      string
	Component string
}

// PvcSpec : PVC specification
type PvcSpec struct {
	AccessMode       string
	StorageClassName string
	Resources        struct {
		StorageSize string
	}
}

// NewPersistentVolumeClaim : creates a new instance of PersistentVolumeClaim
func NewPersistentVolumeClaim(namespace string, pvcLabelName string) *PersistentVolumeClaim {
	var pvc = &PersistentVolumeClaim{
		Metadata: newDefaultMetadata(namespace, pvcLabelName),
		Spec:     newDefaultSpec(),
	}

	return pvc
}

// ProcessTemplates : Interface implementation to process templates with PVC placeholder
func (pvc *PersistentVolumeClaim) ProcessTemplates(filename string) (err error) {
	for placeholder, value := range pvc.Placeholder() {
		_, err := files.ReplacePlaceholderInTemplate(filename, placeholder, value)
		if err != nil {
			return err
		}
	}

	return nil
}

// Placeholder : returns a map with the placeholder names and its values for this type
func (pvc *PersistentVolumeClaim) Placeholder() map[string]string {
	return map[string]string{
		"PVC_METADATA_NAME":               pvc.Metadata.Name,
		"PVC_METADATA_NAMESPACE":          pvc.Metadata.Namespace,
		"PVC_METADATA_LABELS_NAME":        pvc.Metadata.Labels.Name,
		"PVC_METADATA_LABELS_COMPONENT":   pvc.Metadata.Labels.Component,
		"PVC_SPEC_ACCESS_MODE":            pvc.Spec.AccessMode,
		"PVC_SPEC_STORAGE_CLASS_NAME":     pvc.Spec.StorageClassName,
		"PVC_SPEC_RESOURCES_STORAGE_SIZE": pvc.Spec.Resources.StorageSize,
	}
}

// ----- Setter to manipulate the default object
// SetMetadataName : Set PVC Name to Metadata
func (pvc *PersistentVolumeClaim) SetMetadataName(pvcName string) {
	pvc.Metadata.Name = pvcName
}

// SetMetadataNamespace : Set PVC Namespace to Metadata
func (pvc *PersistentVolumeClaim) SetMetadataNamespace(pvcNamespace string) {
	pvc.Metadata.Namespace = pvcNamespace
}

// SetMetadataLabelComponentName : Set PVC Component Name to the Metadata Labels
func (pvc *PersistentVolumeClaim) SetMetadataLabelComponentName(componentName string) {
	pvc.Metadata.Labels.Component = componentName
}

// SetMetadataLabelName : Set PVC label Name to Metadata Labels
func (pvc *PersistentVolumeClaim) SetMetadataLabelName(labelName string) {
	pvc.Metadata.Labels.Name = labelName
}

// ----- internal methods
// newDefaultMetadata : create a new Metadata object
func newDefaultMetadata(namespace string, pvcName string) PvcMetadata {
	return PvcMetadata{
		Namespace: namespace,
		Name:      pvcName,
		Labels:    newDefaultMetadataLabels(),
	}
}

// newDefaultMetadataLabels : create new default label Metadata with configuration
func newDefaultMetadataLabels() PvcMetadataLabels {
	var configuration = models.GetConfiguration()
	return PvcMetadataLabels{
		Name:      configuration.Jenkins.Helm.Master.DeploymentName,
		Component: configuration.Jenkins.Helm.Master.DeploymentName,
	}
}

// newDefaultSpec :
func newDefaultSpec() PvcSpec {
	var configuration = models.GetConfiguration()
	return PvcSpec{
		AccessMode:       configuration.Jenkins.Helm.Master.Persistence.AccessMode,
		StorageClassName: configuration.Jenkins.Helm.Master.Persistence.StorageClass,
		Resources:        struct{ StorageSize string }{StorageSize: configuration.Jenkins.Helm.Master.Persistence.Size},
	}
}

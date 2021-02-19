# Persistent Volume Claims

Persistent Volume Claims (PVC) are used to store Jenkins data on a defined volume on the Kubernetes cluster. It is
possible to distinguish between persistent volumes and non-persistent volumes.

If the data is to be stored on a persistent volume, this volume must be created before deployment. When this volume is
created, it is possible to enter the name of the PVC in the tool.

The tool uses the `pvc_claim` template to create the required configuration in Kubernetes.
It is also possible to use the placeholders in other templates.

## Placeholder variables for Persistent Volume Claims

| Placeholder | Description | Source |
| --- | --- | --- |
| `{{ .PersistentVolumeClaim.Metadata.Name }}` | Placeholder for `metadata.name` for PVC | user input |
| `{{ .PersistentVolumeClaim.Metadata.Namespace }}` | Placeholder for `metadata.namespace` for PVC | user input |
| `{{ .PersistentVolumeClaim.Metadata.Labels.Name }}` | Placeholder for `metadata.labels.app.kubernetes.io/name` for PVC | configuration `JENKINS_MASTER_DEPLOYMENT_NAME` |
| `{{ .PersistentVolumeClaim.Metadata.Labels.Component }}` | Placeholder for `metadata.labels.app.kubernetes.io/component` for PVC | configuration `JENKINS_MASTER_DEPLOYMENT_NAME` |
| `{{ .PersistentVolumeClaim.Spec.AccessMode }}` | Placeholder for `metadata.spec.accessModes` for PVC | configuration `JENKINS_MASTER_PERSISTENCE_ACCESS_MODE` |
| `{{ .PersistentVolumeClaim.Spec.Resources.StorageSize }}` | Placeholder for `metadata.spec.resources.requests.storage` for PVC | configuration `JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE` |
| `{{ .PersistentVolumeClaim.Spec.StorageClassName }}` | Placeholder for `metadata.spec.storageClassName` for PVC | configuration `JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS` |



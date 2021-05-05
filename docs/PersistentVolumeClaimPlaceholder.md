# Persistent Volume Claims

Persistent Volume Claims (PVC) are used to store Jenkins data on a defined volume on the Kubernetes cluster. It is
possible to distinguish between persistent volumes and non-persistent volumes.

If the data is to be stored on a persistent volume, this volume must be created before deployment. When this volume is
created, it is possible to enter the name of the PVC in the tool.

The tool uses the `pvc_claim` template to create the required configuration in Kubernetes.
It is also possible to use the placeholders in other templates.

## Placeholder variables for Persistent Volume Claims

| Placeholder | Description | Source | old config |
| --- | --- | --- | --- |
| `{{ .PersistentVolumeClaim.Spec.AccessMode }}` | Placeholder for `metadata.spec.accessModes` for PVC | configuration `jenkins.persistence.accessMode` | `JENKINS_MASTER_PERSISTENCE_ACCESS_MODE` |
| `{{ .PersistentVolumeClaim.Spec.Resources.StorageSize }}` | Placeholder for `metadata.spec.resources.requests.storage` for PVC | configuration `jenkins.persistence.storageSize` | `JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE` |
| `{{ .PersistentVolumeClaim.Spec.StorageClassName }}` | Placeholder for `metadata.spec.storageClassName` for PVC | configuration `jenkins.persistence.storageClass` | `JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS` |

## More placeholder
| Description | Link |
| --- | --- |
| Common base placeholder | [TemplatePlaceholder.md](TemplatePlaceholder.md) |
| Jenkins configuration as Code (JCasC) `jcasc_config.yaml` placeholder | [JcasCHelmValuesPlaceholder.md](JcasCHelmValuesPlaceholder.md) |
| Jenkins deployment `jenkins_helm_values.yaml` placeholder | [JenkinsHelmValuesPlaceholder.md](JenkinsHelmValuesPlaceholder.md) |
| Nginx Ingress Controller `nginx_ingress_helm_values.yaml` placeholder | [NginxIngressControllerPlaceholder.md](NginxIngressControllerPlaceholder.md) |

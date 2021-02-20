# Jenkins Helm Values Placeholder

Jenkins Helm Values are part of the Jenkins Helm Chart.

The tool uses the `jenkins_helm_values` template to create the required configuration in Kubernetes.
It is also possible to use the placeholders in other templates.

## Placeholder variables for Jenkins Helm Values

| Placeholder | Description | Source |
| --- | --- | --- |
| `{{ .JenkinsHelmValues.Master.Image }}` | Placeholder for `master.image` for Jenkins Helm Values | configuration `JENKINS_MASTER_CONTAINER_IMAGE` |
| `{{ .JenkinsHelmValues.Master.Tag }}` | Placeholder for `master.tag` for Jenkins Helm Values | configuration `JENKINS_MASTER_CONTAINER_IMAGE_TAG` |
| `{{ .JenkinsHelmValues.Master.ImagePullPolicy }}` | Placeholder for `master.imagePullPolicy` for Jenkins Helm Values | configuration `JENKINS_MASTER_CONTAINER_PULL_POLICY` |
| `{{ .JenkinsHelmValues.Master.ImagePullSecretName }}` | Placeholder for `master.imagePullSecretName` for Jenkins Helm Values | configuration `JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME` |
| `{{ .JenkinsHelmValues.Master.CustomJenkinsLabels }}` | Placeholder for `master.customJenkinsLabels` for Jenkins Helm Values | configuration `JENKINS_MASTER_DEFAULT_LABEL` |
| `{{ .JenkinsHelmValues.Master.AdminPassword }}` | Placeholder for `master.adminPassword` for Jenkins Helm Values | configuration `JENKINS_MASTER_ADMIN_PASSWORD` |
| `{{ .JenkinsHelmValues.Master.SidecarsConfigAutoReloadFolder }}` | Placeholder for `master.sidecars.configAutoReload.folder` for Jenkins Helm Values. This entry will also be parsed with the project structure. This allows to use also every template in the URL (e.g. `{{ .Namespace }}`) | configuration `JENKINS_JCASC_CONFIGURATION_URL` |
| `{{ .JenkinsHelmValues.Master.AuthorizationStrategyDenyAnonymousReadAccess }}` | Placeholder for `master.authorizationStrategy` for Jenkins Helm Values. | configuration `JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS` |
| `{{ .JenkinsHelmValues.Persistence.ExistingClaim }}` | Placeholder for `persistence.existingClaim` for Jenkins Helm Values. | user input for existing PVC |
| `{{ .JenkinsHelmValues.Persistence.StorageClass }}` | Placeholder for `persistence.storageClass` for Jenkins Helm Values. | configuration `JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS` |
| `{{ .JenkinsHelmValues.Persistence.AccessMode }}` | Placeholder for `persistence.accessMode` for Jenkins Helm Values. | configuration `JENKINS_MASTER_PERSISTENCE_ACCESS_MODE` |
| `{{ .JenkinsHelmValues.Persistence.Size }}` | Placeholder for `persistence.size` for Jenkins Helm Values. | configuration `JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE` |


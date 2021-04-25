# Jenkins Helm Values Placeholder

Jenkins Helm Values are part of the Jenkins Helm Chart.

The tool uses the `jenkins_helm_values` template to create the required configuration in Kubernetes.
It is also possible to use the placeholders in other templates.

## Placeholder variables for Jenkins Helm Values

| Placeholder | Description | Source | old config |
| --- | --- | --- | --- |
| `{{ .JenkinsHelmValues.AdditionalNamespaces }}` | Placeholder for `k8smanagement.rbac.additionalNamespaces` for Jenkins Helm Values | n/a |
| `{{ .JenkinsHelmValues.Controller.Image }}` | Placeholder for `controller.image` for Jenkins Helm Values | configuration `jenkins.container.image` | `JENKINS_MASTER_CONTAINER_IMAGE` |
| `{{ .JenkinsHelmValues.Controller.Tag }}` | Placeholder for `controller.tag` for Jenkins Helm Values | configuration `jenkins.container.tag` | `JENKINS_MASTER_CONTAINER_IMAGE_TAG` |
| `{{ .JenkinsHelmValues.Controller.ImagePullPolicy }}` | Placeholder for `controller.imagePullPolicy` for Jenkins Helm Values | configuration `jenkins.container.pullPolicy` | `JENKINS_MASTER_CONTAINER_PULL_POLICY` |
| `{{ .JenkinsHelmValues.Controller.ImagePullSecretName }}` | Placeholder for `controller.imagePullSecretName` for Jenkins Helm Values | configuration `jenkins.container.pullSecret` | `JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME` |
| `{{ .JenkinsHelmValues.Controller.CustomJenkinsLabels }}` | Placeholder for `controller.customJenkinsLabels` for Jenkins Helm Values | configuration `jenkins.controller.customJenkinsLabel` | `JENKINS_MASTER_DEFAULT_LABEL` |
| `{{ .JenkinsHelmValues.Controller.AdminPassword }}` | Placeholder for `controller.adminPassword` for Jenkins Helm Values | configuration `jenkins.controller.passwords.adminUser` | `JENKINS_MASTER_ADMIN_PASSWORD` |
| `{{ .JenkinsHelmValues.Controller.SidecarsConfigAutoReloadFolder }}` | Placeholder for `controller.sidecars.configAutoReload.folder` for Jenkins Helm Values. This entry will also be parsed with the project structure. This allows to use also every template in the URL (e.g. `{{ .Base.Namespace }}`) | configuration `jenkins.jcasc.configurationUrl` | `JENKINS_JCASC_CONFIGURATION_URL` |
| `{{ .JenkinsHelmValues.Controller.AuthorizationStrategyAllowAnonymousRead }}` | Placeholder for `controller.authorizationStrategy` for Jenkins Helm Values. | configuration `jenkins.jcasc.authorizationStrategy.allowAnonymousRead` | `JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS` |
| `{{ .JenkinsHelmValues.Persistence.ExistingClaim }}` | Placeholder for `persistence.existingClaim` for Jenkins Helm Values. | user input for existing PVC | n/a |
| `{{ .JenkinsHelmValues.Persistence.StorageClass }}` | Placeholder for `persistence.storageClass` for Jenkins Helm Values. | configuration `jenkins.persistence.storageClass` | `JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS` |
| `{{ .JenkinsHelmValues.Persistence.AccessMode }}` | Placeholder for `persistence.accessMode` for Jenkins Helm Values. | configuration `jenkins.persistence.accessMode` | `JENKINS_MASTER_PERSISTENCE_ACCESS_MODE` |
| `{{ .JenkinsHelmValues.Persistence.Size }}` | Placeholder for `persistence.size` for Jenkins Helm Values. | configuration `jenkins.persistence.storageSize` | `JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE` |

## More placeholder
| Description | Link |
| --- | --- |
| Common base placeholder | [TemplatePlaceholder.md](TemplatePlaceholder.md) |
| Jenkins configuration as Code (JCasC) `jcasc_config.yaml` placeholder | [JcasCHelmValuesPlaceholder.md](JcasCHelmValuesPlaceholder.md) |
| Nginx Ingress Controller `nginx_ingress_helm_values.yaml` placeholder | [NginxIngressControllerPlaceholder.md](NginxIngressControllerPlaceholder.md) |
| Persistent Volume Claim `pvc_claim.yaml` placeholder | [PersistentVolumeClaimPlaceholder.md](PersistentVolumeClaimPlaceholder.md) |

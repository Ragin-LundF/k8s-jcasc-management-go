# Template Placeholder

This documentation describes all possible template placeholders that can be used in all template files.

| Placeholder | Description | Source |
| --- | --- | --- |
| `{{ .Base.DeploymentName }}` | Placeholder for the Jenkins deployment name | configuration `JENKINS_MASTER_DEPLOYMENT_NAME` |
| `{{ .Base.Domain }}` | Placeholder for the domain entered by the user | user input |
| `{{ .Base.ExistingVolumeClaim }}` | Placeholder for the existing peristent volume claim | user input |
| `{{ .Base.IPAddress }}` | Placeholder for the IP address | user input |
| `{{ .Base.JenkinsUriPrefix }}` | Placeholder for the Jenkins URI prefix (e.g. `/jenkins`) | configuration `JENKINS_MASTER_DEFAULT_URI_PREFIX` |
| `{{ .Base.JenkinsURL }}` | Placeholder for a Jenkins URL. | If `NGINX_LOADBALANCER_ANNOTATIONS_ENABLED` is enabled, the entered domain or `<namespace>`.<configuration `NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME`>. Otherwise the IP address will be returned. |
| `{{ .Base.Namespace }}` | Placeholder for the namespace | user input |

## More placeholder
| Description | Link |
| --- | --- |
| Jenkins configuration as Code (JCasC) `jcasc_config.yaml` placeholder | [JcasCHelmValuesPlaceholder.md](JcasCHelmValuesPlaceholder.md) |
| Jenkins deployment `jenkins_helm_values.yaml` placeholder | [JenkinsHelmValuesPlaceholder.md](JenkinsHelmValuesPlaceholder.md) |
| Nginx Ingress Controller `nginx_ingress_helm_values.yaml` placeholder | [NginxIngressControllerPlaceholder.md](NginxIngressControllerPlaceholder.md) |
| Persistent Volume Claim `pvc_claim.yaml` placeholder | [PersistentVolumeClaimPlaceholder.md](PersistentVolumeClaimPlaceholder.md) |

# Template Placeholder

This documentation describes all possible template placeholders that can be used in all template files.

| Placeholder | Description | Source | old config |
| --- | --- | --- | --- |
| `{{ .Base.DeploymentName }}` | Placeholder for the Jenkins deployment name | configuration `jenkins.controller.deploymentName` | `JENKINS_MASTER_DEPLOYMENT_NAME` |
| `{{ .Base.Domain }}` | Placeholder for the domain entered by the user | user input | n/a |
| `{{ .Base.ExistingVolumeClaim }}` | Placeholder for the existing peristent volume claim | user input | n/a |
| `{{ .Base.IPAddress }}` | Placeholder for the IP address | user input | n/a |
| `{{ .Base.JenkinsUriPrefix }}` | Placeholder for the Jenkins URI prefix (e.g. `/jenkins`) | configuration `jenkins.controller.defaultURIPrefix` | `JENKINS_MASTER_DEFAULT_URI_PREFIX` |
| `{{ .Base.JenkinsURL }}` | Placeholder for a Jenkins URL. | If `nginx.loadbalancer.annotations.enabled` is `true`, the entered domain or `<namespace>`.<configuration `nginx.loadbalancer.externalDNS.hostName`>. Otherwise the IP address will be returned. | n/a |
| `{{ .Base.Namespace }}` | Placeholder for the namespace | user input | n/a |

## More placeholder
| Description | Link |
| --- | --- |
| Jenkins configuration as Code (JCasC) `jcasc_config.yaml` placeholder | [JcasCHelmValuesPlaceholder.md](JcasCHelmValuesPlaceholder.md) |
| Jenkins deployment `jenkins_helm_values.yaml` placeholder | [JenkinsHelmValuesPlaceholder.md](JenkinsHelmValuesPlaceholder.md) |
| Nginx Ingress Controller `nginx_ingress_helm_values.yaml` placeholder | [NginxIngressControllerPlaceholder.md](NginxIngressControllerPlaceholder.md) |
| Persistent Volume Claim `pvc_claim.yaml` placeholder | [PersistentVolumeClaimPlaceholder.md](PersistentVolumeClaimPlaceholder.md) |

# Nginx Ingress Controller

The Nginx Ingress Controller can be used to define an ingress controller and load balancer for Jenkins in the namespace.

The tool uses the `nginx_ingress_helm_values` template to create the required configuration in Kubernetes.
It is also possible to use the placeholders in other templates.

## Placeholder variables for Nginx Ingress Controller

| Placeholder | Description | Source | old config |
| --- | --- | --- | --- |
| `{{ .Nginx.Ingress.AnnotationIngressClass }}` | Placeholder for `ingress.annotationIngressClass` for Nginx Ingress Controller | configuration `nginx.ingress.annotationclass` | `NGINX_INGRESS_ANNOTATION_CLASS` |
| `{{ .Nginx.Ingress.DeploymentName }}` | Placeholder for `ingress.deploymentName` for Nginx Ingress Controller | configuration `nginx.ingress.deployment.deploymentName` | `NGINX_INGRESS_DEPLOYMENT_NAME` |
| `{{ .Nginx.Ingress.ContainerImage }}` | Placeholder for `ingress.containerImage` for Nginx Ingress Controller | configuration `nginx.ingress.container.image` | `NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE` |
| `{{ .Nginx.Ingress.ImagePullSecrets }}` | Placeholder for `ingress.imagePullSecrets` for Nginx Ingress Controller | configuration `nginx.ingress.container.pullSecret` | `NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS` |
| `{{ .Nginx.Ingress.EnableControllerForNamespace }}` | Placeholder for `ingress.controllerForNamespace.enabled` for Nginx Ingress Controller | configuration `nginx.ingress.deployment.forEachNamespace` | `NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE` |
| `{{ .Nginx.LoadBalancer.Enabled }}` | Placeholder for `loadbalancer.enabled` for Nginx Ingress Controller load balancer | configuration `nginx.loadbalancer.enabled` | `NGINX_LOADBALANCER_ENABLED` |
| `{{ .Nginx.LoadBalancer.Ports.HTTP.Port }}` | Placeholder for `loadbalancer.ports.http.port` for Nginx Ingress Controller load balancer | configuration `nginx.loadbalancer.ports.http` | `NGINX_LOADBALANCER_HTTP_PORT` |
| `{{ .Nginx.LoadBalancer.Ports.HTTP.TargetPort }}` | Placeholder for `loadbalancer.ports.http.targetPort` for Nginx Ingress Controller load balancer | configuration `nginx.loadbalancer.ports.httpTarget` | `NGINX_LOADBALANCER_HTTP_TARGETPORT` |
| `{{ .Nginx.LoadBalancer.Ports.HTTPS.Port }}` | Placeholder for `loadbalancer.ports.https.port` for Nginx Ingress Controller load balancer | configuration `nginx.loadbalancer.ports.https` | `NGINX_LOADBALANCER_HTTPS_PORT` |
| `{{ .Nginx.LoadBalancer.Ports.HTTPS.TargetPort }}` | Placeholder for `loadbalancer.ports.https.targetPort` for Nginx Ingress Controller load balancer | configuration `nginx.loadbalancer.ports.httpsTarget` | `NGINX_LOADBALANCER_HTTPS_TARGETPORT` |
| `{{ .Nginx.LoadBalancer.Annotations.Enabled }}` | Placeholder for `loadbalancer.annotations.enabled` for Nginx Ingress Controller load balancer | configuration `nginx.loadbalancer.annotations.enabled` | `NGINX_LOADBALANCER_ANNOTATIONS_ENABLED` |
| `{{ .Nginx.LoadBalancer.Annotations.ExternalDnsHostname }}` | Placeholder for `loadbalancer.annotations.external_dns_hostname` for Nginx Ingress Controller load balancer | configuration `nginx.loadbalancer.externalDNS.hostName` | `NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME` |
| `{{ .Nginx.LoadBalancer.Annotations.ExternalDnsTtl }}` | Placeholder for `loadbalancer.annotations.external_dns_ttl` for Nginx Ingress Controller load balancer | configuration `nginx.loadbalancer.externalDNS.ttl` | `NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_TTL` |

## More placeholder
| Description | Link |
| --- | --- |
| Common base placeholder | [TemplatePlaceholder.md](TemplatePlaceholder.md) |
| Jenkins configuration as Code (JCasC) `jcasc_config.yaml` placeholder | [JcasCHelmValuesPlaceholder.md](JcasCHelmValuesPlaceholder.md) |
| Jenkins deployment `jenkins_helm_values.yaml` placeholder | [JenkinsHelmValuesPlaceholder.md](JenkinsHelmValuesPlaceholder.md) |
| Persistent Volume Claim `pvc_claim.yaml` placeholder | [PersistentVolumeClaimPlaceholder.md](PersistentVolumeClaimPlaceholder.md) |

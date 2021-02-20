# Nginx Ingress Controller

The Nginx Ingress Controller can be used to define an ingress controller and load balancer for Jenkins in the namespace.

The tool uses the `nginx_ingress_helm_values` template to create the required configuration in Kubernetes.
It is also possible to use the placeholders in other templates.

## Placeholder variables for Nginx Ingress Controller

| Placeholder | Description | Source |
| --- | --- | --- |
| `{{ .Nginx.Ingress.AnnotationIngressClass }}` | Placeholder for `ingress.annotationIngressClass` for Nginx Ingress Controller | configuration `NGINX_INGRESS_ANNOTATION_CLASS` |
| `{{ .Nginx.Ingress.LoadBalancerIP }}` | Placeholder for `ingress.loadbalancerIp` for Nginx Ingress Controller | user input |
| `{{ .Nginx.Ingress.DeploymentName }}` | Placeholder for `ingress.deploymentName` for Nginx Ingress Controller | configuration `NGINX_INGRESS_DEPLOYMENT_NAME` |
| `{{ .Nginx.Ingress.ContainerImage }}` | Placeholder for `ingress.containerImage` for Nginx Ingress Controller | configuration `NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE` |
| `{{ .Nginx.Ingress.ImagePullSecrets }}` | Placeholder for `ingress.imagePullSecrets` for Nginx Ingress Controller | configuration `NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS` |
| `{{ .Nginx.Ingress.EnableControllerForNamespace }}` | Placeholder for `ingress.controllerForNamespace.enabled` for Nginx Ingress Controller | configuration `NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE` |
| `{{ .Nginx.LoadBalancer.Enabled }}` | Placeholder for `loadbalancer.enabled` for Nginx Ingress Controller load balancer | configuration `NGINX_LOADBALANCER_ENABLED` |
| `{{ .Nginx.LoadBalancer.Ports.HTTP.Port }}` | Placeholder for `loadbalancer.ports.http.port` for Nginx Ingress Controller load balancer | configuration `NGINX_LOADBALANCER_HTTP_PORT` |
| `{{ .Nginx.LoadBalancer.Ports.HTTP.TargetPort }}` | Placeholder for `loadbalancer.ports.http.targetPort` for Nginx Ingress Controller load balancer | configuration `NGINX_LOADBALANCER_HTTP_TARGETPORT` |
| `{{ .Nginx.LoadBalancer.Ports.HTTPS.Port }}` | Placeholder for `loadbalancer.ports.https.port` for Nginx Ingress Controller load balancer | configuration `NGINX_LOADBALANCER_HTTPS_PORT` |
| `{{ .Nginx.LoadBalancer.Ports.HTTPS.TargetPort }}` | Placeholder for `loadbalancer.ports.https.targetPort` for Nginx Ingress Controller load balancer | configuration `NGINX_LOADBALANCER_HTTPS_TARGETPORT` |
| `{{ .Nginx.LoadBalancer.Annotations.Enabled }}` | Placeholder for `loadbalancer.annotations.enabled` for Nginx Ingress Controller load balancer | configuration `NGINX_LOADBALANCER_ANNOTATIONS_ENABLED` |
| `{{ .Nginx.LoadBalancer.Annotations.ExternalDnsHostname }}` | Placeholder for `loadbalancer.annotations.external_dns_hostname` for Nginx Ingress Controller load balancer | If no domain was set by the user, the default is the `<namespace>`.<configuration `NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME`>. Else it uses the entered domain. |
| `{{ .Nginx.LoadBalancer.Annotations.ExternalDnsTtl }}` | Placeholder for `loadbalancer.annotations.external_dns_ttl` for Nginx Ingress Controller load balancer | configuration `NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_TTL` |

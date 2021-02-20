# Template Placeholder

This documentation describes all possible template placeholders that can be used in all template files.

| Placeholder | Description | Source |
| --- | --- | --- |
| `{{ .Namespace }}` | Placeholder for the namespace | user input |
| `{{ .JenkinsSetup.DeploymentName }}` | Placeholder for the Jenkins deployment name | configuration `JENKINS_MASTER_DEPLOYMENT_NAME` |
| `{{ .JenkinsSetup.JenkinsUriPrefix }}` | Placeholder for the Jenkins URI prefix (e.g. `/jenkins`) | configuration `JENKINS_MASTER_DEFAULT_URI_PREFIX` |

| Placeholder | Description | Source |
| --- | --- | --- |
| ##NAMESPACE## | Placeholder for namespace | user input |
| ##PUBLIC_IP_ADDRESS## | Placeholder for public IP address | user input |
| ##JENKINS_URL## | Jenkins URL or IP | user input |
| ##PROJECT_DIRECTORY## | Placeholder for project directory | |
| ##K8S_MGMT_JENKINS_CLOUD_TEMPLATES## | Placeholder to add cloud templates | user input |
| ##K8S_MGMT_JENKINS_SYSTEM_MESSAGE## | Placeholder for Jenkins system message | user input |
| ##JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL## | Placeholder for URL of JobDSL seed Job script | configuration |
| ##PROJECT_JENKINS_JOB_DEFINITION_REPOSITORY## | Placeholder for the Jobs repository | user input |
| ##JENKINS_JCASC_CONFIGURATION_URL## | Placeholder for Jenkins Configuration-as-Code configuration URL | configuration |
| ##JENKINS_MASTER_DEPLOYMENT_NAME## | Placeholder for the deployment name | calculated |
| ##JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS## | Placeholder for anonymous read access true/false configuration | configuration |
| ##JENKINS_MASTER_DEFAULT_URI_PREFIX## | Placeholder for the Jenkins URI prefix (e.g. /jenkins) | configuration |
| ##JENKINS_MASTER_DEFAULT_LABEL## | Placeholder for label of the master instance of Jenkins to pin the seed job on this instance | configuration |
| ##JENKINS_MASTER_ADMIN_PASSWORD## | Placeholder for the Jenkins master admin password (unencrypted), if no acl script was installed | configuration |
| ##JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED## | Placeholder for the encrypted Jenkins master admin password (bcrypt) | configuration |
| ##JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED## | Placeholder for the encrypted Jenkins master user password (default project user) (bcrypt) | configuration |
| ##JENKINS_MASTER_CONTAINER_IMAGE## | Placeholder for the Jenkins Docker image | configuration |
| ##JENKINS_MASTER_CONTAINER_IMAGE_TAG## | Placeholder for the Jenkins Docker image tag | configuration |
| ##JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME## | Placeholder for the Jenkins Docker image pull secret name | configuration |
| ##JENKINS_MASTER_CONTAINER_PULL_POLICY## | Placeholder for the Jenkins Docker image pull policy | configuration |
| ##NGINX_INGRESS_DEPLOYMENT_NAME## | Placeholder for the Nginx Ingress Controller deployment name | calculated |
| ##NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE## | Placeholder for the Nginx Ingress Controller Docker image | configuration |
| ##NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS## | Placeholder for the Nginx Ingress Controller Docker image pull secret | configuration |
| ##NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE## | Placeholder for Nginx Ingress Controller namespace | configuration | configuration |
| ##NGINX_INGRESS_ANNOTATION_CLASS## | Placeholder for Nginx Ingress Controller Kubernetes annotation class | configuration |
| ##NGINX_LOADBALANCER_ENABLED## | Placeholder for Nginx Ingress Controller, if a loadbalancer was enabled | configuration |
| ##NGINX_LOADBALANCER_HTTP_PORT## | Placeholder for Nginx load balancer HTTP port | configuration |
| ##NGINX_LOADBALANCER_HTTP_TARGETPORT## | Placeholder for Nginx load balancer HTTP target port | configuration |
| ##NGINX_LOADBALANCER_HTTPS_PORT## | Placeholder for Nginx load balancer HTTPS port | configuration |
| ##NGINX_LOADBALANCER_HTTPS_TARGETPORT## | Placeholder for Nginx load balancer HTTPS target port | configuration |
| ##NGINX_LOADBALANCER_ANNOTATIONS_ENABLED## | Placeholder for Nginx Annotation support | configuration |
| ##NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_TTL## | Placeholder for external DNS TTL annotation | configuration |
| ##NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME## | Placeholder for external DNS hostname annotation | configuration |
| ##KUBERNETES_SERVER_CERTIFICATE## | Placeholder for Kubernetes server certificate | configuration |
| ##KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID## | Placeholder for Kubernetes Docker Registry credentials ID | configuration |
| ##MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID## | Placeholder for the Maven repository secrets credentials ID | configuration |
| ##NPM_REPOSITORY_SECRETS_CREDENTIALS_ID## | Placeholder for the NPM repository secrets credentials ID | configuration |
| ##VCS_REPOSITORY_SECRETS_CREDENTIALS_ID## | Placeholder for the Version Control System secrets credentials ID | configuration |
| ##JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE## | Placeholder for the size of the PVC if the Jenkins master instance | configuration |
| ##JENKINS_MASTER_PERSISTENCE_ACCESS_MODE## | Placeholder for the Jenkins master persistence access mode (ReadWriteMany, ReadWriteSingle...) | configuration |
| ##JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS## | Placeholder for the Jenkins master persistence storage class (local-path, nfs-client...) | configuration |
| ##JENKINS_MASTER_PERSISTENCE_EXISTING_CLAIM## | Placeholder for the Jenkins master existing PVC | user input |
| ##K8S_MGMT_PERSISTENCE_VOLUME_CLAIM_NAME## | Placeholder for the PVC claim in the FilenamePvcClaim file (alias for ##JENKINS_MASTER_PERSISTENCE_EXISTING_CLAIM##). | user input |


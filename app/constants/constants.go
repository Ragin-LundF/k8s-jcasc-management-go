package constants

// ---- common
const NewLine = "\n"

// ---- directory and file configuration
// config directory
const DirConfig = "config"

// directory for shell scripts inside a project
const DirProjectScripts = "scripts"

// directory for cloud-templates inside the templates directory
const DirProjectTemplateCloudTemplates = "cloud-templates"

// prefix for scripts for installation
const DirProjectScriptsInstallPrefix = "i_"

// prefix for scripts for deinstallation
const DirProjectScriptsUninstallPrefix = "d_"

// directory of the Jenkins Helm Charts
const DirHelmJenkinsMaster = "/charts/jenkins-master"

// directory of the Nginx Ingress Controller Helm Charts
const DirHelmNginxIngressCtrl = "/charts/nginx-ingress-controller"

// filename of the configuration file
const FilenameConfiguration = "k8s_jcasc_mgmt.cnf"

// filename of the custom configuration file for overwrites
const FilenameConfigurationCustom = "k8s_jcasc_custom.cnf"

// filename of the Jenkins Configuration-as-Code Helm Values file
const FilenameJenkinsConfigurationAsCode = "jcasc_config.yaml"

// filename if the Jenkins Helm Values file
const FilenameJenkinsHelmValues = "jenkins_helm_values.yaml"

// filename of the Nginx Ingress Controller Helm Values file
const FilenameNginxIngressControllerHelmValues = "nginx_ingress_helm_values.yaml"

// filename of the Kubernetes PVC Claim values file
const FilenamePvcClaim = "pvc_claim.yaml"

// base filename for secrets (without .gpg extension)
const FilenameSecrets = "secrets.sh"

// secrets filename extension (will be added after FilenameSecrets)
const SecretsFileEncodedEnding = ".gpg"

// filename extension of scripts
const ScriptsFileEnding = ".sh"

// ---- commands (for delegation of the actions)
// install command
const CommandInstall = "install"

// uninstall command
const CommandUninstall = "uninstall"

// upgrade command
const CommandUpgrade = "upgrade"

// encrypt secrets command
const CommandEncryptSecrets = "encryptSecrets"

// decrypt secrets command
const CommandDecryptSecrets = "decryptSecrets"

// apply secrets command
const CommandApplySecrets = "applySecrets"

// apply secrets to all namespaces command
const CommandApplySecretsToAll = "applySecretsToAll"

// create project command
const CommandCreateProject = "createProject"

// create deployment only project command
const CommandCreateDeploymentOnlyProject = "createDeploymentOnlyProject"

// create namespace command
const CommandCreateNamespace = "createNamespace"

// create Jenkins user password command
const CommandCreateJenkinsUserPassword = "createJenkinsUserPassword"

// quit command
const CommandQuit = "quit"

// ---- helm commands (for execution)
// Helm command for install
const HelmCommandInstall = "install"

// Helm command for upgrade
const HelmCommandUpgrade = "upgrade"

// Helm command for uninstall
const HelmCommandUninstall = "uninstall"

// common messages
const CommonJenkinsSystemMessage = "Welcome to Jenkins"

// error
const ErrorPromptFailed = "prompt failed"

// ---- colors
// Normal color
const ColorNormal = "\033[0m"

// Info color
const ColorInfo = "\033[1;34m"

// Error color
const ColorError = "\033[1;31m"

// kubectl field names
const KubectlFieldName = "NAME"

// ---- Template Strings
// Placeholder for namespace
const TemplateNamespace = "##NAMESPACE##"

// Placeholder for public IP address
const TemplatePublicIpAddress = "##PUBLIC_IP_ADDRESS##"

// Placeholder for project directory
const TemplateProjectDirectory = "##PROJECT_DIRECTORY##"

// Placeholder to add cloud templates
const TemplateJenkinsCloudTemplates = "##K8S_MGMT_JENKINS_CLOUD_TEMPLATES##"

// Placeholder for Jenkins system message
const TemplateJenkinsSystemMessage = "##K8S_MGMT_JENKINS_SYSTEM_MESSAGE##"

// Placeholder for URL of JobDSL seed Job script
const TemplateJenkinsJobDslSeedJobScriptUrl = "##JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL##"

// Placeholder for the Jobs repository
const TemplateJobDefinitionRepository = "##PROJECT_JENKINS_JOB_DEFINITION_REPOSITORY##"

// Placeholder for Jenkins Configuration-as-Code configuration URL
const TemplateJenkinsJcascConfigurationUrl = "##JENKINS_JCASC_CONFIGURATION_URL##"

// Placeholder for the deployment name
const TemplateJenkinsMasterDeploymentName = "##JENKINS_MASTER_DEPLOYMENT_NAME##"

// Placeholder for anonymous read access true/false configuration
const TemplateJenkinsMasterDenyAnonymousReadAccess = "##JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS##"

// Placeholder for the Jenkins URI prefix (e.g. /jenkins)
const TemplateJenkinsMasterDefaultUriPrefix = "##JENKINS_MASTER_DEFAULT_URI_PREFIX##"

// Placeholder for label of the master instance of Jenkins to pin the seed job on this instance
const TemplateJenkinsMasterDefaultLabel = "##JENKINS_MASTER_DEFAULT_LABEL##"

// Placeholder for the Jenkins master admin password (unencrypted), if no acl script was installed
const TemplateJenkinsMasterAdminPassword = "##JENKINS_MASTER_ADMIN_PASSWORD##"

// Placeholder for the encrypted Jenkins master admin password (bcrypt)
const TemplateJenkinsMasterAdminPasswordEncrypted = "##JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED##"

// Placeholder for the encrypted Jenkins master user password (default project user) (bcrypt)
const TemplateJenkinsMasterUserPasswordEncrypted = "##JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED##"

// Placeholder for the Jenkins Docker image
const TemplateJenkinsMasterContainerImage = "##JENKINS_MASTER_CONTAINER_IMAGE##"

// Placeholder for the Jenkins Docker image tag
const TemplateJenkinsMasterContainerImageTag = "##JENKINS_MASTER_CONTAINER_IMAGE_TAG##"

// Placeholder for the Jenkins Docker image pull secret name
const TemplateJenkinsMasterContainerImagePullSecretName = "##JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME##"

// Placeholder for the Jenkins Docker image pull policy
const TemplateJenkinsMasterContainerPullPolicy = "##JENKINS_MASTER_CONTAINER_PULL_POLICY##"

// Placeholder for the Nginx Ingress Controller deployment name
const TemplateNginxIngressDeploymentName = "##NGINX_INGRESS_DEPLOYMENT_NAME##"

// Placeholder for the Nginx Ingress Controller Docker image
const TemplateNginxIngressControllerContainerImage = "##NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE##"

// Placeholder for the Nginx Ingress Controller Docker image pull secret
const TemplateNginxIngressControllerContainerPullSecrets = "##NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS##"

// Placeholder for Nginx Ingress Controller namespace
const TemplateNginxIngressControllerContainerForNamespace = "##NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE##"

// Placeholder for Nginx Ingress Controller Kubernetes annotation class
const TemplateNginxIngressAnnotationClass = "##NGINX_INGRESS_ANNOTATION_CLASS##"

// Placeholder for Nginx Ingress Controller, if a loadbalancer was enabled
const TemplateNginxLoadbalancerEnabled = "##NGINX_LOADBALANCER_ENABLED##"

// Placeholder for Nginx load balancer HTTP port
const TemplateNginxLoadbalancerHttpPort = "##NGINX_LOADBALANCER_HTTP_PORT##"

// Placeholder for Nginx load balancer HTTP target port
const TemplateNginxLoadbalancerHttpTargetPort = "##NGINX_LOADBALANCER_HTTP_TARGETPORT##"

// Placeholder for Nginx load balancer HTTPS port
const TemplateNginxLoadbalancerHttpsPort = "##NGINX_LOADBALANCER_HTTPS_PORT##"

// Placeholder for Nginx load balancer HTTPS target port
const TemplateNginxLoadbalancerHttpsTargetPort = "##NGINX_LOADBALANCER_HTTPS_TARGETPORT##"

// Placeholder for Kubernetes server certificate
const TemplateKubernetesServerCertificate = "##KUBERNETES_SERVER_CERTIFICATE##"

// Placeholder for Kubernetes Docker Registry credentials ID
const TemplateCredentialsIdKubernetesDockerRegistry = "##KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID##"

// Placeholder for the Maven repository secrets credentials ID
const TemplateCredentialsIdMaven = "##MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID##"

// Placeholder for the NPM repository secrets credentials ID
const TemplateCredentialsIdNpm = "##NPM_REPOSITORY_SECRETS_CREDENTIALS_ID##"

// Placeholder for the Version Control System secrets credentials ID
const TemplateCredentialsIdVcs = "##VCS_REPOSITORY_SECRETS_CREDENTIALS_ID##"

// Placeholder for the size of the PVC if the Jenkins master instance
const TemplatePvcStorageSize = "##JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE##"

// Placeholder for the Jenkins master persistence access mode (ReadWriteMany, ReadWriteSingle...)
const TemplatePvcAccessMode = "##JENKINS_MASTER_PERSISTENCE_ACCESS_MODE##"

// Placeholder for the Jenkins master persistence storage class (local-path, nfs-client...)
const TemplatePvcStorageClass = "##JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS##"

// Placeholder for the Jenkins master existing PVC
const TemplatePvcExistingVolumeClaim = "##JENKINS_MASTER_PERSISTENCE_EXISTING_CLAIM##"

// Placeholder for the PVC claim in the FilenamePvcClaim file
const TemplatePvcName = "##K8S_MGMT_PERSISTENCE_VOLUME_CLAIM_NAME##"

// GUI Constants
// GUI name of --dry-run option
const InstallDryRunActive = "dry-run"

// GUI name for execution (no dry-run)
const InstallDryRunInactive = "execute"

// ---- Utils constants
// Prefix for encrypted Jenkins user password
const UtilsJenkinsUserPassBcryptPrefix = "#jbcrypt:"

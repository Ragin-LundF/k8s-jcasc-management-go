package constants

// ---- common

// NewLine defines the new line characters
const NewLine = "\n"

// ---- directory and file configuration

// DirConfig is the config directory
const DirConfig = "config"

// DirProjectScripts is the directory for shell scripts inside a project
const DirProjectScripts = "scripts"

// DirProjectTemplateCloudTemplates is the directory for cloud-templates inside the templates directory
const DirProjectTemplateCloudTemplates = "cloud-templates"

// DirProjectScriptsInstallPrefix is the prefix for scripts for installation
const DirProjectScriptsInstallPrefix = "i_"

// DirProjectScriptsUninstallPrefix is the prefix for scripts for deinstallation
const DirProjectScriptsUninstallPrefix = "d_"

// DirHelmJenkinsMaster : directory of the Jenkins Helm Charts
const DirHelmJenkinsMaster = "/charts/jenkins-master"

// DirHelmNginxIngressCtrl : directory of the Nginx Ingress Controller Helm Charts
const DirHelmNginxIngressCtrl = "/charts/nginx-ingress-controller"

// FilenameConfiguration : filename of the configuration file
const FilenameConfiguration = "k8s_jcasc_mgmt.cnf"

// FilenameConfigurationYaml : filename of the yaml configuration file
const FilenameConfigurationYaml = "k8s_jcasc_mgmt.yaml"

// FilenameConfigurationCustom : filename of the custom configuration file for overwrites
const FilenameConfigurationCustom = "k8s_jcasc_custom.cnf"

// FilenameJenkinsConfigurationAsCode : filename of the Jenkins Configuration-as-Code Helm Values file
const FilenameJenkinsConfigurationAsCode = "jcasc_config.yaml"

// FilenameJenkinsHelmValues : filename if the Jenkins Helm Values file
const FilenameJenkinsHelmValues = "jenkins_helm_values.yaml"

// FilenameNginxIngressControllerHelmValues : filename of the Nginx Ingress Controller Helm Values file
const FilenameNginxIngressControllerHelmValues = "nginx_ingress_helm_values.yaml"

// FilenamePvcClaim : filename of the Kubernetes PVC Claim values file
const FilenamePvcClaim = "pvc_claim.yaml"

// FilenameSecrets : base filename for secrets (without .gpg extension)
const FilenameSecrets = "secrets.sh"

// SecretsFileEncodedEnding : secrets filename extension (will be added after FilenameSecrets)
const SecretsFileEncodedEnding = ".gpg"

// ScriptsFileEnding : filename extension of scripts
const ScriptsFileEnding = ".sh"

// ---- commands (for delegation of the actions)

// CommandInstall : install command
const CommandInstall = "install"

// CommandUninstall : uninstall command
const CommandUninstall = "uninstall"

// CommandUpgrade : upgrade command
const CommandUpgrade = "upgrade"

// CommandEncryptSecrets : encrypt secrets command
const CommandEncryptSecrets = "encryptSecrets"

// CommandDecryptSecrets : decrypt secrets command
const CommandDecryptSecrets = "decryptSecrets"

// CommandApplySecrets : apply secrets command
const CommandApplySecrets = "applySecrets"

// CommandApplySecretsToAll : apply secrets to all namespaces command
const CommandApplySecretsToAll = "applySecretsToAll"

// CommandCreateProject : create project command
const CommandCreateProject = "createProject"

// CommandCreateDeploymentOnlyProject : create deployment only project command
const CommandCreateDeploymentOnlyProject = "createDeploymentOnlyProject"

// CommandCreateNamespace : create namespace command
const CommandCreateNamespace = "createNamespace"

// CommandCreateJenkinsUserPassword : create Jenkins user password command
const CommandCreateJenkinsUserPassword = "createJenkinsUserPassword"

// CommandTools : tools section in the main menu
const CommandTools = "tools"

// CommandQuit : quit command
const CommandQuit = "quit"

// ---- helm commands (for execution)

// HelmCommandInstall : Helm command for install
const HelmCommandInstall = "install"

// HelmCommandUpgrade : Helm command for upgrade
const HelmCommandUpgrade = "upgrade"

// HelmCommandUninstall : Helm command for uninstall
const HelmCommandUninstall = "uninstall"

// CommonJenkinsSystemMessage : common messages
const CommonJenkinsSystemMessage = "Welcome to Jenkins"

// ErrorPromptFailed : error in CLI prompt
const ErrorPromptFailed = "prompt failed"

// ---- colors

// ColorNormal : Normal color
const ColorNormal = "\033[0m"

// ColorInfo : Info color
const ColorInfo = "\033[1;34m"

// ColorError : Error color
const ColorError = "\033[1;31m"

// KubectlFieldName : kubectl field names
const KubectlFieldName = "NAME"

// ---- Template Strings

// TemplateNamespace : Placeholder for namespace
const TemplateNamespace = "##NAMESPACE##"

// TemplatePublicIPAddress : Placeholder for public IP address
const TemplatePublicIPAddress = "##PUBLIC_IP_ADDRESS##"

// TemplateJenkinsUrl : Jenkins URL or IP
const TemplateJenkinsUrl = "##JENKINS_URL##"

// TemplateProjectDirectory : Placeholder for project directory
const TemplateProjectDirectory = "##PROJECT_DIRECTORY##"

// TemplateJenkinsCloudTemplates : Placeholder to add cloud templates
const TemplateJenkinsCloudTemplates = "##K8S_MGMT_JENKINS_CLOUD_TEMPLATES##"

// TemplateJenkinsSystemMessage : Placeholder for Jenkins system message
const TemplateJenkinsSystemMessage = "##K8S_MGMT_JENKINS_SYSTEM_MESSAGE##"

// TemplateJenkinsJobDslSeedJobScriptURL : Placeholder for URL of JobDSL seed Job script
const TemplateJenkinsJobDslSeedJobScriptURL = "##JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL##"

// TemplateJobDefinitionRepository : Placeholder for the Jobs repository
const TemplateJobDefinitionRepository = "##PROJECT_JENKINS_JOB_DEFINITION_REPOSITORY##"

// TemplateJenkinsJcascConfigurationURL : Placeholder for Jenkins Configuration-as-Code configuration URL
const TemplateJenkinsJcascConfigurationURL = "##JENKINS_JCASC_CONFIGURATION_URL##"

// TemplateJenkinsMasterDeploymentName : Placeholder for the deployment name
const TemplateJenkinsMasterDeploymentName = "##JENKINS_MASTER_DEPLOYMENT_NAME##"

// TemplateJenkinsMasterDenyAnonymousReadAccess : Placeholder for anonymous read access true/false configuration
const TemplateJenkinsMasterDenyAnonymousReadAccess = "##JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS##"

// TemplateJenkinsMasterDefaultURIPrefix : Placeholder for the Jenkins URI prefix (e.g. /jenkins)
const TemplateJenkinsMasterDefaultURIPrefix = "##JENKINS_MASTER_DEFAULT_URI_PREFIX##"

// TemplateJenkinsMasterDefaultLabel : Placeholder for label of the master instance of Jenkins to pin the seed job on this instance
const TemplateJenkinsMasterDefaultLabel = "##JENKINS_MASTER_DEFAULT_LABEL##"

// TemplateJenkinsMasterAdminPassword : Placeholder for the Jenkins master admin password (unencrypted), if no acl script was installed
const TemplateJenkinsMasterAdminPassword = "##JENKINS_MASTER_ADMIN_PASSWORD##"

// TemplateJenkinsMasterAdminPasswordEncrypted : Placeholder for the encrypted Jenkins master admin password (bcrypt)
const TemplateJenkinsMasterAdminPasswordEncrypted = "##JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED##"

// TemplateJenkinsMasterUserPasswordEncrypted : Placeholder for the encrypted Jenkins master user password (default project user) (bcrypt)
const TemplateJenkinsMasterUserPasswordEncrypted = "##JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED##"

// TemplateJenkinsMasterContainerImage : Placeholder for the Jenkins Docker image
const TemplateJenkinsMasterContainerImage = "##JENKINS_MASTER_CONTAINER_IMAGE##"

// TemplateJenkinsMasterContainerImageTag : Placeholder for the Jenkins Docker image tag
const TemplateJenkinsMasterContainerImageTag = "##JENKINS_MASTER_CONTAINER_IMAGE_TAG##"

// TemplateJenkinsMasterContainerImagePullSecretName : Placeholder for the Jenkins Docker image pull secret name
const TemplateJenkinsMasterContainerImagePullSecretName = "##JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME##"

// TemplateJenkinsMasterContainerPullPolicy : Placeholder for the Jenkins Docker image pull policy
const TemplateJenkinsMasterContainerPullPolicy = "##JENKINS_MASTER_CONTAINER_PULL_POLICY##"

// TemplateNginxIngressDeploymentName : Placeholder for the Nginx Ingress Controller deployment name
const TemplateNginxIngressDeploymentName = "##NGINX_INGRESS_DEPLOYMENT_NAME##"

// TemplateNginxIngressControllerContainerImage : Placeholder for the Nginx Ingress Controller Docker image
const TemplateNginxIngressControllerContainerImage = "##NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE##"

// TemplateNginxIngressControllerContainerPullSecrets : Placeholder for the Nginx Ingress Controller Docker image pull secret
const TemplateNginxIngressControllerContainerPullSecrets = "##NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS##"

// TemplateNginxIngressControllerContainerForNamespace : Placeholder for Nginx Ingress Controller namespace
const TemplateNginxIngressControllerContainerForNamespace = "##NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE##"

// TemplateNginxIngressAnnotationClass : Placeholder for Nginx Ingress Controller Kubernetes annotation class
const TemplateNginxIngressAnnotationClass = "##NGINX_INGRESS_ANNOTATION_CLASS##"

// TemplateNginxLoadbalancerEnabled : Placeholder for Nginx Ingress Controller, if a loadbalancer was enabled
const TemplateNginxLoadbalancerEnabled = "##NGINX_LOADBALANCER_ENABLED##"

// TemplateNginxLoadbalancerHTTPPort : Placeholder for Nginx load balancer HTTP port
const TemplateNginxLoadbalancerHTTPPort = "##NGINX_LOADBALANCER_HTTP_PORT##"

// TemplateNginxLoadbalancerHTTPTargetPort : Placeholder for Nginx load balancer HTTP target port
const TemplateNginxLoadbalancerHTTPTargetPort = "##NGINX_LOADBALANCER_HTTP_TARGETPORT##"

// TemplateNginxLoadbalancerHTTPSPort : Placeholder for Nginx load balancer HTTPS port
const TemplateNginxLoadbalancerHTTPSPort = "##NGINX_LOADBALANCER_HTTPS_PORT##"

// TemplateNginxLoadbalancerHTTPSTargetPort : Placeholder for Nginx load balancer HTTPS target port
const TemplateNginxLoadbalancerHTTPSTargetPort = "##NGINX_LOADBALANCER_HTTPS_TARGETPORT##"

// TemplateNginxLoadbalancerAnnotationsEnabled : Placeholder for Nginx Annotation support
const TemplateNginxLoadbalancerAnnotationsEnabled = "##NGINX_LOADBALANCER_ANNOTATIONS_ENABLED##"

// TemplateNginxLoadbalancerAnnotationsExtDnsTtl : Placeholder for external DNS TTL annotation
const TemplateNginxLoadbalancerAnnotationsExtDnsTtl = "##NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_TTL##"

// TemplateNginxLoadbalancerAnnotationsExtDnsTtl : Placeholder for external DNS hostname annotation
const TemplateNginxLoadbalancerAnnotationsExtDnsHostname = "##NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME##"

// TemplateKubernetesServerCertificate : Placeholder for Kubernetes server certificate
const TemplateKubernetesServerCertificate = "##KUBERNETES_SERVER_CERTIFICATE##"

// TemplateCredentialsIDKubernetesDockerRegistry : Placeholder for Kubernetes Docker Registry credentials ID
const TemplateCredentialsIDKubernetesDockerRegistry = "##KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID##"

// TemplateCredentialsIDMaven : Placeholder for the Maven repository secrets credentials ID
const TemplateCredentialsIDMaven = "##MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID##"

// TemplateCredentialsIDNpm : Placeholder for the NPM repository secrets credentials ID
const TemplateCredentialsIDNpm = "##NPM_REPOSITORY_SECRETS_CREDENTIALS_ID##"

// TemplateCredentialsIDVcs : Placeholder for the Version Control System secrets credentials ID
const TemplateCredentialsIDVcs = "##VCS_REPOSITORY_SECRETS_CREDENTIALS_ID##"

// TemplatePvcStorageSize : Placeholder for the size of the PVC if the Jenkins master instance
const TemplatePvcStorageSize = "##JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE##"

// TemplatePvcAccessMode : Placeholder for the Jenkins master persistence access mode (ReadWriteMany, ReadWriteSingle...)
const TemplatePvcAccessMode = "##JENKINS_MASTER_PERSISTENCE_ACCESS_MODE##"

// TemplatePvcStorageClass : Placeholder for the Jenkins master persistence storage class (local-path, nfs-client...)
const TemplatePvcStorageClass = "##JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS##"

// TemplatePvcExistingVolumeClaim : Placeholder for the Jenkins master existing PVC
const TemplatePvcExistingVolumeClaim = "##JENKINS_MASTER_PERSISTENCE_EXISTING_CLAIM##"

// TemplatePvcName : Placeholder for the PVC claim in the FilenamePvcClaim file
const TemplatePvcName = "##K8S_MGMT_PERSISTENCE_VOLUME_CLAIM_NAME##"

// ---- GUI Constants

// InstallDryRunActive : GUI name of --dry-run option
const InstallDryRunActive = "dry-run"

// InstallDryRunInactive : GUI name for execution (no dry-run)
const InstallDryRunInactive = "execute"

// ---- Utils constants

// UtilsJenkinsUserPassBcryptPrefix : Prefix for encrypted Jenkins user password
const UtilsJenkinsUserPassBcryptPrefix = "#jbcrypt:"

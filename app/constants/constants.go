package constants

// common
const NewLine = "\n"

// directory and file configuration
const DirConfig = "config"
const DirProjectScripts = "scripts"
const DirProjectTemplateCloudTemplates = "cloud-templates"
const DirProjectScriptsInstallPrefix = "i_"
const DirProjectScriptsUninstallPrefix = "d_"
const DirHelmJenkinsMaster = "/charts/jenkins-master"
const DirHelmNginxIngressCtrl = "/charts/nginx-ingress-controller"
const FilenameConfiguration = "k8s_jcasc_mgmt.cnf"
const FilenameConfigurationCustom = "k8s_jcasc_custom.cnf"
const FilenameJenkinsConfigurationAsCode = "jcasc_config.yaml"
const FilenameJenkinsHelmValues = "jenkins_helm_values.yaml"
const FilenameNginxIngressControllerHelmValues = "nginx_ingress_helm_values.yaml"
const FilenamePvcClaim = "pvc_claim.yaml"
const FilenameSecrets = "secrets.sh"
const SecretsFileEncodedEnding = ".gpg"
const ScriptsFileEnding = ".sh"

// commands
const CommandMenu = "menu"
const CommandInstall = "install"
const CommandUninstall = "uninstall"
const CommandUpgrade = "upgrade"
const CommandEncryptSecrets = "encryptSecrets"
const CommandDecryptSecrets = "decryptSecrets"
const CommandApplySecrets = "applySecrets"
const CommandApplySecretsToAll = "applySecretsToAll"
const CommandCreateProject = "createProject"
const CommandCreateDeploymentOnlyProject = "createDeploymentOnlyProject"
const CommandCreateJenkinsUserPassword = "createJenkinsUserPassword"
const CommandQuit = "quit"

// helm commands
const HelmCommandInstall = "install"
const HelmCommandUpgrade = "upgrade"

// error
const ErrorPromptFailed = "prompt failed"

// colors
const ColorNormal = "\033[0m"
const ColorInfo = "\033[1;34m"
const ColorNotice = "\033[1;36m"
const ColorWarning = "\033[1;33m"
const ColorError = "\033[1;31m"
const ColorDebug = "\033[0;36m"

// kubectl field names
const KubectlFieldName = "NAME"

// Template Strings
const TemplateJenkinsCloudTemplates = "##K8S_MGMT_JENKINS_CLOUD_TEMPLATES##"
const TemplateJenkinsSystemMessage = "##K8S_MGMT_JENKINS_SYSTEM_MESSAGE##"
const TemplateJenkinsJobDslSeedJobScriptUrl = "##JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL##"
const TemplateJenkinsMasterDeploymentName = "##JENKINS_MASTER_DEPLOYMENT_NAME##"
const TemplateJenkinsMasterDefaultUriPrefix = "##JENKINS_MASTER_DEFAULT_URI_PREFIX##"
const TemplateNginxIngressDeploymentName = "##NGINX_INGRESS_DEPLOYMENT_NAME##"
const TemplateNginxIngressControllerContainerImage = "##NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE##"
const TemplateNginxIngressControllerContainerPullSecrets = "##NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS##"
const TemplateNginxIngressControllerContainerForNamespace = "##NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE##"
const TemplateNginxIngressAnnotationClass = "##NGINX_INGRESS_ANNOTATION_CLASS##"
const TemplateNginxLoadbalancerEnabled = "##NGINX_LOADBALANCER_ENABLED##"
const TemplateNginxLoadbalancerHttpPort = "##NGINX_LOADBALANCER_HTTP_PORT##"
const TemplateNginxLoadbalancerHttpTargetPort = "##NGINX_LOADBALANCER_HTTP_TARGETPORT##"
const TemplateNginxLoadbalancerHttpsPort = "##NGINX_LOADBALANCER_HTTPS_PORT##"
const TemplateNginxLoadbalancerHttpsTargetPort = "##NGINX_LOADBALANCER_HTTPS_TARGETPORT##"

package models

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestGetAlternativeSecretsFilesEmpty(t *testing.T) {
	assert.Empty(t, GetSecretsFiles())
}

func TestGetAlternativeSecretsFilesWithAlternatives(t *testing.T) {
	configuration = GetConfiguration()
	configuration.GlobalSecretsFile = "./secrets.sh"

	var secretsFile = GetGlobalSecretsFile()
	var secretsFileA = strings.Replace(secretsFile, "secrets.sh", "secrets_environment_a.sh", -1)
	var secretsFileB = strings.Replace(secretsFile, "secrets.sh", "secrets_environment_b.sh", -1)
	_ = ioutil.WriteFile(secretsFileA, []byte(""), 0644)
	_ = ioutil.WriteFile(secretsFileB, []byte(""), 0644)

	var alternativeSecretFiles = GetSecretsFiles()
	assert.NotEmpty(t, alternativeSecretFiles)
	assert.True(t, len(*alternativeSecretFiles) == 2)

	os.Remove(secretsFileA)
	os.Remove(secretsFileB)
}

func TestAssignDryRun(t *testing.T) {
	assert.False(t, GetConfiguration().K8sManagement.DryRunOnly)
}

func TestAssignDryRunTrue(t *testing.T) {
	AssignDryRun(true)
	assert.True(t, GetConfiguration().K8sManagement.DryRunOnly)
}

func TestAssignDryRunFalse(t *testing.T) {
	AssignDryRun(false)
	assert.False(t, GetConfiguration().K8sManagement.DryRunOnly)
}

func TestAssignCliOnlyMode(t *testing.T) {
	assert.False(t, GetConfiguration().CliOnly)
}

func TestAssignCliOnlyModeTrue(t *testing.T) {
	AssignCliOnlyMode(true)
	assert.True(t, GetConfiguration().CliOnly)
}

func TestAssignCliOnlyModeFalse(t *testing.T) {
	AssignCliOnlyMode(false)
	assert.False(t, GetConfiguration().CliOnly)
}

func TestAssignToConfigurationCommon(t *testing.T) {
	var logLevel = "info"
	AssignToConfiguration("LOG_LEVEL", logLevel)
	var globalSecretsFile = "/tmp/secrets.sh"
	AssignToConfiguration("GLOBAL_SECRETS_FILE", globalSecretsFile)
	var ipConfigFile = "/tmp/ip_config.cnf"
	AssignToConfiguration("IP_CONFIG_FILE", ipConfigFile)
	var ipConfigDummyPrefix = "dummy_"
	AssignToConfiguration("IP_CONFIG_FILE_DUMMY_PREFIX", ipConfigDummyPrefix)
	var projectsBaseDirectory = "/tmp/prj"
	AssignToConfiguration("PROJECTS_BASE_DIRECTORY", projectsBaseDirectory)
	var templatesBaseDirectory = "/tmp/tmpl"
	AssignToConfiguration("TEMPLATES_BASE_DIRECTORY", templatesBaseDirectory)

	assert.Equal(t, logLevel, GetConfiguration().LogLevel)
	assert.Equal(t, globalSecretsFile, GetConfiguration().GlobalSecretsFile)
	assert.Equal(t, ipConfigFile, GetConfiguration().IPConfig.IPConfigFile)
	assert.Equal(t, ipConfigDummyPrefix, GetConfiguration().IPConfig.IPConfigFileDummyPrefix)
	assert.Equal(t, projectsBaseDirectory, GetConfiguration().Directories.ProjectsBaseDirectory)
	assert.Equal(t, templatesBaseDirectory, GetConfiguration().Directories.TemplatesBaseDirectory)
}

func TestAssignToConfigurationJcasc(t *testing.T) {
	var jenkinsJcascConfigURL = "https://domain.tld/repo/jobs.git"
	AssignToConfiguration("JENKINS_JCASC_CONFIGURATION_URL", jenkinsJcascConfigURL)

	assert.Equal(t, jenkinsJcascConfigURL, GetConfiguration().Jenkins.JCasC.ConfigurationURL)
}

func TestAssignToConfigurationJobDsl(t *testing.T) {
	var jenkinsJobDslBaseURL = "http://github.com"
	AssignToConfiguration("JENKINS_JOBDSL_BASE_URL", jenkinsJobDslBaseURL)
	var jenkinsJobDslRepoValidatePattern = "([A..Z])"
	AssignToConfiguration("JENKINS_JOBDSL_REPO_VALIDATE_PATTERN", jenkinsJobDslRepoValidatePattern)
	var jenkinsJobDslSeedJobScriptURL = "https://domain.tld/seed_jobs/jobs.git"
	AssignToConfiguration("JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL", jenkinsJobDslSeedJobScriptURL)

	assert.Equal(t, jenkinsJobDslBaseURL, GetConfiguration().Jenkins.JobDSL.BaseURL)
	assert.Equal(t, jenkinsJobDslRepoValidatePattern, GetConfiguration().Jenkins.JobDSL.RepoValidatePattern)
	assert.Equal(t, jenkinsJobDslSeedJobScriptURL, GetConfiguration().Jenkins.JobDSL.SeedJobScriptURL)
}

func TestAssignToConfigurationJenkinsHelm(t *testing.T) {
	var jenkinsMasterAdminPassword = "mypass"
	AssignToConfiguration("JENKINS_MASTER_ADMIN_PASSWORD", jenkinsMasterAdminPassword)
	var jenkinsMasterAdminPasswordEncrypted = "abcEncrypted"
	AssignToConfiguration("JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED", jenkinsMasterAdminPasswordEncrypted)
	var jenkinsMasterProjectUserPasswordEncrypted = "abcUserEncrypted"
	AssignToConfiguration("JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED", jenkinsMasterProjectUserPasswordEncrypted)
	var jenkinsMasterDefaultLabel = "jenkins-master"
	AssignToConfiguration("JENKINS_MASTER_DEFAULT_LABEL", jenkinsMasterDefaultLabel)
	var jenkinsMasterDenyAnonymousReadAccess = "true"
	AssignToConfiguration("JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS", jenkinsMasterDenyAnonymousReadAccess)
	var jenkinsMasterDefaultURIPrefix = "/jenkins"
	AssignToConfiguration("JENKINS_MASTER_DEFAULT_URI_PREFIX", jenkinsMasterDefaultURIPrefix)
	var jenkinsMasterDeploymentName = "jenkins-master"
	AssignToConfiguration("JENKINS_MASTER_DEPLOYMENT_NAME", jenkinsMasterDeploymentName)

	assert.Equal(t, jenkinsMasterAdminPassword, GetConfiguration().Jenkins.Helm.Master.AdminPassword)
	assert.Equal(t, jenkinsMasterAdminPasswordEncrypted, GetConfiguration().Jenkins.Helm.Master.AdminPasswordEncrypted)
	assert.Equal(t, jenkinsMasterProjectUserPasswordEncrypted, GetConfiguration().Jenkins.Helm.Master.DefaultProjectUserPasswordEncrypted)
	assert.Equal(t, jenkinsMasterDefaultLabel, GetConfiguration().Jenkins.Helm.Master.Label)
	assert.Equal(t, jenkinsMasterDenyAnonymousReadAccess, GetConfiguration().Jenkins.Helm.Master.DenyAnonymousReadAccess)
	assert.Equal(t, jenkinsMasterDefaultURIPrefix, GetConfiguration().Jenkins.Helm.Master.DefaultURIPrefix)
	assert.Equal(t, jenkinsMasterDeploymentName, GetConfiguration().Jenkins.Helm.Master.DeploymentName)

}

func TestAssignToConfigurationJenkinsHelmContainer(t *testing.T) {
	var jenkinsMasterContainerImage = "jenkins-lts"
	AssignToConfiguration("JENKINS_MASTER_CONTAINER_IMAGE", jenkinsMasterContainerImage)
	var jenkinsMasterContainerImageTag = "latest"
	AssignToConfiguration("JENKINS_MASTER_CONTAINER_IMAGE_TAG", jenkinsMasterContainerImageTag)
	var jenkinsMasterContainerPullPolicy = "always"
	AssignToConfiguration("JENKINS_MASTER_CONTAINER_PULL_POLICY", jenkinsMasterContainerPullPolicy)
	var jenkinsMasterContainerImagePullSecretName = "pull_secret"
	AssignToConfiguration("JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME", jenkinsMasterContainerImagePullSecretName)

	assert.Equal(t, jenkinsMasterContainerImage, GetConfiguration().Jenkins.Helm.Master.Container.Image)
	assert.Equal(t, jenkinsMasterContainerImageTag, GetConfiguration().Jenkins.Helm.Master.Container.ImageTag)
	assert.Equal(t, jenkinsMasterContainerPullPolicy, GetConfiguration().Jenkins.Helm.Master.Container.PullPolicy)
	assert.Equal(t, jenkinsMasterContainerImagePullSecretName, GetConfiguration().Jenkins.Helm.Master.Container.PullSecretName)
}

func TestAssignToConfigurationJenkinsHelmPersistence(t *testing.T) {
	var jenkinsMasterPvcStorageClass = "local-path"
	AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS", jenkinsMasterPvcStorageClass)
	var jenkinsMasterPvcAccessMode = "ReadWriteMany"
	AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_ACCESS_MODE", jenkinsMasterPvcAccessMode)
	var jenkinsMasterPvcSize = "3"
	AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE", jenkinsMasterPvcSize)

	assert.Equal(t, jenkinsMasterPvcStorageClass, GetConfiguration().Jenkins.Helm.Master.Persistence.StorageClass)
	assert.Equal(t, jenkinsMasterPvcAccessMode, GetConfiguration().Jenkins.Helm.Master.Persistence.AccessMode)
	assert.Equal(t, jenkinsMasterPvcSize, GetConfiguration().Jenkins.Helm.Master.Persistence.Size)
}

func TestAssignToConfigurationNginxIngressController(t *testing.T) {
	var nginxIngressControllerContainerName = "nginx"
	AssignToConfiguration("NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE", nginxIngressControllerContainerName)
	var nginxIngressControllerContainerPullSecret = "ngPsecret"
	AssignToConfiguration("NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS", nginxIngressControllerContainerPullSecret)
	var nginxIngressControllerForNamespace = "true"
	AssignToConfiguration("NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE", nginxIngressControllerForNamespace)
	var nginxIngressControllerDeploymentName = "nginx-ingress-ctrl"
	AssignToConfiguration("NGINX_INGRESS_DEPLOYMENT_NAME", nginxIngressControllerDeploymentName)
	var nginxIngressControllerAnnotationClass = "nginx-annotation"
	AssignToConfiguration("NGINX_INGRESS_ANNOTATION_CLASS", nginxIngressControllerAnnotationClass)

	assert.Equal(t, nginxIngressControllerContainerName, GetConfiguration().Nginx.Ingress.Controller.Container.Name)
	assert.Equal(t, nginxIngressControllerContainerPullSecret, GetConfiguration().Nginx.Ingress.Controller.Container.PullSecret)
	nginxIngressControllerForNamespaceBool, _ := strconv.ParseBool(nginxIngressControllerForNamespace)
	assert.Equal(t, nginxIngressControllerForNamespaceBool, GetConfiguration().Nginx.Ingress.Controller.Container.Namespace)
	assert.Equal(t, nginxIngressControllerDeploymentName, GetConfiguration().Nginx.Ingress.Controller.DeploymentName)
	assert.Equal(t, nginxIngressControllerAnnotationClass, GetConfiguration().Nginx.Ingress.AnnotationClass)
}

func TestAssignToConfigurationNginxLoadbalancer(t *testing.T) {
	var nginxLoadbalancerEnabled = "true"
	AssignToConfiguration("NGINX_LOADBALANCER_ENABLED", nginxLoadbalancerEnabled)
	var nginxLoadbalancerHTTPPort = "80"
	AssignToConfiguration("NGINX_LOADBALANCER_HTTP_PORT", nginxLoadbalancerHTTPPort)
	var nginxLoadbalancerHTTPTargetPort = "8080"
	AssignToConfiguration("NGINX_LOADBALANCER_HTTP_TARGETPORT", nginxLoadbalancerHTTPTargetPort)
	var nginxLoadbalancerHTTPSPort = "443"
	AssignToConfiguration("NGINX_LOADBALANCER_HTTPS_PORT", nginxLoadbalancerHTTPSPort)
	var nginxLoadbalancerHTTPSTargetPort = "8443"
	AssignToConfiguration("NGINX_LOADBALANCER_HTTPS_TARGETPORT", nginxLoadbalancerHTTPSTargetPort)

	nginxLoadbalancerEnabledBool, _ := strconv.ParseBool(nginxLoadbalancerEnabled)
	assert.Equal(t, nginxLoadbalancerEnabledBool, GetConfiguration().LoadBalancer.Enabled)
	nginxLoadbalancerHTTPPortInt, _ := strconv.ParseUint(nginxLoadbalancerHTTPPort, 10, 16)
	assert.Equal(t, nginxLoadbalancerHTTPPortInt, GetConfiguration().LoadBalancer.Port.HTTP)
	nginxLoadbalancerHTTPTargetPortInt, _ := strconv.ParseUint(nginxLoadbalancerHTTPTargetPort, 10, 16)
	assert.Equal(t, nginxLoadbalancerHTTPTargetPortInt, GetConfiguration().LoadBalancer.Port.HTTPTarget)
	nginxLoadbalancerHTTPSPortInt, _ := strconv.ParseUint(nginxLoadbalancerHTTPSPort, 10, 16)
	assert.Equal(t, nginxLoadbalancerHTTPSPortInt, GetConfiguration().LoadBalancer.Port.HTTPS)
	nginxLoadbalancerHTTPSTargetPortInt, _ := strconv.ParseUint(nginxLoadbalancerHTTPSTargetPort, 10, 16)
	assert.Equal(t, nginxLoadbalancerHTTPSTargetPortInt, GetConfiguration().LoadBalancer.Port.HTTPSTarget)
}

func TestAssignToConfigurationCredentials(t *testing.T) {
	var k8sDockerRegistryCredentialsID = "docker-reg"
	AssignToConfiguration("KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID", k8sDockerRegistryCredentialsID)
	var credentialsIDMaven = "mvn-cred-id"
	AssignToConfiguration("MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID", credentialsIDMaven)
	var credentialsIDNpm = "npm-cred-id"
	AssignToConfiguration("NPM_REPOSITORY_SECRETS_CREDENTIALS_ID", credentialsIDNpm)
	var credentialsIDVcs = "vcs-cred-id"
	AssignToConfiguration("VCS_REPOSITORY_SECRETS_CREDENTIALS_ID", credentialsIDVcs)

	assert.Equal(t, k8sDockerRegistryCredentialsID, GetConfiguration().CredentialIds.DefaultDockerRegistry)
	assert.Equal(t, credentialsIDMaven, GetConfiguration().CredentialIds.DefaultMavenRepository)
	assert.Equal(t, credentialsIDNpm, GetConfiguration().CredentialIds.DefaultNpmRepository)
	assert.Equal(t, credentialsIDVcs, GetConfiguration().CredentialIds.DefaultVcsRepository)
}

func TestAssignToConfigurationK8sMgmt(t *testing.T) {
	var k8sMgmtVersionCheck = "true"
	AssignToConfiguration("K8S_MGMT_VERSION_CHECK", k8sMgmtVersionCheck)
	var k8sMgmtAlternativeConfigFile = "k8s_custom.cnf"
	AssignToConfiguration("K8S_MGMT_ALTERNATIVE_CONFIG_FILE", k8sMgmtAlternativeConfigFile)
	var k8sMgmtBasePath = "./basePath"
	AssignToConfiguration("K8S_MGMT_BASE_PATH", k8sMgmtBasePath)

	k8sMgmtVersionCheckBool, _ := strconv.ParseBool(k8sMgmtVersionCheck)
	assert.Equal(t, k8sMgmtVersionCheckBool, GetConfiguration().K8sManagement.VersionCheck)
	assert.Equal(t, k8sMgmtAlternativeConfigFile, GetConfiguration().AlternativeConfigFile)
	assert.Equal(t, k8sMgmtBasePath, GetConfiguration().BasePath)
}

func TestAssignToConfigurationK8sMgmtLogging(t *testing.T) {
	var k8sMgmtLoggingLogfile = "output_data.log"
	AssignToConfiguration("K8S_MGMT_LOGGING_LOGFILE", k8sMgmtLoggingLogfile)
	var k8sMgmtLoggingEncoding = "console"
	AssignToConfiguration("K8S_MGMT_LOGGING_ENCODING", k8sMgmtLoggingEncoding)
	var k8sMgmtLoggingOverwriteOnStart = "true"
	AssignToConfiguration("K8S_MGMT_LOGGING_OVERWRITE_ON_START", k8sMgmtLoggingOverwriteOnStart)

	assert.Equal(t, k8sMgmtLoggingLogfile, GetConfiguration().K8sManagement.Logging.LogFile)
	assert.Equal(t, k8sMgmtLoggingEncoding, GetConfiguration().K8sManagement.Logging.LogEncoding)
	k8sMgmtLoggingOverwriteOnStartBool, _ := strconv.ParseBool(k8sMgmtLoggingOverwriteOnStart)
	assert.Equal(t, k8sMgmtLoggingOverwriteOnStartBool, GetConfiguration().K8sManagement.Logging.LogOverwriteOnStart)
}

func TestReplaceUnneededChars(t *testing.T) {
	var resultString = replaceUnneededChars("test")

	assert.Equal(t, "test", resultString)
}

func TestReplaceUnneededCharsSingleQuote(t *testing.T) {
	var resultString = replaceUnneededChars("'test'")

	assert.Equal(t, "test", resultString)
}

func TestReplaceUnneededCharsDoubleQuote(t *testing.T) {
	var resultString = replaceUnneededChars("\"test\"")

	assert.Equal(t, "test", resultString)
}

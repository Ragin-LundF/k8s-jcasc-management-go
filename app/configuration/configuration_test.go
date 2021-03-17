package configuration

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestGetAlternativeSecretsFilesEmpty(t *testing.T) {
	var conf = EmptyConfiguration()
	conf.K8SManagement.BasePath = "./"
	conf.K8SManagement.Project.SecretFiles = "./secrets.sh"

	assert.True(t, len(conf.GetSecretsFiles()) == 1)
}

func TestGetAlternativeSecretsFilesWithDifferentBasePath(t *testing.T) {
	var conf = EmptyConfiguration()
	conf.K8SManagement.BasePath = "../"
	conf.K8SManagement.Project.SecretFiles = "./secrets.sh"

	_ = ioutil.WriteFile("../secrets.sh", []byte(""), 0644)
	var secretsFile = conf.getGlobalSecretsFile()
	var secretsFileA = strings.Replace(secretsFile, "secrets.sh", "secrets_environment_a.sh", -1)
	var secretsFileB = strings.Replace(secretsFile, "secrets.sh", "secrets_environment_b.sh", -1)
	_ = ioutil.WriteFile(secretsFile, []byte(""), 0644)
	_ = ioutil.WriteFile(secretsFileA, []byte(""), 0644)
	_ = ioutil.WriteFile(secretsFileB, []byte(""), 0644)

	var alternativeSecretFiles = conf.GetSecretsFiles()
	assert.NotEmpty(t, alternativeSecretFiles)
	assert.True(t, len(alternativeSecretFiles) == 3)
	for _, secretFile := range alternativeSecretFiles {
		runeSecretFile := []rune(secretFile)
		assert.Equal(t, "secrets", string(runeSecretFile[0:7]))
	}

	_ = os.Remove(secretsFileA)
	_ = os.Remove(secretsFileB)
	_ = os.Remove("../secrets.sh")
}

func TestLoadConfiguration(t *testing.T) {
	LoadConfiguration("../../", true, true)
	var conf = GetConfiguration()

	assert.NotNil(t, conf)
	assert.NotNil(t, conf.GetIPConfigurationFile())
	assert.True(t, strings.HasSuffix(conf.GetIPConfigurationFile(), "/config/ip_config.yaml"))
	assert.NotNil(t, conf.GetGlobalSecretsPath())
	assert.True(t, strings.HasSuffix(conf.GetGlobalSecretsPath(), "/config/"))
	assert.NotNil(t, conf.getGlobalSecretsFile())
	assert.True(t, strings.HasSuffix(conf.getGlobalSecretsFile(), "/config/secrets.sh"))
	assert.NotNil(t, conf.GetProjectBaseDirectory())
	assert.True(t, strings.HasSuffix(conf.GetProjectBaseDirectory(), "/projects/"))
	assert.NotNil(t, conf.GetProjectTemplateDirectory())
	assert.True(t, strings.HasSuffix(conf.GetProjectTemplateDirectory(), "/templates/"))
	assert.NotNil(t, conf.GetSecretsFiles())
	assert.True(t, len(conf.GetSecretsFiles()) == 1)

	assertCustomConfig(conf, t)
	assertJenkinsConfig(conf, t)
	assertK8SManagement(conf, t)
	assertKubernetesConfig(conf, t)
	assertNginxConfig(conf, t)
}

func assertCustomConfig(conf *config, t *testing.T) {
	assert.NotNil(t, conf.CustomConfig)
	assert.NotNil(t, conf.CustomConfig.K8SManagement)
	assert.NotNil(t, conf.CustomConfig.K8SManagement.BasePath)
	assert.NotNil(t, conf.CustomConfig.K8SManagement.ConfigFile)
}

func assertJenkinsConfig(conf *config, t *testing.T) {
	assert.NotNil(t, conf.Jenkins)
	assert.NotNil(t, conf.Jenkins.Container)
	assert.NotNil(t, conf.Jenkins.Container.Image)
	assert.Equal(t, conf.Jenkins.Container.Image, "jenkins/jenkins")
	assert.NotNil(t, conf.Jenkins.Container.PullPolicy)
	assert.Equal(t, conf.Jenkins.Container.PullPolicy, "Always")
	assert.NotNil(t, conf.Jenkins.Container.PullSecret)
	assert.Empty(t, conf.Jenkins.Container.PullSecret)

	assert.NotNil(t, conf.Jenkins.Controller)
	assert.NotNil(t, conf.Jenkins.Controller.CustomJenkinsLabel)
	assert.Equal(t, conf.Jenkins.Controller.CustomJenkinsLabel, "jenkins-controller-for-seed")
	assert.NotNil(t, conf.Jenkins.Controller.DefaultURIPrefix)
	assert.Equal(t, conf.Jenkins.Controller.DefaultURIPrefix, "/jenkins")
	assert.NotNil(t, conf.Jenkins.Controller.DeploymentName)
	assert.Equal(t, conf.Jenkins.Controller.DeploymentName, "jenkins-controller")
	assert.NotNil(t, conf.Jenkins.Controller.Passwords)
	assert.NotNil(t, conf.Jenkins.Controller.Passwords.AdminUser)
	assert.Equal(t, conf.Jenkins.Controller.Passwords.AdminUser, "admin")
	assert.NotNil(t, conf.Jenkins.Controller.Passwords.AdminUserEncrypted)
	assert.Equal(t, conf.Jenkins.Controller.Passwords.AdminUserEncrypted, "$2a$04$UNxiNvJN6R3me9vybVQr/OzpMhgobih8qbxDpGy3lZmmmwc6t48ty")
	assert.NotNil(t, conf.Jenkins.Controller.Passwords.DefaultUserEncrypted)
	assert.Equal(t, conf.Jenkins.Controller.Passwords.DefaultUserEncrypted, "$2a$04$BFPq6fSa9KGKrlIktz/C8eSFrrG/gglnW1eXWMSjgtCSx36mMOSNm")

	assert.NotNil(t, conf.Jenkins.Jcasc)
	assert.NotNil(t, conf.Jenkins.Jcasc.AuthorizationStrategy)
	assert.NotNil(t, conf.Jenkins.Jcasc.AuthorizationStrategy.AllowAnonymousRead)
	assert.True(t, conf.Jenkins.Jcasc.AuthorizationStrategy.AllowAnonymousRead)
	assert.NotNil(t, conf.Jenkins.Jcasc.ConfigurationURL)
	assert.Equal(t, conf.Jenkins.Jcasc.ConfigurationURL, "https://raw.githubusercontent.com/Ragin-LundF/k8s-jcasc-project-config/main/{{ .Base.Namespace }}/jcasc_config.yaml")
	assert.NotNil(t, conf.Jenkins.Jcasc.CredentialIDs)
	assert.NotNil(t, conf.Jenkins.Jcasc.CredentialIDs.Docker)
	assert.Equal(t, conf.Jenkins.Jcasc.CredentialIDs.Docker, "docker-registry-credentialsid")
	assert.NotNil(t, conf.Jenkins.Jcasc.CredentialIDs.Maven)
	assert.Equal(t, conf.Jenkins.Jcasc.CredentialIDs.Maven, "repository-credentialsid")
	assert.NotNil(t, conf.Jenkins.Jcasc.CredentialIDs.Npm)
	assert.Equal(t, conf.Jenkins.Jcasc.CredentialIDs.Npm, "repository-credentialsid")
	assert.NotNil(t, conf.Jenkins.Jcasc.CredentialIDs.Vcs)
	assert.Equal(t, conf.Jenkins.Jcasc.CredentialIDs.Vcs, "vcs-notification-credentialsid")
	assert.NotNil(t, conf.Jenkins.Jcasc.SeedJobURL)
	assert.Equal(t, conf.Jenkins.Jcasc.SeedJobURL, "https://github.com/Ragin-LundF/jenkins-jobdsl-remote.git")

	assert.NotNil(t, conf.Jenkins.JobDSL)
	assert.NotNil(t, conf.Jenkins.JobDSL.BaseURL)
	assert.Equal(t, conf.Jenkins.JobDSL.BaseURL, "http://github.com")
	assert.NotNil(t, conf.Jenkins.JobDSL.RepoValidatePattern)
	assert.Equal(t, conf.Jenkins.JobDSL.RepoValidatePattern, ".*\\.git")

	assert.NotNil(t, conf.Jenkins.Persistence)
	assert.NotNil(t, conf.Jenkins.Persistence.AccessMode)
	assert.Equal(t, conf.Jenkins.Persistence.AccessMode, "ReadWriteOnce")
	assert.NotNil(t, conf.Jenkins.Persistence.StorageClass)
	assert.Equal(t, conf.Jenkins.Persistence.StorageClass, "nfs-client")
	assert.NotNil(t, conf.Jenkins.Persistence.StorageSize)
	assert.Equal(t, conf.Jenkins.Persistence.StorageSize, "2Gi")
}

func assertK8SManagement(conf *config, t *testing.T) {
	assert.NotNil(t, conf.K8SManagement)
	assert.NotNil(t, conf.K8SManagement.BasePath)

	assert.NotNil(t, conf.K8SManagement.Project)
	assert.NotNil(t, conf.K8SManagement.Project.BaseDirectory)
	assert.Equal(t, conf.K8SManagement.Project.BaseDirectory, "./projects/")
	assert.NotNil(t, conf.K8SManagement.Project.SecretFiles)
	assert.Equal(t, conf.K8SManagement.Project.SecretFiles, "./config/secrets.sh")
	assert.NotNil(t, conf.K8SManagement.Project.TemplateDirectory)
	assert.Equal(t, conf.K8SManagement.Project.TemplateDirectory, "./templates/")

	assert.NotNil(t, conf.K8SManagement.CliOnly)
	assert.True(t, conf.K8SManagement.CliOnly)

	assert.NotNil(t, conf.K8SManagement.DryRunOnly)
	assert.True(t, conf.K8SManagement.DryRunOnly)

	assert.NotNil(t, conf.K8SManagement.IPConfig)
	assert.NotNil(t, conf.K8SManagement.IPConfig.DummyPrefix)
	assert.Equal(t, conf.K8SManagement.IPConfig.DummyPrefix, "dummy")
	assert.NotNil(t, conf.K8SManagement.IPConfig.File)
	assert.Equal(t, conf.K8SManagement.IPConfig.File, "./config/ip_config.yaml")

	assert.NotNil(t, conf.K8SManagement.Log)
	assert.NotNil(t, conf.K8SManagement.Log.File)
	assert.Equal(t, conf.K8SManagement.Log.File, "output.log")
	assert.NotNil(t, conf.K8SManagement.Log.Encoding)
	assert.Equal(t, conf.K8SManagement.Log.Encoding, "console")
	assert.NotNil(t, conf.K8SManagement.Log.Level)
	assert.Equal(t, conf.K8SManagement.Log.Level, "INFO")
	assert.NotNil(t, conf.K8SManagement.Log.OverwriteOnRestart)
	assert.True(t, conf.K8SManagement.Log.OverwriteOnRestart)

	assert.NotNil(t, conf.K8SManagement.VersionCheck)
	assert.True(t, conf.K8SManagement.VersionCheck)
}

func assertKubernetesConfig(conf *config, t *testing.T) {
	assert.NotNil(t, conf.Kubernetes)
	assert.NotNil(t, conf.Kubernetes.Certificates)
	assert.NotNil(t, conf.Kubernetes.Certificates.Default)
	assert.True(t, len(conf.Kubernetes.Certificates.Default) > 10)
}

func assertNginxConfig(conf *config, t *testing.T) {
	assert.NotNil(t, conf.Nginx)
	assert.NotNil(t, conf.Nginx.Ingress)
	assert.NotNil(t, conf.Nginx.Ingress.Annotationclass)
	assert.Equal(t, conf.Nginx.Ingress.Annotationclass, "nginx")

	assert.NotNil(t, conf.Nginx.Ingress.Container)
	assert.NotNil(t, conf.Nginx.Ingress.Container.Image)
	assert.Equal(t, conf.Nginx.Ingress.Container.Image, "bitnami/nginx-ingress-controller:latest")
	assert.NotNil(t, conf.Nginx.Ingress.Container.PullSecret)
	assert.Empty(t, conf.Nginx.Ingress.Container.PullSecret)

	assert.NotNil(t, conf.Nginx.Ingress.Deployment)
	assert.NotNil(t, conf.Nginx.Ingress.Deployment.DeploymentName)
	assert.Equal(t, conf.Nginx.Ingress.Deployment.DeploymentName, "nginx-ingress")
	assert.NotNil(t, conf.Nginx.Ingress.Deployment.ForEachNamespace)
	assert.False(t, conf.Nginx.Ingress.Deployment.ForEachNamespace)

	assert.NotNil(t, conf.Nginx.Loadbalancer)
}

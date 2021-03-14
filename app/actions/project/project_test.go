package project

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/configuration"
	"os"
	"strings"
	"testing"
)

// ----- Constants for testing (base configuration)
var testRootDirectory = fmt.Sprintf("..%v..%v..%v", string(os.PathSeparator), string(os.PathSeparator), string(os.PathSeparator))
var testProjectRootDirectory = "projects"
var testProjectName = testRootDirectory + testProjectRootDirectory + string(os.PathSeparator) + "__k8s_project_test__"

const testNamespace = "my-namespace"
const testConfigJenkinsMasterDefaultUriPrefix = "/jenkins"
const testConfigJenkinsMasterDeploymentName = "jenkins-controller"

const testConfigJenkinsMasterPvcAccessMode = "ReadWriteOnce"
const testConfigJenkinsMasterPvcStorageClassName = "hostpath"
const testConfigJenkinsMasterPvcSize = "1Gi"

const testNginxIngressAnnotationClass = "nginx"
const testNginxIngressDeploymentName = "nginx-ingress"
const testNginxIngressControllerContainerImage = "bitnami/nginx-ingress-controller:latest"
const testNginxIngressControllerContainerPullSecrets = "mypullsecret"
const testNginxIngressControllerForNamespace = true

const testNginxLoadBalancerEnabled = true
const testNginxLoadBalancerHttpPort = uint64(80)
const testNginxLoadBalancerHttpTargetPort = uint64(8080)
const testNginxLoadBalancerHttpsPort = uint64(443)
const testNginxLoadBalancerHttpsTargetPort = uint64(8443)

const testNginxLoadBalancerAnnotationsEnabled = true
const testNginxLoadBalancerAnnotationsExtDnsHostname = "domain.tld"
const testNginxLoadBalancerAnnotationsExtDnsTtl = uint64(3600)

const testJenkinsHelmMasterImage = "jenkins/jenkins"
const testJenkinsHelmMasterImageTag = "latest"
const testJenkinsHelmMasterPullPolicy = "Always"
const testJenkinsHelmMasterPullSecret = "my-secret"

const testJenkinsHelmMasterDefaultLabel = "jenkins-controller-for-seed"

const testJenkinsHelmMasterDenyAnonymousReadAccess = false
const testJenkinsHelmMasterAdminPassword = "admin"
const testJenkinsHelmMasterAdminPasswordEncrypted = "$2a$04$UNxiNvJN6R3me9vybVQr/OzpMhgobih8qbxDpGy3lZmmmwc6t48ty"
const testJenkinsHelmMasterUserPasswordEncrypted = "$2a$04$BFPq6fSa9KGKrlIktz/C8eSFrrG/gglnW1eXWMSjgtCSx36mMOSNm"

const testJcascDockerCredentialsId = "docker-credentials"
const testJcascMavenCredentialsId = "maven-credentials"
const testJcascNpmCredentialsId = "npm-credentials"
const testJcascVcsCredentialsId = "vcs-credentials"

const testJenkinsHelmMasterJcascConfigUrl = "https://raw.githubusercontent.com/Ragin-LundF/k8s-jcasc-project-config/master/{{ .Base.Namespace }}/jcasc_config.yaml"
const testJenkinsHelmMasterJcascConfigSeedUrl = "https://seed-job.domain.tld/seed.git"

const testJcascKubernetesCertificate = "LN2-test-certificate"

func TestProjectTemplates(t *testing.T) {
	testDefaultProjectConfiguration(t, true)
	var cloudTemplates = []string{"gradle_java.yaml", "node.yaml"}
	var project = NewProject()
	project.SetNamespace(testNamespace)

	// add some default values
	project.SetJobsSeedRepository("https://my-config.domain.tld/seedJob.git")
	project.SetJobsDefinitionRepository("https://my-job-repo.domain.tld/jobs.git")
	project.SetPersistentVolumeClaimExistingName("my-pvc-claim")

	// assign cloud templates
	var cloudTemplatesString, err = ActionReadCloudTemplatesAsString(cloudTemplates)
	assert.Nil(t, err)
	project.SetCloudKubernetesAdditionalTemplates(cloudTemplatesString)

	err = project.ProcessTemplates(testProjectName)
	assert.Nil(t, err)

	// remove generated templates
	_ = os.RemoveAll(testProjectName)
}

func TestProjectValidationErrorWithEmptyIP(t *testing.T) {
	testDefaultProjectConfiguration(t, false)
	configuration.GetConfiguration().Nginx.Loadbalancer.Annotations.Enabled = false

	var project = NewProject()
	project.SetNamespace(testNamespace)
	var err = project.validateProject()

	assert.Error(t, err)
}

func TestProjectValidationErrorWithEmptyNamespace(t *testing.T) {
	testDefaultProjectConfiguration(t, false)

	var project = NewProject()
	project.SetIPAddress("127.0.0.1")
	var err = project.validateProject()

	assert.Error(t, err)
}

func testDefaultProjectConfiguration(t *testing.T, setupTestProject bool) {
	configuration.LoadConfiguration("../../../", false, false)

	var cfg = configuration.GetConfiguration()
	cfg.K8SManagement.Project.TemplateDirectory = testRootDirectory + "templates/"

	cfg.Jenkins.Controller.DeploymentName = testConfigJenkinsMasterDeploymentName
	cfg.Jenkins.Controller.DefaultURIPrefix = testConfigJenkinsMasterDefaultUriPrefix

	cfg.Jenkins.Persistence.AccessMode = testConfigJenkinsMasterPvcAccessMode
	cfg.Jenkins.Persistence.StorageSize = testConfigJenkinsMasterPvcSize
	cfg.Jenkins.Persistence.StorageClass = testConfigJenkinsMasterPvcStorageClassName

	cfg.Nginx.Ingress.Annotationclass = testNginxIngressAnnotationClass
	cfg.Nginx.Ingress.Container.Image = testNginxIngressControllerContainerImage
	cfg.Nginx.Ingress.Container.PullSecret = testNginxIngressControllerContainerPullSecrets
	cfg.Nginx.Ingress.Deployment.DeploymentName = testNginxIngressDeploymentName
	cfg.Nginx.Ingress.Deployment.ForEachNamespace = testNginxIngressControllerForNamespace

	cfg.Nginx.Loadbalancer.Enabled = testNginxLoadBalancerEnabled
	cfg.Nginx.Loadbalancer.Ports.HTTP = testNginxLoadBalancerHttpPort
	cfg.Nginx.Loadbalancer.Ports.HTTPTarget = testNginxLoadBalancerHttpTargetPort
	cfg.Nginx.Loadbalancer.Ports.HTTPS = testNginxLoadBalancerHttpsPort
	cfg.Nginx.Loadbalancer.Ports.HTTPSTarget = testNginxLoadBalancerHttpsTargetPort
	cfg.Nginx.Loadbalancer.Annotations.Enabled = testNginxLoadBalancerAnnotationsEnabled
	cfg.Nginx.Loadbalancer.ExternalDNS.HostName = testNginxLoadBalancerAnnotationsExtDnsHostname
	cfg.Nginx.Loadbalancer.ExternalDNS.TTL = testNginxLoadBalancerAnnotationsExtDnsTtl

	cfg.Jenkins.Container.Image = testJenkinsHelmMasterImage
	cfg.Jenkins.Container.Tag = testJenkinsHelmMasterImageTag
	cfg.Jenkins.Container.PullPolicy = testJenkinsHelmMasterPullPolicy
	cfg.Jenkins.Container.PullSecret = testJenkinsHelmMasterPullSecret

	cfg.Jenkins.Controller.CustomJenkinsLabel = testJenkinsHelmMasterDefaultLabel

	cfg.Jenkins.Jcasc.AuthorizationStrategy.AllowAnonymousRead = testJenkinsHelmMasterDenyAnonymousReadAccess
	cfg.Jenkins.Jcasc.ConfigurationURL = testJenkinsHelmMasterJcascConfigUrl
	cfg.Jenkins.Controller.Passwords.AdminUser = testJenkinsHelmMasterAdminPassword
	cfg.Jenkins.Controller.Passwords.AdminUserEncrypted = testJenkinsHelmMasterAdminPasswordEncrypted
	cfg.Jenkins.Controller.Passwords.DefaultUserEncrypted = testJenkinsHelmMasterUserPasswordEncrypted

	cfg.Jenkins.Jcasc.CredentialIDs.Docker = testJcascDockerCredentialsId
	cfg.Jenkins.Jcasc.CredentialIDs.Maven = testJcascMavenCredentialsId
	cfg.Jenkins.Jcasc.CredentialIDs.Npm = testJcascNpmCredentialsId
	cfg.Jenkins.Jcasc.CredentialIDs.Vcs = testJcascVcsCredentialsId

	cfg.Jenkins.Jcasc.SeedJobURL = testJenkinsHelmMasterJcascConfigSeedUrl

	cfg.Kubernetes.Certificates.Default = testJcascKubernetesCertificate

	if setupTestProject {
		var err = ActionCreateNewProjectDirectory(testProjectName)
		assert.Nil(t, err)
		err = ActionCopyTemplatesToNewDirectory(testProjectName, true, false)
		assert.Nil(t, err)
	}
}

// TestCommandExec is the test executor for mocks
type TestCommandExec struct{}

// CombinedOutput is the mock implementation of CombinedOutput
func (c TestCommandExec) CombinedOutput(command string, args ...string) ([]byte, error) {
	var result []byte
	var commandAsString = command + " " + strings.Join(args, " ")
	result = combinedOutputCurrentContext(args)
	if result != nil {
		return result, nil
	}
	result = combinedOutputGetContexts(args)
	if result != nil {
		return result, nil
	}

	return []byte(commandAsString + "...executed"), nil
}

// combinedOutputCurrentContext returns the kubernetes config current-context
func combinedOutputCurrentContext(args []string) []byte {
	if cap(args) == 2 {
		if args[0] == "config" && args[1] == "current-context" {
			return []byte("default-k8s")
		}
	}
	return nil
}

// combinedOutputGetContexts returns the kubernetes config get-contexts -o name
func combinedOutputGetContexts(args []string) []byte {
	if cap(args) == 4 {
		if args[0] == "config" && args[1] == "get-contexts" && args[2] == "-o" && args[3] == "name" {
			return []byte("default-k8s")
		}
	}
	return nil
}

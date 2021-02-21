package project

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/actions/createprojectactions"
	"k8s-management-go/app/models"
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
const testConfigJenkinsMasterDeploymentName = "jenkins-master"

const testConfigJenkinsMasterPvcAccessMode = "ReadWriteOnce"
const testConfigJenkinsMasterPvcStorageClassName = "hostpath"
const testConfigJenkinsMasterPvcSize = "1Gi"

const testNginxIngressAnnotationClass = "nginx"
const testNginxIngressDeploymentName = "nginx-ingress"
const testNginxIngressControllerContainerImage = "bitnami/nginx-ingress-controller:latest"
const testNginxIngressControllerContainerPullSecrets = "mypullsecret"
const testNginxIngressControllerForNamespace = "true"

const testNginxLoadBalancerEnabled = "true"
const testNginxLoadBalancerHttpPort = "80"
const testNginxLoadBalancerHttpTargetPort = "8080"
const testNginxLoadBalancerHttpsPort = "443"
const testNginxLoadBalancerHttpsTargetPort = "8443"

const testNginxLoadBalancerAnnotationsEnabled = "true"
const testNginxLoadBalancerAnnotationsExtDnsHostname = "domain.tld"
const testNginxLoadBalancerAnnotationsExtDnsTtl = "3600"

const testJenkinsHelmMasterImage = "jenkins/jenkins"
const testJenkinsHelmMasterImageTag = "latest"
const testJenkinsHelmMasterPullPolicy = "Always"
const testJenkinsHelmMasterPullSecret = "my-secret"

const testJenkinsHelmMasterDefaultLabel = "jenkins-master-for-seed"

const testJenkinsHelmMasterDenyAnonymousReadAccess = "true"
const testJenkinsHelmMasterAdminPassword = "admin"
const testJenkinsHelmMasterAdminPasswordEncrypted = "$2a$04$UNxiNvJN6R3me9vybVQr/OzpMhgobih8qbxDpGy3lZmmmwc6t48ty"
const testJenkinsHelmMasterUserPasswordEncrypted = "$2a$04$BFPq6fSa9KGKrlIktz/C8eSFrrG/gglnW1eXWMSjgtCSx36mMOSNm"

const testJcascDockerCredentialsId = "docker-credentials"
const testJcascMavenCredentialsId = "maven-credentials"
const testJcascNpmCredentialsId = "npm-credentials"
const testJcascVcsCredentialsId = "vcs-credentials"

const testJenkinsHelmMasterJcascConfigUrl = "https://raw.githubusercontent.com/Ragin-LundF/k8s-jcasc-project-config/master/{{ .Base.Namespace }}/jcasc_config.yaml"

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
	var cloudTemplatesString, err = createprojectactions.ActionReadCloudTemplatesAsString(cloudTemplates)
	assert.Nil(t, err)
	project.SetCloudKubernetesAdditionalTemplates(cloudTemplatesString)

	err = project.ProcessTemplates(testProjectName)
	assert.Nil(t, err)

	// _ = os.RemoveAll(testProjectName)
}

func TestProjectValidationErrorWithEmptyIP(t *testing.T) {
	testDefaultProjectConfiguration(t, false)
	models.AssignToConfiguration("NGINX_LOADBALANCER_ANNOTATIONS_ENABLED", "false")

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
	models.AssignToConfiguration("TEMPLATES_BASE_DIRECTORY", testRootDirectory+"templates/")

	models.AssignToConfiguration("JENKINS_MASTER_DEPLOYMENT_NAME", testConfigJenkinsMasterDeploymentName)
	models.AssignToConfiguration("JENKINS_MASTER_DEFAULT_URI_PREFIX", testConfigJenkinsMasterDefaultUriPrefix)

	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_ACCESS_MODE", testConfigJenkinsMasterPvcAccessMode)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE", testConfigJenkinsMasterPvcSize)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS", testConfigJenkinsMasterPvcStorageClassName)

	models.AssignToConfiguration("NGINX_INGRESS_ANNOTATION_CLASS", testNginxIngressAnnotationClass)
	models.AssignToConfiguration("NGINX_INGRESS_DEPLOYMENT_NAME", testNginxIngressDeploymentName)

	models.AssignToConfiguration("NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE", testNginxIngressControllerContainerImage)
	models.AssignToConfiguration("NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS", testNginxIngressControllerContainerPullSecrets)
	models.AssignToConfiguration("NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE", testNginxIngressControllerForNamespace)

	models.AssignToConfiguration("NGINX_LOADBALANCER_ENABLED", testNginxLoadBalancerEnabled)
	models.AssignToConfiguration("NGINX_LOADBALANCER_HTTP_PORT", testNginxLoadBalancerHttpPort)
	models.AssignToConfiguration("NGINX_LOADBALANCER_HTTP_TARGETPORT", testNginxLoadBalancerHttpTargetPort)
	models.AssignToConfiguration("NGINX_LOADBALANCER_HTTPS_PORT", testNginxLoadBalancerHttpsPort)
	models.AssignToConfiguration("NGINX_LOADBALANCER_HTTPS_TARGETPORT", testNginxLoadBalancerHttpsTargetPort)

	models.AssignToConfiguration("NGINX_LOADBALANCER_ANNOTATIONS_ENABLED", testNginxLoadBalancerAnnotationsEnabled)
	models.AssignToConfiguration("NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME", testNginxLoadBalancerAnnotationsExtDnsHostname)
	models.AssignToConfiguration("NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_TTL", testNginxLoadBalancerAnnotationsExtDnsTtl)

	models.AssignToConfiguration("JENKINS_MASTER_CONTAINER_IMAGE", testJenkinsHelmMasterImage)
	models.AssignToConfiguration("JENKINS_MASTER_CONTAINER_IMAGE_TAG", testJenkinsHelmMasterImageTag)
	models.AssignToConfiguration("JENKINS_MASTER_CONTAINER_PULL_POLICY", testJenkinsHelmMasterPullPolicy)
	models.AssignToConfiguration("JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME", testJenkinsHelmMasterPullSecret)

	models.AssignToConfiguration("JENKINS_MASTER_DEFAULT_LABEL", testJenkinsHelmMasterDefaultLabel)

	models.AssignToConfiguration("JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS", testJenkinsHelmMasterDenyAnonymousReadAccess)
	models.AssignToConfiguration("JENKINS_MASTER_ADMIN_PASSWORD", testJenkinsHelmMasterAdminPassword)
	models.AssignToConfiguration("JENKINS_JCASC_CONFIGURATION_URL", testJenkinsHelmMasterJcascConfigUrl)
	models.AssignToConfiguration("JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED", testJenkinsHelmMasterAdminPasswordEncrypted)
	models.AssignToConfiguration("JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED", testJenkinsHelmMasterUserPasswordEncrypted)

	models.AssignToConfiguration("KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID", testJcascDockerCredentialsId)
	models.AssignToConfiguration("MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID", testJcascMavenCredentialsId)
	models.AssignToConfiguration("NPM_REPOSITORY_SECRETS_CREDENTIALS_ID", testJcascNpmCredentialsId)
	models.AssignToConfiguration("VCS_REPOSITORY_SECRETS_CREDENTIALS_ID", testJcascVcsCredentialsId)

	models.AssignToConfiguration("KUBERNETES_SERVER_CERTIFICATE", testJcascKubernetesCertificate)

	if setupTestProject {
		var err = createprojectactions.ActionCreateNewProjectDirectory(testProjectName)
		assert.Nil(t, err)
		err = createprojectactions.ActionCopyTemplatesToNewDirectory(testProjectName, true, false)
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

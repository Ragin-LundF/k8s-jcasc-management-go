package project

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/actions/createprojectactions"
	"k8s-management-go/app/models"
	"testing"
)

// ----- Constants for testing (base configuration)
const testRootDirectory = "../../../"
const testProjectName = testRootDirectory + "projects/__k8s_project_test__"

const testNamespace = "my-namespace"
const testPvcNamespace = "my-namespace"
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

func TestProjectTemplates(t *testing.T) {
	testDefaultProjectConfiguration(t, true)
	var project = NewProject("my-namespace")
	project.PersistentVolumeClaim.SetMetadataName("my-pvc-claim")

	var err = project.ProcessTemplates(testProjectName)
	assert.Nil(t, err)

	// _ = os.RemoveAll(testProjectName)
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

	if setupTestProject {
		var err = createprojectactions.ActionCreateNewProjectDirectory(testProjectName)
		assert.Nil(t, err)
		err = createprojectactions.ActionCopyTemplatesToNewDirectory(testProjectName, true, false)
		assert.Nil(t, err)
	}
}

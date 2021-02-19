package project

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/actions/createprojectactions"
	"k8s-management-go/app/models"
	"testing"
)

const testRootDirectory = "../../../"
const testProjectName = testRootDirectory + "projects/__k8s_project_test__"

const testPvcNamespace = "my-Namespace"
const testConfigJenkinsMasterDeploymentName = "jenkins-master"
const testConfigJenkinsMasterPvcAccessMode = "ReadWriteOnce"
const testConfigJenkinsMasterPvcStorageClassName = "hostpath"
const testConfigJenkinsMasterPvcSize = "1Gi"

func TestProjectTemplates(t *testing.T) {
	defaultProjectConfiguration(t)
	var project = NewProject("my-namespace")
	project.PersistentVolumeClaim.SetMetadataName("my-pvc-claim")

	var err = project.ProcessTemplates(testProjectName)
	assert.Nil(t, err)

	// _ = os.RemoveAll(testProjectName)
}

func defaultProjectConfiguration(t *testing.T) {
	models.AssignToConfiguration("TEMPLATES_BASE_DIRECTORY", testRootDirectory+"templates/")

	models.AssignToConfiguration("JENKINS_MASTER_DEPLOYMENT_NAME", testConfigJenkinsMasterDeploymentName)
	models.AssignToConfiguration("JENKINS_MASTER_DEFAULT_URI_PREFIX", "/jenkins")

	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_ACCESS_MODE", testConfigJenkinsMasterPvcAccessMode)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE", testConfigJenkinsMasterPvcSize)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS", testConfigJenkinsMasterPvcStorageClassName)

	models.AssignToConfiguration("NGINX_INGRESS_ANNOTATION_CLASS", "nginx")
	models.AssignToConfiguration("NGINX_INGRESS_DEPLOYMENT_NAME", "nginx-ingress")

	models.AssignToConfiguration("NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE", "bitnami/nginx-ingress-controller:latest")
	models.AssignToConfiguration("NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS", "mypullsecret")
	models.AssignToConfiguration("NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE", "true")

	models.AssignToConfiguration("NGINX_LOADBALANCER_ENABLED", "true")
	models.AssignToConfiguration("NGINX_LOADBALANCER_HTTP_PORT", "80")
	models.AssignToConfiguration("NGINX_LOADBALANCER_HTTP_TARGETPORT", "8080")
	models.AssignToConfiguration("NGINX_LOADBALANCER_HTTPS_PORT", "443")
	models.AssignToConfiguration("NGINX_LOADBALANCER_HTTPS_TARGETPORT", "8443")

	models.AssignToConfiguration("NGINX_LOADBALANCER_ANNOTATIONS_ENABLED", "true")
	models.AssignToConfiguration("NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME", "domain.tld")
	models.AssignToConfiguration("NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_TTL", "3600")

	var err = createprojectactions.ActionCreateNewProjectDirectory(testProjectName)
	assert.Nil(t, err)
	err = createprojectactions.ActionCopyTemplatesToNewDirectory(testProjectName, true, false)
	assert.Nil(t, err)
}

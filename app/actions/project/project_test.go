package project

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/models"
	"testing"
)

func TestProjectTemplates(t *testing.T) {
	defaultPprojectConfiguration()

	var project = Project{
		PersistentVolumeClaim: NewPersistentVolumeClaim(pvcNamespace, pvcName),
		Namespace:             struct{ Name string }{Name: "Test"},
	}
	var err = project.ProcessTemplates("../../../templates/")
	assert.Nil(t, err)
}

func defaultPprojectConfiguration() {
	models.AssignToConfiguration("JENKINS_MASTER_DEPLOYMENT_NAME", configJenkinsMasterDeploymentName)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_ACCESS_MODE", configJenkinsMasterPvcAccessMode)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE", configJenkinsMasterPvcSize)
	models.AssignToConfiguration("JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS", configJenkinsMasterPvcStorageClassName)
}

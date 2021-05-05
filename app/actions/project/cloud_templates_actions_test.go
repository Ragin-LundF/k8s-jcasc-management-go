package project

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/configuration"
	"strings"
	"testing"
)

func TestActionReadCloudTemplates(t *testing.T) {
	createTestConfiguration()
	configuration.GetConfiguration().K8SManagement.Project.TemplateDirectory = "../../../templates"
	var templates = ActionReadCloudTemplates()

	assert.True(t, len(templates) == 2)
	assert.Equal(t, templates[0], "gradle_java.yaml")
	assert.Equal(t, templates[1], "node.yaml")
}

func TestActionReadCloudTemplatesAsString(t *testing.T) {
	createTestConfiguration()
	configuration.GetConfiguration().K8SManagement.Project.TemplateDirectory = "../../../templates"

	var result, err = ActionReadCloudTemplatesAsString([]string{"gradle_java.yaml"})

	assert.Nil(t, err)
	assert.True(t, len(result) > 100)
	assert.True(t, strings.Contains(result, "name: \"gradle_java\""))
}

package namespaceactions

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/models"
	"testing"
)

func init() {
	models.ResetIPAndNamespaces()
	models.AddIPAndNamespaceToConfiguration("namespace-a", "1.2.3.4")
	models.AddIPAndNamespaceToConfiguration("stage-alpha", "1.2.3.5")
	models.AddIPAndNamespaceToConfiguration("projectA", "1.2.3.6")
	models.AddIPAndNamespaceToConfiguration("projectB", "1.2.3.7")
	models.AddIPAndNamespaceToConfiguration("beta-stage", "1.2.3.8")
	models.AddIPAndNamespaceToConfiguration("production-stage", "1.2.3.9")
}

func TestActionReadNamespaceWithFilterNil(t *testing.T) {
	var namespaces = ActionReadNamespaceWithFilter(nil)

	assert.Len(t, namespaces, 6)
}

func TestActionReadNamespaceWithFilterValue(t *testing.T) {
	var filter = "project"
	var namespaces = ActionReadNamespaceWithFilter(&filter)

	assert.Len(t, namespaces, 2)
}

func TestActionReadNamespaceWithFilterValueNotExisting(t *testing.T) {
	var filter = "notexisting"
	var namespaces = ActionReadNamespaceWithFilter(&filter)

	assert.Len(t, namespaces, 0)
}

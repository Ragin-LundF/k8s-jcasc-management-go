package namespaceactions

import (
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
	if len(namespaces) == 6 {
		t.Log("Success. Found all entries")
	} else {
		t.Errorf("Failed. Number of returned items with nil filter was wrong [%v]", len(namespaces))
	}
}

func TestActionReadNamespaceWithFilterValue(t *testing.T) {
	var filter = "project"
	var namespaces = ActionReadNamespaceWithFilter(&filter)
	if len(namespaces) == 2 {
		t.Log("Success. Found all filtered entries")
	} else {
		t.Errorf("Failed. Number of returned items with project filter was wrong [%v]", len(namespaces))
	}
}

func TestActionReadNamespaceWithFilterValueNotExisting(t *testing.T) {
	var filter = "notexisting"
	var namespaces = ActionReadNamespaceWithFilter(&filter)
	if len(namespaces) == 0 {
		t.Log("Success. Found no entries")
	} else {
		t.Errorf("Failed. Number of returned items with project filter was wrong [%v]", len(namespaces))
	}
}

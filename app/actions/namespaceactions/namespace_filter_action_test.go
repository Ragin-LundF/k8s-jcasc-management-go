package namespaceactions

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/configuration"
	"testing"
)

func init() {
	configuration.LoadConfiguration("../../../", false, false)
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = nil

	var newIpAndNamespace = configuration.DeploymentStruct{
		IPAddress: "1.2.3.4",
		Namespace: "namespace-a",
	}
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = append(configuration.GetConfiguration().K8SManagement.IPConfig.Deployments, newIpAndNamespace)
	newIpAndNamespace = configuration.DeploymentStruct{
		IPAddress: "1.2.3.5",
		Namespace: "stage-alpha",
	}
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = append(configuration.GetConfiguration().K8SManagement.IPConfig.Deployments, newIpAndNamespace)
	newIpAndNamespace = configuration.DeploymentStruct{
		IPAddress: "1.2.3.6",
		Namespace: "projectA",
	}
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = append(configuration.GetConfiguration().K8SManagement.IPConfig.Deployments, newIpAndNamespace)
	newIpAndNamespace = configuration.DeploymentStruct{
		IPAddress: "1.2.3.7",
		Namespace: "projectB",
	}
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = append(configuration.GetConfiguration().K8SManagement.IPConfig.Deployments, newIpAndNamespace)
	newIpAndNamespace = configuration.DeploymentStruct{
		IPAddress: "1.2.3.8",
		Namespace: "beta-stage",
	}
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = append(configuration.GetConfiguration().K8SManagement.IPConfig.Deployments, newIpAndNamespace)
	newIpAndNamespace = configuration.DeploymentStruct{
		IPAddress: "1.2.3.9",
		Namespace: "production-stage",
	}
	configuration.GetConfiguration().K8SManagement.IPConfig.Deployments = append(configuration.GetConfiguration().K8SManagement.IPConfig.Deployments, newIpAndNamespace)
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

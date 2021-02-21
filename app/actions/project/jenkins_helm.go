package project

import "k8s-management-go/app/models"

// ----- Structures
// jenkinsHelm : Model which describes the jenkins helm values
type jenkinsHelmValues struct {
	Master      jenkinsHelmMaster
	Persistence jenkinsHelmPersistence
}

// jenkinsHelmMaster : Model which describes the Jenkins master section in the helm values
type jenkinsHelmMaster struct {
	Image                                        string
	Tag                                          string
	ImagePullPolicy                              string
	ImagePullSecretName                          string
	CustomJenkinsLabels                          string
	AuthorizationStrategyDenyAnonymousReadAccess string
	AdminPassword                                string
	SidecarsConfigAutoReloadFolder               string
}

// jenkinsHelmPersistence : Model which describes the persistence section in the helm values
type jenkinsHelmPersistence struct {
	StorageClass string
	AccessMode   string
	Size         string
}

// NewJenkinsHelmValues : Create new Jenkins Helm values structure
func NewJenkinsHelmValues() *jenkinsHelmValues {
	return &jenkinsHelmValues{
		Master:      newDefaultJenkinsHelmMaster(),
		Persistence: newDefaultJenkinsHelmPersistence(),
	}
}

// ----- internal methods
// newDefaultJenkinsHelmMaster : create a new default jenkinsHelmMaster structure
func newDefaultJenkinsHelmMaster() jenkinsHelmMaster {
	var configuration = models.GetConfiguration()
	return jenkinsHelmMaster{
		Image:                          configuration.Jenkins.Helm.Master.Container.Image,
		Tag:                            configuration.Jenkins.Helm.Master.Container.ImageTag,
		ImagePullPolicy:                configuration.Jenkins.Helm.Master.Container.PullPolicy,
		ImagePullSecretName:            configuration.Jenkins.Helm.Master.Container.PullSecretName,
		CustomJenkinsLabels:            configuration.Jenkins.Helm.Master.Label,
		AdminPassword:                  configuration.Jenkins.Helm.Master.AdminPassword,
		SidecarsConfigAutoReloadFolder: configuration.Jenkins.JCasC.ConfigurationURL,
		AuthorizationStrategyDenyAnonymousReadAccess: configuration.Jenkins.Helm.Master.DenyAnonymousReadAccess,
	}
}

func newDefaultJenkinsHelmPersistence() jenkinsHelmPersistence {
	var configuration = models.GetConfiguration()
	return jenkinsHelmPersistence{
		StorageClass: configuration.Jenkins.Helm.Master.Persistence.StorageClass,
		AccessMode:   configuration.Jenkins.Helm.Master.Persistence.AccessMode,
		Size:         configuration.Jenkins.Helm.Master.Persistence.Size,
	}
}

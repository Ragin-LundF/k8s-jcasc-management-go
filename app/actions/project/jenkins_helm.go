package project

import (
	"k8s-management-go/app/configuration"
	"strconv"
)

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
	return jenkinsHelmMaster{
		Image:                          configuration.GetConfiguration().Jenkins.Container.Image,
		Tag:                            configuration.GetConfiguration().Jenkins.Container.Tag,
		ImagePullPolicy:                configuration.GetConfiguration().Jenkins.Container.PullPolicy,
		ImagePullSecretName:            configuration.GetConfiguration().Jenkins.Container.PullSecret,
		CustomJenkinsLabels:            configuration.GetConfiguration().Jenkins.Controller.CustomJenkinsLabel,
		AdminPassword:                  configuration.GetConfiguration().Jenkins.Controller.Passwords.AdminUser,
		SidecarsConfigAutoReloadFolder: configuration.GetConfiguration().Jenkins.Jcasc.ConfigurationURL,
		AuthorizationStrategyDenyAnonymousReadAccess: strconv.FormatBool(!configuration.GetConfiguration().Jenkins.Jcasc.AuthorizationStrategy.AllowAnonymousRead),
	}
}

func newDefaultJenkinsHelmPersistence() jenkinsHelmPersistence {
	return jenkinsHelmPersistence{
		StorageClass: configuration.GetConfiguration().Jenkins.Persistence.StorageClass,
		AccessMode:   configuration.GetConfiguration().Jenkins.Persistence.AccessMode,
		Size:         configuration.GetConfiguration().Jenkins.Persistence.StorageSize,
	}
}

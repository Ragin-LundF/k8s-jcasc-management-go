package project

import (
	"k8s-management-go/app/configuration"
	"strconv"
)

// ----- Structures
// jenkinsHelm : Model which describes the jenkins helm values
type jenkinsHelmValues struct {
	Controller           jenkinsHelmMaster      `yaml:"controller,omitempty"`
	Persistence          jenkinsHelmPersistence `yaml:"persistence,omitempty"`
	AdditionalNamespaces []string               `yaml:"additionalNamespaces,omitempty"`
}

// jenkinsHelmMaster : Model which describes the Jenkins master section in the helm values
type jenkinsHelmMaster struct {
	Image                                   string `yaml:"image,omitempty"`
	Tag                                     string `yaml:"tag,omitempty"`
	ImagePullPolicy                         string `yaml:"imagePullPolicy,omitempty"`
	ImagePullSecretName                     string `yaml:"imagePullSecretName,omitempty"`
	CustomJenkinsLabels                     string `yaml:"customJenkinsLabel,omitempty"`
	AuthorizationStrategyAllowAnonymousRead string `yaml:"authorizationStrategyAllowAnonymousRead,omitempty"`
	AdminPassword                           string `yaml:"adminPassword,omitempty"`
	SidecarsConfigAutoReloadFolder          string `yaml:"sidecarsConfigAutoReloadFolder,omitempty"`
}

// jenkinsHelmPersistence : Model which describes the persistence section in the helm values
type jenkinsHelmPersistence struct {
	AccessMode   string `yaml:"accessMode,omitempty"`
	Size         string `yaml:"storageSize,omitempty"`
	StorageClass string `yaml:"storageClass,omitempty"`
}

// NewJenkinsHelmValues : Create new Jenkins Helm values structure
func NewJenkinsHelmValues() *jenkinsHelmValues {
	return &jenkinsHelmValues{
		Controller:           NewDefaultJenkinsHelmController(),
		Persistence:          NewDefaultJenkinsHelmPersistence(),
		AdditionalNamespaces: NewDefaultAdditionalNamespaces(),
	}
}

// NewDefaultJenkinsHelmController : create a new default jenkinsHelmMaster structure
func NewDefaultJenkinsHelmController() jenkinsHelmMaster {
	return jenkinsHelmMaster{
		Image:                                   configuration.GetConfiguration().Jenkins.Container.Image,
		Tag:                                     configuration.GetConfiguration().Jenkins.Container.Tag,
		ImagePullPolicy:                         configuration.GetConfiguration().Jenkins.Container.PullPolicy,
		ImagePullSecretName:                     configuration.GetConfiguration().Jenkins.Container.PullSecret,
		CustomJenkinsLabels:                     configuration.GetConfiguration().Jenkins.Controller.CustomJenkinsLabel,
		AdminPassword:                           configuration.GetConfiguration().Jenkins.Controller.Passwords.AdminUser,
		SidecarsConfigAutoReloadFolder:          configuration.GetConfiguration().Jenkins.Jcasc.ConfigurationURL,
		AuthorizationStrategyAllowAnonymousRead: strconv.FormatBool(configuration.GetConfiguration().Jenkins.Jcasc.AuthorizationStrategy.AllowAnonymousRead),
	}
}

func NewDefaultJenkinsHelmPersistence() jenkinsHelmPersistence {
	return jenkinsHelmPersistence{
		StorageClass: configuration.GetConfiguration().Jenkins.Persistence.StorageClass,
		AccessMode:   configuration.GetConfiguration().Jenkins.Persistence.AccessMode,
		Size:         configuration.GetConfiguration().Jenkins.Persistence.StorageSize,
	}
}

func NewDefaultAdditionalNamespaces() []string {
	return []string{}
}

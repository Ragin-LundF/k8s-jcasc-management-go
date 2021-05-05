package project

import (
	"k8s-management-go/app/actions/kubernetesactions"
	"k8s-management-go/app/configuration"
	"strings"
)

// ----- Structures
// jcascConfig : Model which describes the JcasC (Jenkins configuration as code) config helm values
type jcascConfig struct {
	Clouds        *clouds        `yaml:"clouds,omitempty"`
	CredentialIDs *credentialIDs `yaml:"credentialIDs,omitempty"`
	JobsConfig    *jobsConfig    `yaml:"jobsConfig,omitempty"`
	SecurityRealm *securityRealm `yaml:"securityRealm,omitempty"`
	SystemMessage string         `yaml:"systemMessage,omitempty"`
}

// clouds : Model which describes the Clouds section in the helm values
type clouds struct {
	Kubernetes kubernetes `yaml:"kubernetes,omitempty"`
}

// credentialIDs : Model which describes the common Kubernetes settings
type credentialIDs struct {
	DockerRegistryCredentialsID         string `yaml:"docker,omitempty"`
	MavenRepositorySecretsCredentialsID string `yaml:"maven,omitempty"`
	NpmRepositorySecretsCredentialsID   string `yaml:"npm,omitempty"`
	VcsRepositoryCredentialsID          string `yaml:"vcs,omitempty"`
}

// jobsConfig : Model which describes the jobs configuration
type jobsConfig struct {
	JobsSeedRepository       string `yaml:"seedJobRepository,omitempty"`
	JobsDefinitionRepository string `yaml:"jobsDefinitionRepository,omitempty"`
}

// securityRealm : Model which describes the security realm section in the helm values
type securityRealm struct {
	LocalUsers *securityRealmLocalUsers `yaml:"localUsers,omitempty"`
}

// securityRealmLocalUsers : Model which describes the security realm local users section in the helm values
type securityRealmLocalUsers struct {
	AdminPassword string `yaml:"adminPassword,omitempty"`
	UserPassword  string `yaml:"userPassword,omitempty"`
}

// kubernetes : Model which describes the Kubernetes section in the helm values
type kubernetes struct {
	Templates kubernetesTemplates `yaml:"templates,omitempty"`
}

// kubernetesTemplates : Model which describes the Kubernetes Templates section in the helm values
type kubernetesTemplates struct {
	AdditionalCloudTemplateFiles []string `yaml:"additionalCloudTemplateFiles"`
	AdditionalCloudTemplates     string   `yaml:"-"`
}

// ----- Setter to manipulate the default object

// SetJenkinsSystemMessage : Set the Jenkins system message
func (jcascCfg *jcascConfig) SetJenkinsSystemMessage(jenkinsSystemMessage string) {
	jcascCfg.SystemMessage = jenkinsSystemMessage
}

// SetAdminPassword : Set admin password to local security realm user
func (jcascCfg *jcascConfig) SetAdminPassword(adminPassword string) {
	jcascCfg.SecurityRealm.LocalUsers.AdminPassword = adminPassword
}

// SetUserPassword : Set user password to local security realm user
func (jcascCfg *jcascConfig) SetUserPassword(userPassword string) {
	jcascCfg.SecurityRealm.LocalUsers.UserPassword = userPassword
}

// SetCloudKubernetesAdditionalTemplates : Set additional templates for cloud.kubernetes.templates
func (jcascCfg *jcascConfig) SetCloudKubernetesAdditionalTemplates(additionalTemplates string) {
	jcascCfg.Clouds.Kubernetes.Templates.AdditionalCloudTemplates = additionalTemplates
}

// SetCloudKubernetesAdditionalTemplateFiles : Set additional templates for cloud.kubernetes.templates
func (jcascCfg *jcascConfig) SetCloudKubernetesAdditionalTemplateFiles(additionalTemplateFiles []string) {
	jcascCfg.Clouds.Kubernetes.Templates.AdditionalCloudTemplateFiles = additionalTemplateFiles
}

// SetJobsSeedRepository : Set seed jobs repository for jobs configuration
func (jcascCfg *jcascConfig) SetJobsSeedRepository(seedRepository string) {
	jcascCfg.JobsConfig.JobsSeedRepository = seedRepository
}

// SetJobsDefinitionRepository : Set jobs repository for jobs configuration
func (jcascCfg *jcascConfig) SetJobsDefinitionRepository(jobsRepository string) {
	jcascCfg.JobsConfig.JobsDefinitionRepository = jobsRepository
}

// JobsAvailable : method to check if jobs are available. Can be used in the templates to disable the jobs section
func (jobsCfg *jobsConfig) JobsAvailable() bool {
	if jobsCfg.JobsDefinitionRepository != "" && jobsCfg.JobsSeedRepository != "" {
		return true
	}
	return false
}

// ServerCertificate : Get the server certificate for the current context
func (k8s *kubernetes) ServerCertificate() string {
	var currentContext = strings.ToUpper(kubernetesactions.GetKubernetesConfig().CurrentContext())
	if configuration.GetConfiguration().Kubernetes.Certificates.Contexts != nil {
		for context, certificate := range configuration.GetConfiguration().Kubernetes.Certificates.Contexts {
			if strings.EqualFold(currentContext, context) {
				return certificate
			}
		}
	}

	return configuration.GetConfiguration().Kubernetes.Certificates.Default
}

// ----- internal methods

// newJCascConfig : Create new Jenkins Helm values structure
func newJCascConfig() *jcascConfig {
	return &jcascConfig{
		CredentialIDs: newCredentialIDs(),
		Clouds:        newClouds(),
		JobsConfig: &jobsConfig{
			JobsSeedRepository:       configuration.GetConfiguration().Jenkins.Jcasc.SeedJobURL,
			JobsDefinitionRepository: "",
		},
		SecurityRealm: &securityRealm{
			LocalUsers: newSecurityRealmLocalUsers(),
		},
		SystemMessage: "",
	}
}

// newCredentialIDs : create new default credential IDs
func newCredentialIDs() *credentialIDs {
	return &credentialIDs{
		DockerRegistryCredentialsID:         configuration.GetConfiguration().Jenkins.Jcasc.CredentialIDs.Docker,
		MavenRepositorySecretsCredentialsID: configuration.GetConfiguration().Jenkins.Jcasc.CredentialIDs.Maven,
		NpmRepositorySecretsCredentialsID:   configuration.GetConfiguration().Jenkins.Jcasc.CredentialIDs.Npm,
		VcsRepositoryCredentialsID:          configuration.GetConfiguration().Jenkins.Jcasc.CredentialIDs.Vcs,
	}
}

// newClouds : create new default clouds
func newClouds() *clouds {
	return &clouds{
		Kubernetes: *newCloudKubernetes(),
	}
}

// newCloudKubernetes : create new default newCloudKubernetes
func newCloudKubernetes() *kubernetes {
	return &kubernetes{
		Templates: *newCloudKubernetesSubTemplates(),
	}
}

// newCloudKubernetesSubTemplates : create new default sub-templates for cloud.kubernetes.templates
func newCloudKubernetesSubTemplates() *kubernetesTemplates {
	return &kubernetesTemplates{
		AdditionalCloudTemplates: "",
	}
}

// newSecurityRealmLocalUsers : create a new default securityRealmLocalUsers structure
func newSecurityRealmLocalUsers() *securityRealmLocalUsers {
	return &securityRealmLocalUsers{
		AdminPassword: configuration.GetConfiguration().Jenkins.Controller.Passwords.AdminUserEncrypted,
		UserPassword:  configuration.GetConfiguration().Jenkins.Controller.Passwords.DefaultUserEncrypted,
	}
}

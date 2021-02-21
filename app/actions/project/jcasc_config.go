package project

import "k8s-management-go/app/models"

// ----- Structures
// jcascConfig : Model which describes the JcasC (Jenkins configuration as code) config helm values
type jcascConfig struct {
	Clouds        clouds
	CredentialIDs credentialIDs
	JobsConfig    jobsConfig
	SecurityRealm securityRealm
	SystemMessage string
}

// clouds : Model which describes the Clouds section in the helm values
type clouds struct {
	Kubernetes kubernetes
}

// credentialIDs : Model which describes the common Kubernetes settings
type credentialIDs struct {
	DockerRegistryCredentialsID         string
	MavenRepositorySecretsCredentialsID string
	NpmRepositorySecretsCredentialsID   string
	VcsRepositoryCredentialsID          string
}

// jobsConfig : Model which describes the jobs configuration
type jobsConfig struct {
	JobsSeedRepository       string
	JobsDefinitionRepository string
}

// securityRealm : Model which describes the security realm section in the helm values
type securityRealm struct {
	LocalUsers securityRealmLocalUsers
}

// securityRealmLocalUsers : Model which describes the security realm local users section in the helm values
type securityRealmLocalUsers struct {
	AdminPassword string
	UserPassword  string
}

// kubernetes : Model which describes the Kubernetes section in the helm values
type kubernetes struct {
	ServerCertificate string
	Templates         kubernetesTemplates
}

// kubernetesTemplates : Model which describes the Kubernetes Templates section in the helm values
type kubernetesTemplates struct {
	AdditionalCloudTemplates string
}

// NewJCascConfig : Create new Jenkins Helm values structure
func NewJCascConfig() *jcascConfig {
	return &jcascConfig{
		CredentialIDs: newCredentialIDs(),
		Clouds:        newClouds(),
		JobsConfig: jobsConfig{
			JobsSeedRepository:       "",
			JobsDefinitionRepository: "",
		},
		SecurityRealm: securityRealm{
			LocalUsers: newSecurityRealmLocalUsers(),
		},
		SystemMessage: "",
	}
}

// ----- Setter to manipulate the default object
// SetJenkinsSystemMessage : Set the Jenkins system message
func (jcascConfig *jcascConfig) SetJenkinsSystemMessage(jenkinsSystemMessage string) {
	jcascConfig.SystemMessage = jenkinsSystemMessage
}

// SetAdminPassword : Set admin password to local security realm user
func (jcascConfig *jcascConfig) SetAdminPassword(adminPassword string) {
	jcascConfig.SecurityRealm.LocalUsers.AdminPassword = adminPassword
}

// SetUserPassword : Set user password to local security realm user
func (jcascConfig *jcascConfig) SetUserPassword(userPassword string) {
	jcascConfig.SecurityRealm.LocalUsers.UserPassword = userPassword
}

// SetCloudKubernetesAdditionalTemplates : Set additional templates for cloud.kubernetes.templates
func (jcascConfig *jcascConfig) SetCloudKubernetesAdditionalTemplates(additionalTemplates string) {
	jcascConfig.Clouds.Kubernetes.Templates.AdditionalCloudTemplates = additionalTemplates
}

// SetJobsSeedRepository : Set seed jobs repository for jobs configuration
func (jcascConfig *jcascConfig) SetJobsSeedRepository(seedRepository string) {
	jcascConfig.JobsConfig.JobsSeedRepository = seedRepository
}

// SetJobsDefinitionRepository : Set jobs repository for jobs configuration
func (jcascConfig *jcascConfig) SetJobsDefinitionRepository(jobsRepository string) {
	jcascConfig.JobsConfig.JobsDefinitionRepository = jobsRepository
}

// JobsAvailable : method to check if jobs are available. Can be used in the templates to disable the jobs section
func (jobsConfig *jobsConfig) JobsAvailable() bool {
	if jobsConfig.JobsDefinitionRepository != "" && jobsConfig.JobsSeedRepository != "" {
		return true
	}
	return false
}

// ----- internal methods

// newCredentialIDs : create new default credential IDs
func newCredentialIDs() credentialIDs {
	var configuration = models.GetConfiguration()
	return credentialIDs{
		DockerRegistryCredentialsID:         configuration.CredentialIds.DefaultDockerRegistry,
		MavenRepositorySecretsCredentialsID: configuration.CredentialIds.DefaultMavenRepository,
		NpmRepositorySecretsCredentialsID:   configuration.CredentialIds.DefaultNpmRepository,
		VcsRepositoryCredentialsID:          configuration.CredentialIds.DefaultVcsRepository,
	}
}

// newClouds : create new default clouds
func newClouds() clouds {
	return clouds{
		Kubernetes: newCloudKubernetes(),
	}
}

// newCloudKubernetes : create new default newCloudKubernetes
func newCloudKubernetes() kubernetes {
	var configuration = models.GetConfiguration()
	return kubernetes{
		ServerCertificate: configuration.Kubernetes.ServerCertificate,
		Templates:         newCloudKubernetesSubTemplates(),
	}
}

// newCloudKubernetesSubTemplates : create new default sub-templates for cloud.kubernetes.templates
func newCloudKubernetesSubTemplates() kubernetesTemplates {
	return kubernetesTemplates{
		AdditionalCloudTemplates: "",
	}
}

// newSecurityRealmLocalUsers : create a new default securityRealmLocalUsers structure
func newSecurityRealmLocalUsers() securityRealmLocalUsers {
	var configuration = models.GetConfiguration()
	return securityRealmLocalUsers{
		AdminPassword: configuration.Jenkins.Helm.Master.AdminPasswordEncrypted,
		UserPassword:  configuration.Jenkins.Helm.Master.DefaultProjectUserPasswordEncrypted,
	}
}

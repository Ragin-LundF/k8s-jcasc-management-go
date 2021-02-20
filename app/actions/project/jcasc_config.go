package project

import "k8s-management-go/app/models"

// ----- Structures
// jcascConfig : Model which describes the JcasC (Jenkins configuration as code) config helm values
type jcascConfig struct {
	CredentialIDs credentialIDs
	Clouds        clouds
	SystemMessage string
	SecurityRealm securityRealm
}

// credentialIDs : Structure for common Kubernetes settings
type credentialIDs struct {
	DockerRegistryCredentialsID         string
	MavenRepositorySecretsCredentialsID string
	NpmRepositorySecretsCredentialsId   string
	VcsRepositoryCredentialsID          string
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

// clouds : Model which describes the clouds section in the helm values
type clouds struct {
	ServerCertificate        string
	AdditionalCloudTemplates string
}

// NewJCascConfig : Create new Jenkins Helm values structure
func NewJCascConfig() *jcascConfig {
	return &jcascConfig{
		CredentialIDs: newCredentialIDs(),
		Clouds:        newClouds(),
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

// ----- internal methods

// newCredentialIDs : create new default credential IDs
func newCredentialIDs() credentialIDs {
	var configuration = models.GetConfiguration()
	return credentialIDs{
		DockerRegistryCredentialsID:         configuration.CredentialIds.DefaultDockerRegistry,
		MavenRepositorySecretsCredentialsID: configuration.CredentialIds.DefaultMavenRepository,
		NpmRepositorySecretsCredentialsId:   configuration.CredentialIds.DefaultNpmRepository,
		VcsRepositoryCredentialsID:          configuration.CredentialIds.DefaultVcsRepository,
	}
}

// newClouds : create new default clouds
func newClouds() clouds {
	var configuration = models.GetConfiguration()
	return clouds{
		ServerCertificate: configuration.Kubernetes.ServerCertificate,
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

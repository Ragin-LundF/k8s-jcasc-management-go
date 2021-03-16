package project

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/actions/kubernetesactions"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/cmdexecutor"
	"strings"
	"testing"
)

func TestCreateJCascConfig(t *testing.T) {
	testDefaultProjectConfiguration(t, false)
	cmdexecutor.Executor = TestCommandExec{}
	var jcascConfig = newJCascConfig()

	assert.Equal(t, testJcascKubernetesCertificate, jcascConfig.Clouds.Kubernetes.ServerCertificate())
	assert.Empty(t, jcascConfig.Clouds.Kubernetes.Templates.AdditionalCloudTemplates)

	assert.Empty(t, jcascConfig.SystemMessage)

	assert.False(t, jcascConfig.JobsConfig.JobsAvailable())
	assert.Equal(t, testJenkinsHelmMasterJcascConfigSeedUrl, jcascConfig.JobsConfig.JobsSeedRepository)
	assert.Empty(t, jcascConfig.JobsConfig.JobsDefinitionRepository)

	assert.Equal(t, testJenkinsHelmMasterAdminPasswordEncrypted, jcascConfig.SecurityRealm.LocalUsers.AdminPassword)
	assert.Equal(t, testJenkinsHelmMasterUserPasswordEncrypted, jcascConfig.SecurityRealm.LocalUsers.UserPassword)

	assert.Equal(t, testJcascDockerCredentialsId, jcascConfig.CredentialIDs.DockerRegistryCredentialsID)
	assert.Equal(t, testJcascMavenCredentialsId, jcascConfig.CredentialIDs.MavenRepositorySecretsCredentialsID)
	assert.Equal(t, testJcascNpmCredentialsId, jcascConfig.CredentialIDs.NpmRepositorySecretsCredentialsID)
	assert.Equal(t, testJcascVcsCredentialsId, jcascConfig.CredentialIDs.VcsRepositoryCredentialsID)
}

func TestCreateJCascConfigCloudsContextCertificate(t *testing.T) {
	// set custom context
	testDefaultProjectConfiguration(t, false)

	var customCertificate = testJcascKubernetesCertificate + "-custom"
	cmdexecutor.Executor = TestCommandExecCustomContext{}
	kubernetesactions.ReloadKubernetesContext()

	var cfg = configuration.GetConfiguration()
	cfg.Kubernetes.Certificates.Contexts = map[string]string{
		"custom-k8s": customCertificate,
	}

	var jcascConfig = newJCascConfig()

	assert.Equal(t, customCertificate, jcascConfig.Clouds.Kubernetes.ServerCertificate())
	assert.Empty(t, jcascConfig.Clouds.Kubernetes.Templates.AdditionalCloudTemplates)

	// reset executor and context
	cmdexecutor.Executor = TestCommandExec{}
	kubernetesactions.ReloadKubernetesContext()
}

func TestCreateJCascConfigSystemMessage(t *testing.T) {
	testDefaultProjectConfiguration(t, false)
	cmdexecutor.Executor = TestCommandExec{}
	var jcascConfig = newJCascConfig()

	assert.Empty(t, jcascConfig.SystemMessage)
}

func TestCreateJCascConfigJobsConfig(t *testing.T) {
	var jobsRepository = "jobs.git"
	var jobsSeedRepository = "seed.git"

	testDefaultProjectConfiguration(t, false)
	cmdexecutor.Executor = TestCommandExec{}
	var jcascConfig = newJCascConfig()
	jcascConfig.SetJobsDefinitionRepository(jobsRepository)
	jcascConfig.SetJobsSeedRepository(jobsSeedRepository)

	assert.True(t, jcascConfig.JobsConfig.JobsAvailable())
	assert.Equal(t, jobsRepository, jcascConfig.JobsConfig.JobsDefinitionRepository)
	assert.Equal(t, jobsSeedRepository, jcascConfig.JobsConfig.JobsSeedRepository)
}

//NOSONAR
func TestCreateJCascConfigSecurityRealmPasswords(t *testing.T) {
	var adminPassword = "new_enc_pass_admin" //NOSONAR
	var userPassword = "new_enc_pass_user"   //NOSONAR

	testDefaultProjectConfiguration(t, false)
	cmdexecutor.Executor = TestCommandExec{}
	var jcascConfig = newJCascConfig()
	jcascConfig.SetAdminPassword(adminPassword)
	jcascConfig.SetUserPassword(userPassword)

	assert.Equal(t, adminPassword, jcascConfig.SecurityRealm.LocalUsers.AdminPassword)
	assert.Equal(t, userPassword, jcascConfig.SecurityRealm.LocalUsers.UserPassword)
}

// TestCommandExecCustomContext is the test executor for mocks
type TestCommandExecCustomContext struct{}

// CombinedOutput is the mock implementation of CombinedOutput
func (c TestCommandExecCustomContext) CombinedOutput(command string, args ...string) ([]byte, error) {
	var result []byte
	var commandAsString = command + " " + strings.Join(args, " ")
	result = combinedOutputCurrentCustomContext(args)
	if result != nil {
		return result, nil
	}
	result = combinedOutputGetContexts(args)
	if result != nil {
		return result, nil
	}

	return []byte(commandAsString + "...executed"), nil
}

// combinedOutputCurrentCustomContext returns the custom kubernetes config current-context
func combinedOutputCurrentCustomContext(args []string) []byte {
	if cap(args) == 2 {
		if args[0] == "config" && args[1] == "current-context" {
			return []byte("custom-k8s")
		}
	}
	return nil
}

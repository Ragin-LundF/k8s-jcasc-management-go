package configuration

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Init() config {
	return config{}
}

func TestGetAlternativeSecretsFilesEmpty(t *testing.T) {
	var conf = config{}
	conf.K8SManagement.Project.BaseDirectory = "./"
	conf.K8SManagement.Project.SecretFiles = "./secrets.sh"

	assert.True(t, len(conf.GetSecretsFiles()) == 1)
}

func TestGetAlternativeSecretsFilesWithDifferentBasePath(t *testing.T) {
	var conf = config{}
	conf.K8SManagement.Project.BaseDirectory = "../"
	conf.K8SManagement.Project.SecretFiles = "./secrets.sh"

	_ = ioutil.WriteFile("../secrets.sh", []byte(""), 0644)
	var secretsFile = conf.getGlobalSecretsFile()
	var secretsFileA = strings.Replace(secretsFile, "secrets.sh", "secrets_environment_a.sh", -1)
	var secretsFileB = strings.Replace(secretsFile, "secrets.sh", "secrets_environment_b.sh", -1)
	_ = ioutil.WriteFile(secretsFile, []byte(""), 0644)
	_ = ioutil.WriteFile(secretsFileA, []byte(""), 0644)
	_ = ioutil.WriteFile(secretsFileB, []byte(""), 0644)

	var alternativeSecretFiles = conf.GetSecretsFiles()
	assert.NotEmpty(t, alternativeSecretFiles)
	assert.True(t, len(alternativeSecretFiles) == 3)
	for _, secretFile := range alternativeSecretFiles {
		runeSecretFile := []rune(secretFile)
		assert.Equal(t, "secrets", string(runeSecretFile[0:7]))
	}

	_ = os.Remove(secretsFileA)
	_ = os.Remove(secretsFileB)
	_ = os.Remove("../secrets.sh")
}

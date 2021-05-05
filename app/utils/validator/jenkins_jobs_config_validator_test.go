package validator

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/configuration"
	"testing"
)

func TestValidateJenkinsJobConfig(t *testing.T) {
	assignPattern()
	var jenkinsJobConfig = "https://github.com/repo.git"

	err := ValidateJenkinsJobConfig(jenkinsJobConfig)

	assert.NoError(t, err)
}

func TestValidateJenkinsJobConfigWithoutGitRepo(t *testing.T) {
	assignPattern()
	var jenkinsJobConfig = "https://github.com/repo"

	err := ValidateJenkinsJobConfig(jenkinsJobConfig)

	assert.Error(t, err)
}

func assignPattern() {
	configuration.LoadConfiguration("../../../", false, false)
	configuration.GetConfiguration().Jenkins.JobDSL.RepoValidatePattern = ".*\\.git"
}

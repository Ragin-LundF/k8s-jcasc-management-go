package validator

import (
	"github.com/stretchr/testify/assert"
	"k8s-management-go/app/models"
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
	models.AssignToConfiguration("JENKINS_JOBDSL_REPO_VALIDATE_PATTERN", ".*\\.git")
}

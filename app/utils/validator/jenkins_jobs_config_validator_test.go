package validator

import (
	"k8s-management-go/app/models"
	"testing"
)

func TestValidateJenkinsJobConfig(t *testing.T) {
	assignPattern()
	var jenkinsJobConfig = "https://github.com/repo.git"

	err := ValidateJenkinsJobConfig(jenkinsJobConfig)
	if err != nil {
		t.Error("Failed. Validator returned error.")
	} else {
		t.Log("Success. Validator successfully accepted repository.")
	}
}

func TestValidateJenkinsJobConfigWithoutGitRepo(t *testing.T) {
	assignPattern()
	var jenkinsJobConfig = "https://github.com/repo"

	err := ValidateJenkinsJobConfig(jenkinsJobConfig)
	if err != nil {
		t.Log("Success. Validator successfully returned an error.")
	} else {
		t.Error("Failed. Validator accepted the repo without extensions.")
	}
}

func assignPattern() {
	models.AssignToConfiguration("JENKINS_JOBDSL_REPO_VALIDATE_PATTERN", ".*\\.git")
}

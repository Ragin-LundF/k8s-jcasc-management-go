package validator

import (
	"errors"
	"k8s-management-go/app/models"
	"regexp"
)

func ValidateJenkinsJobConfig(input string) error {
	// Job repository should not be longer than 512 characters
	if len(input) > 512 {
		return errors.New("Should not be longer than 512 characters. ")
	}
	// Regex regex to validate repository
	regex := regexp.MustCompile(models.GetConfiguration().Jenkins.JobDSL.RepoValidatePattern)
	if !regex.Match([]byte(input)) {
		return errors.New("Wrong repository name! ")
	}

	return nil
}

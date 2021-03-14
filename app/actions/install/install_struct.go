package install

import (
	"k8s-management-go/app/actions/project"
)

// ProjectConfig : describes the project config for install/uninstall actions
type ProjectConfig struct {
	Project                project.Project
	ProjectPath            string
	JenkinsHelmValuesFile  string
	JenkinsHelmValuesExist bool
	NginxHelmValuesExist   bool
	SecretsFileName        string
	SecretsPassword        *string
	HelmCommand            string
}

// PvcClaimValuesYaml is a structure that represents the PVC Claim values yaml file
type PvcClaimValuesYaml struct {
	Kind       string
	APIVersion string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string
	}
	Spec struct {
		AccessModes      []string
		StorageClassName string
		Resources        struct {
			Requests struct {
				Storage string
			}
		}
	}
}

// NewInstallProjectConfig returns a new install of the ProjectConfig
func NewInstallProjectConfig() ProjectConfig {
	return ProjectConfig{
		Project:                project.NewProject(),
		ProjectPath:            "",
		JenkinsHelmValuesFile:  "",
		JenkinsHelmValuesExist: false,
		NginxHelmValuesExist:   false,
		SecretsFileName:        "",
		SecretsPassword:        nil,
		HelmCommand:            "",
	}
}

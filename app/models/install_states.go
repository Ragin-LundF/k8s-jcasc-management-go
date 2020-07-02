package models

// StateData defines the data for project creation
type StateData struct {
	ProjectPath            string
	Namespace              string
	DeploymentName         string
	JenkinsHelmValuesFile  string
	JenkinsHelmValuesExist bool
	NginxHelmValuesExist   bool
	SecretsPassword        *string
	HelmCommand            string
}

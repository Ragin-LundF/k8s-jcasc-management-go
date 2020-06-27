package models

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

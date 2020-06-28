package models

type ProjectConfig struct {
	Namespace                   string
	IpAddress                   string
	JenkinsSystemMsg            string
	JobsCfgRepo                 string
	ExistingPvc                 string
	SelectedCloudTemplates      []string
	CreateDeploymentOnlyProject bool
}

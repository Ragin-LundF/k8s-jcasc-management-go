package models

type ProjectConfig struct {
	Namespace                   string
	IpAddress                   string
	JenkinsSystemMsg            string
	JobsCfgRepo                 string
	SelectedCloudTemplates      []string
	ExistingPvc                 string
	CreateDeploymentOnlyProject bool
}

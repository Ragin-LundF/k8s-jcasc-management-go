package models

// ProjectConfig defines the values that are necessary for the project configuration
type ProjectConfig struct {
	Namespace                   string
	IPAddress                   string
	JenkinsDomain               string
	JenkinsSystemMsg            string
	JobsCfgRepo                 string
	SelectedCloudTemplates      []string
	ExistingPvc                 string
	CreateDeploymentOnlyProject bool
}

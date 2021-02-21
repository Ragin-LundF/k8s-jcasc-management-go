package project

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"text/template"
)

// Project : Structure for a complete project
type Project struct {
	Base                  *base
	JCasc                 *jcascConfig
	JenkinsHelmValues     *jenkinsHelmValues
	PersistentVolumeClaim *persistentVolumeClaim
	Nginx                 *nginx
}

// base : represents common Jenkins settings
type base struct {
	DeploymentName      string
	Domain              string
	ExistingVolumeClaim string
	JenkinsUriPrefix    string
	IPAddress           string
	Namespace           string
}

// NewProject : create a new Project
func NewProject() Project {
	return Project{
		Base:                  newBase(),
		JenkinsHelmValues:     NewJenkinsHelmValues(),
		JCasc:                 NewJCascConfig(),
		PersistentVolumeClaim: NewPersistentVolumeClaim(),
		Nginx:                 NewNginx(),
	}
}

// JenkinsURL : If load balancer annotations are enabled it returns a domain. Else the IP address.
func (base *base) JenkinsURL() string {
	var configuration = models.GetConfiguration()

	if configuration.LoadBalancer.Annotations.Enabled {
		if base.Domain == "" {
			return fmt.Sprintf("%v.%v", base.Namespace, configuration.LoadBalancer.Annotations.ExtDNS.Hostname)
		} else {
			return base.Domain
		}
	} else {
		return base.IPAddress
	}
}

// ----- Setter to manipulate the default object
// SetIPAddress : Set the Jenkins IP address
func (project *Project) SetIPAddress(ipAddress string) {
	project.Base.IPAddress = ipAddress
}

// SetNamespace : Set the Jenkins namespace
func (project *Project) SetNamespace(namespace string) {
	project.Base.Namespace = namespace
}

// SetNamespace : Set the Jenkins domain
func (project *Project) SetDomain(domain string) {
	project.Base.Domain = domain
}

// SetJenkinsSystemMessage : Set the Jenkins system message
func (project *Project) SetJenkinsSystemMessage(jenkinsSystemMessage string) {
	project.JCasc.SetJenkinsSystemMessage(jenkinsSystemMessage)
}

// SetAdminPassword : Set admin password to local security realm user
func (project *Project) SetAdminPassword(adminPassword string) {
	project.JCasc.SetAdminPassword(adminPassword)
}

// SetUserPassword : Set user password to local security realm user
func (project *Project) SetUserPassword(userPassword string) {
	project.JCasc.SetUserPassword(userPassword)
}

// SetCloudKubernetesAdditionalTemplates : Set additional templates for cloud.kubernetes.templates
func (project *Project) SetCloudKubernetesAdditionalTemplates(additionalTemplates string) {
	project.JCasc.SetCloudKubernetesAdditionalTemplates(additionalTemplates)
}

// SetJobsSeedRepository : Set seed jobs repository for jobs configuration
func (project *Project) SetJobsSeedRepository(seedRepository string) {
	project.JCasc.SetJobsSeedRepository(seedRepository)
}

// SetJobsDefinitionRepository : Set jobs repository for jobs configuration
func (project *Project) SetJobsDefinitionRepository(jobsRepository string) {
	project.JCasc.SetJobsDefinitionRepository(jobsRepository)
}

// SetPersistentVolumeClaimExistingName : Set an existing PVC
func (project *Project) SetPersistentVolumeClaimExistingName(existingPvcName string) {
	project.Base.ExistingVolumeClaim = existingPvcName
}

// ProcessTemplates : Interface implementation to process templates with PVC placeholder
func (project *Project) ProcessTemplates(projectDirectory string) (err error) {
	err = project.validateProject()
	if err != nil {
		return err
	}

	templateFiles, err := files.LoadTemplateFilesOfDirectory(projectDirectory)
	if err != nil {
		return err
	}

	for _, templateFile := range templateFiles {
		err = processWithTemplateEngine(templateFile, *project)
		if err != nil {
			_ = os.Remove(projectDirectory)
			loggingstate.AddErrorEntryAndDetails(
				fmt.Sprintf("-> Unable to process file [%v] with template engine", templateFile),
				err.Error())
			return err
		}
	}

	return nil
}

// processWithTemplateEngine : Process files with template engine
func processWithTemplateEngine(filename string, project Project) (err error) {
	// replace JCasC URL
	var jcascUrl bytes.Buffer
	jcascUrlTemplate, err := template.New("JcasC-URL").Parse(project.JenkinsHelmValues.Master.SidecarsConfigAutoReloadFolder)
	if err != nil {
		return err
	}

	err = jcascUrlTemplate.Execute(&jcascUrl, project)
	if err != nil {
		return err
	}
	project.JenkinsHelmValues.Master.SidecarsConfigAutoReloadFolder = jcascUrl.String()

	// replace sub-templates
	var jcascCloudsKubernetesSubTemplates bytes.Buffer
	jcascCloudsKuberentesTemplate, err := template.New("JcasC-CloudKuberetesTemplates").Parse(project.JCasc.Clouds.Kubernetes.Templates.AdditionalCloudTemplates)
	if err != nil {
		return err
	}

	err = jcascCloudsKuberentesTemplate.Execute(&jcascCloudsKubernetesSubTemplates, project)
	if err != nil {
		return err
	}
	project.JCasc.SetCloudKubernetesAdditionalTemplates(jcascCloudsKubernetesSubTemplates.String())

	// replace templates
	templates, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}

	var processedTemplate bytes.Buffer
	err = templates.Execute(&processedTemplate, project)
	if err != nil {
		return err
	}

	if processedTemplate.Len() > 0 {
		_ = ioutil.WriteFile(filename, processedTemplate.Bytes(), 0)
	}

	return nil
}

// newBase : Base Jenkins setup
func newBase() *base {
	var configuration = models.GetConfiguration()

	return &base{
		DeploymentName:      configuration.Jenkins.Helm.Master.DeploymentName,
		Domain:              "",
		ExistingVolumeClaim: "",
		JenkinsUriPrefix:    configuration.Jenkins.Helm.Master.DefaultURIPrefix,
		IPAddress:           "",
		Namespace:           "",
	}
}

// validateProject : Validate the project that it can be processed
func (project *Project) validateProject() (err error) {
	var configuration = models.GetConfiguration()
	if project.Base.Namespace == "" {
		return errors.New("Error: No namespace available ")
	}

	if !configuration.LoadBalancer.Annotations.Enabled && project.Base.IPAddress == "" {
		return errors.New("Error: If NGINX_LOADBALANCER_ANNOTATIONS_ENABLED is set to false, an IP address is required ")
	}

	return nil
}

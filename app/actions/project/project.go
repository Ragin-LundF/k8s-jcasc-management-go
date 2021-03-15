package project

import (
	"bytes"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"text/template"
)

// Project : Structure for a complete project
type Project struct {
	Base                  *base                  `yaml:"base,omitempty"`
	JCasc                 *jcascConfig           `yaml:"jcasc,omitempty"`
	JenkinsHelmValues     *jenkinsHelmValues     `yaml:"jenkinsHelmValues,omitempty"`
	PersistentVolumeClaim *persistentVolumeClaim `yaml:"persistentVolumeClaim,omitempty"`
	Nginx                 *nginx                 `yaml:"nginx,omitempty"`
}

// base : represents common Jenkins settings
type base struct {
	DeploymentName      string `yaml:"deploymentName,omitempty"`
	Domain              string `yaml:"domain,omitempty"`
	ExistingVolumeClaim string `yaml:"existingVolumeClaim,omitempty"`
	JenkinsUriPrefix    string `yaml:"jenkinsURIPrefix,omitempty"`
	IPAddress           string `yaml:"ipAddress,omitempty"`
	Namespace           string `yaml:"namespace,omitempty"`
	DeploymentOnly      bool   `yaml:"deploymentOnly,omitempty"`
}

// NewProject : create a new Project
func NewProject() Project {
	return Project{
		Base:                  newBase(),
		JenkinsHelmValues:     newJenkinsHelmValues(),
		JCasc:                 newJCascConfig(),
		PersistentVolumeClaim: newPersistentVolumeClaim(),
		Nginx:                 newNginx(),
	}
}

// JenkinsURL : If load balancer annotations are enabled it returns a domain. Else the IP address.
func (bse base) JenkinsURL() string {
	if configuration.GetConfiguration().Nginx.Loadbalancer.Annotations.Enabled {
		if bse.Domain == "" {
			return fmt.Sprintf("%v.%v", bse.Namespace, configuration.GetConfiguration().Nginx.Loadbalancer.ExternalDNS.HostName)
		} else {
			return bse.Domain
		}
	} else {
		return bse.IPAddress
	}
}

// ----- Setter to manipulate the default object
// SetIPAddress : Set the Jenkins IP address
func (prj *Project) SetIPAddress(ipAddress string) {
	prj.Base.IPAddress = ipAddress
}

// SetNamespace : Set the Jenkins namespace
func (prj *Project) SetNamespace(namespace string) {
	prj.Base.Namespace = namespace
}

// SetNamespace : Set the Jenkins domain
func (prj *Project) SetDomain(domain string) {
	prj.Base.Domain = domain
}

// SetJenkinsSystemMessage : Set the Jenkins system message
func (prj *Project) SetJenkinsSystemMessage(jenkinsSystemMessage string) {
	prj.JCasc.SetJenkinsSystemMessage(jenkinsSystemMessage)
}

// SetAdminPassword : Set admin password to local security realm user
func (prj *Project) SetAdminPassword(adminPassword string) {
	prj.JCasc.SetAdminPassword(adminPassword)
}

// SetUserPassword : Set user password to local security realm user
func (prj *Project) SetUserPassword(userPassword string) {
	prj.JCasc.SetUserPassword(userPassword)
}

// SetCloudKubernetesAdditionalTemplates : Set additional templates for cloud.kubernetes.templates
func (prj *Project) SetCloudKubernetesAdditionalTemplates(additionalTemplates string) {
	prj.JCasc.SetCloudKubernetesAdditionalTemplates(additionalTemplates)
}

// SetCloudKubernetesAdditionalTemplateFiles : Set additional template files for cloud.kubernetes.templates
func (prj *Project) SetCloudKubernetesAdditionalTemplateFiles(additionalTemplateFiles []string) {
	prj.JCasc.SetCloudKubernetesAdditionalTemplateFiles(additionalTemplateFiles)
}

// SetJobsSeedRepository : Set seed jobs repository for jobs configuration
func (prj *Project) SetJobsSeedRepository(seedRepository string) {
	prj.JCasc.SetJobsSeedRepository(seedRepository)
}

// SetJobsDefinitionRepository : Set jobs repository for jobs configuration
func (prj *Project) SetJobsDefinitionRepository(jobsRepository string) {
	prj.JCasc.SetJobsDefinitionRepository(jobsRepository)
}

// SetPersistentVolumeClaimExistingName : Set an existing PVC
func (prj *Project) SetPersistentVolumeClaimExistingName(existingPvcName string) {
	prj.Base.ExistingVolumeClaim = existingPvcName
}

// SaveProjectConfiguration: Save the project configuration
func (prj *Project) SaveProjectConfiguration(projectDirectory string) (err error) {
	marshalledOutput, err := yaml.Marshal(prj)
	if err != nil {
		return err
	}

	var projectConfigPath = files.AppendPath(projectDirectory, constants.FilenameProjectConfiguration)
	_ = ioutil.WriteFile(projectConfigPath, marshalledOutput, 0644)

	return nil
}

// ProcessTemplates : Interface implementation to process templates with PVC placeholder
func (prj *Project) ProcessTemplates(projectDirectory string) (err error) {
	err = prj.validateProject()
	if err != nil {
		return err
	}

	templateFiles, err := files.LoadTemplateFilesOfDirectory(projectDirectory)
	if err != nil {
		return err
	}

	for _, templateFile := range templateFiles {
		err = processWithTemplateEngine(templateFile, *prj)
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

// CalculateRequiredDeploymentFiles : calculate which project files are required
func (prj *Project) CalculateRequiredDeploymentFiles() []string {
	var deploymentFiles []string
	// ingress controller
	deploymentFiles = append(deploymentFiles, constants.FilenameNginxIngressControllerHelmValues)
	// if it is not a deployment only project, copy more files
	if !prj.Base.DeploymentOnly {
		deploymentFiles = append(deploymentFiles, constants.FilenameJenkinsHelmValues)
		// copy Jenkins values.yaml
		deploymentFiles = append(deploymentFiles, constants.FilenameJenkinsHelmValues)
		// copy Jenkins JCasC config.yaml
		deploymentFiles = append(deploymentFiles, constants.FilenameJenkinsConfigurationAsCode)
		// copy existing PVC values.yaml
		if len(prj.Base.ExistingVolumeClaim) > 0 {
			deploymentFiles = append(deploymentFiles, constants.FilenamePvcClaim)
		}
		// copy secrets to project
		if configuration.GetConfiguration().K8SManagement.Project.SecretFiles == "" {
			deploymentFiles = append(deploymentFiles, constants.FilenameSecrets)
		}
	}
	return deploymentFiles
}

// processWithTemplateEngine : Process files with template engine
func processWithTemplateEngine(filename string, project Project) (err error) {
	// replace JCasC URL
	var jcascUrl bytes.Buffer
	jcascUrlTemplate, err := template.New("JcasC-URL").Parse(project.JenkinsHelmValues.Controller.SidecarsConfigAutoReloadFolder)
	if err != nil {
		return err
	}

	err = jcascUrlTemplate.Execute(&jcascUrl, project)
	if err != nil {
		return err
	}
	project.JenkinsHelmValues.Controller.SidecarsConfigAutoReloadFolder = jcascUrl.String()

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
	return &base{
		DeploymentName:      configuration.GetConfiguration().Jenkins.Controller.DeploymentName,
		Domain:              "",
		ExistingVolumeClaim: "",
		JenkinsUriPrefix:    configuration.GetConfiguration().Jenkins.Controller.DefaultURIPrefix,
		IPAddress:           "",
		Namespace:           "",
	}
}

// validateProject : Validate the project that it can be processed
func (prj *Project) validateProject() (err error) {
	if prj.Base.Namespace == "" {
		return errors.New("Error: No namespace available ")
	}

	var enabledAnnotations = configuration.GetConfiguration().Nginx.Loadbalancer.Annotations.Enabled
	if !enabledAnnotations && prj.Base.IPAddress == "" {
		return errors.New("Error: If nginx.loadbalancer.annotations.enabled is set to false, an IP address is required ")
	}

	return nil
}

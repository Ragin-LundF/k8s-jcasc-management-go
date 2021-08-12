package install

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s-management-go/app/actions/project"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
)

// ProjectConfig : describes the project config for install/uninstall actions
type ProjectConfig struct {
	Project         project.Project
	ProjectPath     string
	SecretsFileName string
	SecretsPassword *string
	HelmCommand     string
	ConfigLoaded    bool
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
	var prjConfig = ProjectConfig{
		Project:         project.NewProject(),
		ProjectPath:     "",
		SecretsFileName: "",
		SecretsPassword: nil,
		HelmCommand:     "",
		ConfigLoaded:    false,
	}
	prjConfig.Project.StoreConfigOnly = false
	return prjConfig
}

// LoadProjectConfigIfExists : Load a configuration if exists.
func (projectConfig *ProjectConfig) LoadProjectConfigIfExists(namespace string) (err error) {
	projectConfig.Project.Base.Namespace = namespace
	if err = projectConfig.CalculateDirectoriesForInstall(); err != nil {
		return err
	}

	var configFile = files.AppendPath(projectConfig.ProjectPath, constants.FilenameProjectConfiguration)
	if files.FileOrDirectoryExists(configFile) {
		yamlFile, err := ioutil.ReadFile(configFile)
		if err != nil {
			return err
		}

		// parse YAML
		var prj project.Project
		err = yaml.Unmarshal(yamlFile, &prj)
		if err != nil {
			return err
		}

		// config loaded successfully, assign it
		projectConfig.Project = prj
		projectConfig.ConfigLoaded = true
	} else {
		configureProjectWithoutConfig(projectConfig)
	}

	return nil
}

func configureProjectWithoutConfig(projectConfig *ProjectConfig) {
	// Jenkins
	if !files.FileOrDirectoryExists(files.AppendPath(projectConfig.ProjectPath, constants.FilenameJenkinsHelmValues)) {
		projectConfig.Project.Base.DeploymentOnly = true
	}
}

// PrepareInstallYAML : Prepare temporary YAML if required or return path to project file
func (projectConfig *ProjectConfig) PrepareInstallYAML(filename string) (fileWithPath string, err error) {
	fileWithPath = files.AppendPath(projectConfig.ProjectPath, filename)

	// if config was loaded process with config
	if projectConfig.ConfigLoaded && projectConfig.Project.CalculateIfDeploymentFileIsRequired(filename) {
		// check if file not exists to create a temporary file
		if !files.FileOrDirectoryExists(fileWithPath) {
			// filename with temp prefix
			if err = project.CopyTemplate(projectConfig.ProjectPath, filename, true); err != nil {
				return fileWithPath, err
			}
			fileWithPath = calculateTempFilename(projectConfig.ProjectPath, filename)

			// process template with config
			err = projectConfig.Project.ProcessWithTemplateEngine(fileWithPath)
		}
	}

	return fileWithPath, err
}

func calculateTempFilename(projectDir string, filename string) string {
	return files.AppendPath(projectDir, constants.FilenameTempPrefix+filename)
}

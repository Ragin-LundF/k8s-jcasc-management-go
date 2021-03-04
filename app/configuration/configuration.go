package configuration

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"log"
)

// config represents the configuration files
type config struct {
	K8SManagement struct {
		Log struct {
			Level              string `yaml:"level"`
			File               string `yaml:"file"`
			Encoding           string `yaml:"encoding"`
			OverwriteOnRestart bool   `yaml:"overwriteOnRestart"`
		} `yaml:"log"`
		Ipconfig struct {
			File        string `yaml:"file"`
			DummyPrefix string `yaml:"dummyPrefix"`
		} `yaml:"ipconfig"`
		Project struct {
			BaseDirectory     string `yaml:"baseDirectory"`
			TemplateDirectory string `yaml:"templateDirectory"`
			SecretFiles       string `yaml:"secretFiles"`
		} `yaml:"project"`
		VersionCheck bool `yaml:"versionCheck"`
		DryRunDebug  bool
		CliOnly      bool
	} `yaml:"k8sManagement"`
	Jenkins struct {
		Jcasc struct {
			ConfigurationURL      string `yaml:"configurationUrl"`
			AuthorizationStrategy struct {
				AllowAnonymousRead bool `yaml:"allowAnonymousRead"`
			} `yaml:"authorizationStrategy"`
			CredentialIDs struct {
				Docker string `yaml:"docker"`
				Maven  string `yaml:"maven"`
				Npm    string `yaml:"npm"`
				Vcs    string `yaml:"vcs"`
			} `yaml:"credentialIDs"`
			SeedJobURL string `yaml:"seedJobURL"`
		} `yaml:"jcasc"`
		JobDSL struct {
			BaseURL             string `yaml:"baseURL"`
			RepoValidatePattern string `yaml:"repoValidatePattern"`
		} `yaml:"jobDSL"`
		Controller struct {
			Passwords struct {
				AdminUser            string `yaml:"adminUser"`
				AdminUserEncrypted   string `yaml:"adminUserEncrypted"`
				DefaultUserEncrypted string `yaml:"defaultUserEncrypted"`
			} `yaml:"passwords"`
			CustomJenkinsLabel string `yaml:"customJenkinsLabel"`
			DeploymentName     string `yaml:"deploymentName"`
		} `yaml:"controller"`
		Persistence struct {
			AccessMode   string `yaml:"accessMode"`
			StorageClass string `yaml:"storageClass"`
			StorageSize  string `yaml:"storageSize"`
		} `yaml:"persistence"`
		Container struct {
			Image      string `yaml:"image"`
			Tag        string `yaml:"tag"`
			PullPolicy string `yaml:"pullPolicy"`
			PullSecret string `yaml:"pullSecret"`
		} `yaml:"container"`
	} `yaml:"jenkins"`
	Nginx struct {
		Ingress struct {
			Container struct {
				Image      string `yaml:"image"`
				PullSecret string `yaml:"pullSecret"`
			} `yaml:"container"`
			Deployment struct {
				ForEachNamespace bool   `yaml:"forEachNamespace"`
				DeploymentName   string `yaml:"deploymentName"`
			} `yaml:"deployment"`
			Annotationclass string `yaml:"annotationclass"`
		} `yaml:"ingress"`
		Loadbalancer struct {
			Enabled bool `yaml:"enabled"`
			Ports   struct {
				HTTP        int `yaml:"http"`
				HTTPTarget  int `yaml:"httpTarget"`
				HTTPS       int `yaml:"https"`
				HTTPSTarget int `yaml:"httpsTarget"`
			} `yaml:"ports"`
			Annotations struct {
				Enabled bool `yaml:"enabled"`
			} `yaml:"annotations"`
			ExternalDNS struct {
				HostName string `yaml:"hostName"`
				TTL      int    `yaml:"ttl"`
			} `yaml:"externalDNS"`
		} `yaml:"loadbalancer"`
	} `yaml:"nginx"`
	Kubernetes struct {
		Certificates struct {
			Default  string      `yaml:"default"`
			Contexts interface{} `yaml:"contexts"`
		} `yaml:"certificates"`
	} `yaml:"kubernetes"`
}

var conf *config

// GetConfiguration returns the current configuration
func GetConfiguration() *config {
	return conf
}

func LoadConfiguration(basePath string, dryRunDebug bool, cliOnly bool) {
	conf = &config{}
	conf.initBaseConfig(basePath)
	conf.K8SManagement.DryRunDebug = dryRunDebug
	conf.K8SManagement.CliOnly = cliOnly
}

func (conf *config) initBaseConfig(basePath string) {
	yamlFile, err := ioutil.ReadFile(files.AppendPath(files.AppendPath(basePath, constants.DirConfig), constants.FilenameConfigurationYaml))
	if err != nil {
		log.Printf("Unable to read base config: #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unable to unmarshal: %v", err)
	}
}

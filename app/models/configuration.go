package models

import (
	"k8s-management-go/app/utils/files"
	"strconv"
	"strings"
)

var configuration Configuration

type Configuration struct {
	BasePath string
	// Log Level
	LogLevel string
	// secrets file
	GlobalSecretsFile string
	// Alternative ConfigFile
	AlternativeConfigFile string
	// IP config
	IpConfig struct {
		IpConfigFile            string
		IpConfigFileDummyPrefix string
	}
	// Directory configuration
	Directories struct {
		ProjectsBaseDirectory  string
		TemplatesBaseDirectory string
	}
	// Jenkins configuration
	Jenkins struct {
		// JCasC relevant data
		JCasC struct {
			ConfigurationUrl string
		}
		// JobDSL relevant data
		JobDSL struct {
			BaseUrl             string
			RepoValidatePattern string
			SeedJobScriptUrl    string
		}
		// Jenkins Helm Chart relevant data
		Helm struct {
			Master struct {
				AdminPassword                       string
				AdminPasswordEncrypted              string
				DefaultProjectUserPasswordEncrypted string
				Label                               string
				DenyAnonymousReadAccess             string
				DefaultUriPrefix                    string
				DeploymentName                      string
				Persistence                         struct {
					StorageClass string
					AccessMode   string
					Size         string
				}
				Container struct {
					Image          string
					ImageTag       string
					PullPolicy     string
					PullSecretName string
				}
			}
		}
	}
	// Nginx relevant data
	Nginx struct {
		Ingress struct {
			AnnotationClass string
			Controller      struct {
				DeploymentName string
				Container      struct {
					Name       string
					PullSecret string
					Namespace  bool
				}
			}
		}
	}
	// Loadbalancer relevant data
	LoadBalancer struct {
		Enabled bool
		Port    struct {
			Http        uint64
			HttpTarget  uint64
			Https       uint64
			HttpsTarget uint64
		}
	}
	// Kubernetes relevant data
	Kubernetes struct {
		ServerCertificate string
	}
	// Default credential ids
	CredentialIds struct {
		DefaultDockerRegistry  string
		DefaultMavenRepository string
		DefaultNpmRepository   string
		DefaultVcsRepository   string
	}
	// internal configuration
	K8sManagement struct {
		VersionCheck bool
		DryRunOnly   bool
	}
}

func GetConfiguration() Configuration {
	return configuration
}

// helper method for IP configuration file
func GetIpConfigurationFile() string {
	return FilePathWithBasePath(configuration.IpConfig.IpConfigFile)
}

// helper method for secrets file
func GetGlobalSecretsFile() string {
	return FilePathWithBasePath(configuration.GlobalSecretsFile)
}

// Get project base directory with full path
func GetProjectBaseDirectory() string {
	return calculateFullDirectoryPath(configuration.Directories.ProjectsBaseDirectory)
}

// Get project template directory with full path
func GetProjectTemplateDirectory() string {
	return calculateFullDirectoryPath(configuration.Directories.TemplatesBaseDirectory)
}

// calculate full directory path
func calculateFullDirectoryPath(targetDir string) string {
	if strings.HasPrefix(targetDir, "./") {
		// if it starts with current directory add base path
		return FilePathWithBasePath(targetDir)
	} else if strings.HasPrefix(targetDir, "../") {
		// if it starts with subdirectory add base path
		return FilePathWithBasePath(targetDir)
	} else {
		// it seems to be an absolute path, use only the project directory
		return targetDir
	}
}

// helper method to calculate the correct filepath
func FilePathWithBasePath(configurationFilePath string) string {
	if configuration.BasePath != "" {
		configurationFilePath = files.AppendPath(configuration.BasePath, configurationFilePath)
	}
	return configurationFilePath
}

func AssignToConfiguration(key string, value string) {
	if key != "" && value != "" {
		switch key {
		case "LOG_LEVEL":
			configuration.LogLevel = value
		case "GLOBAL_SECRETS_FILE":
			configuration.GlobalSecretsFile = value
		case "IP_CONFIG_FILE":
			configuration.IpConfig.IpConfigFile = value
		case "IP_CONFIG_FILE_DUMMY_PREFIX":
			configuration.IpConfig.IpConfigFileDummyPrefix = value
		case "PROJECTS_BASE_DIRECTORY":
			configuration.Directories.ProjectsBaseDirectory = value
		case "TEMPLATES_BASE_DIRECTORY":
			configuration.Directories.TemplatesBaseDirectory = value
		case "JENKINS_JCASC_CONFIGURATION_URL":
			configuration.Jenkins.JCasC.ConfigurationUrl = value
		case "JENKINS_JOBDSL_BASE_URL":
			configuration.Jenkins.JobDSL.BaseUrl = value
		case "JENKINS_JOBDSL_REPO_VALIDATE_PATTERN":
			configuration.Jenkins.JobDSL.RepoValidatePattern = value
		case "JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL":
			configuration.Jenkins.JobDSL.SeedJobScriptUrl = value
		case "JENKINS_MASTER_ADMIN_PASSWORD":
			configuration.Jenkins.Helm.Master.AdminPassword = value
		case "JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED":
			configuration.Jenkins.Helm.Master.AdminPasswordEncrypted = value
		case "JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED":
			configuration.Jenkins.Helm.Master.DefaultProjectUserPasswordEncrypted = value
		case "JENKINS_MASTER_DEFAULT_LABEL":
			configuration.Jenkins.Helm.Master.Label = value
		case "JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS":
			configuration.Jenkins.Helm.Master.DenyAnonymousReadAccess = value
		case "JENKINS_MASTER_DEFAULT_URI_PREFIX":
			configuration.Jenkins.Helm.Master.DefaultUriPrefix = value
		case "JENKINS_MASTER_DEPLOYMENT_NAME":
			configuration.Jenkins.Helm.Master.DeploymentName = value
		case "JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS":
			configuration.Jenkins.Helm.Master.Persistence.StorageClass = value
		case "JENKINS_MASTER_PERSISTENCE_ACCESS_MODE":
			configuration.Jenkins.Helm.Master.Persistence.AccessMode = value
		case "JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE":
			configuration.Jenkins.Helm.Master.Persistence.Size = value
		case "JENKINS_MASTER_CONTAINER_IMAGE":
			configuration.Jenkins.Helm.Master.Container.Image = value
		case "JENKINS_MASTER_CONTAINER_IMAGE_TAG":
			configuration.Jenkins.Helm.Master.Container.ImageTag = value
		case "JENKINS_MASTER_CONTAINER_PULL_POLICY":
			configuration.Jenkins.Helm.Master.Container.PullPolicy = value
		case "JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME":
			configuration.Jenkins.Helm.Master.Container.PullSecretName = value
		case "NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE":
			configuration.Nginx.Ingress.Controller.Container.Name = value
		case "NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS":
			configuration.Nginx.Ingress.Controller.Container.PullSecret = value
		case "NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE":
			configuration.Nginx.Ingress.Controller.Container.Namespace, _ = strconv.ParseBool(value)
		case "NGINX_INGRESS_DEPLOYMENT_NAME":
			configuration.Nginx.Ingress.Controller.DeploymentName = value
		case "NGINX_INGRESS_ANNOTATION_CLASS":
			configuration.Nginx.Ingress.AnnotationClass = value
		case "NGINX_LOADBALANCER_ENABLED":
			configuration.LoadBalancer.Enabled, _ = strconv.ParseBool(value)
		case "NGINX_LOADBALANCER_HTTP_PORT":
			configuration.LoadBalancer.Port.Http, _ = strconv.ParseUint(value, 10, 16)
		case "NGINX_LOADBALANCER_HTTP_TARGETPORT":
			configuration.LoadBalancer.Port.HttpTarget, _ = strconv.ParseUint(value, 10, 16)
		case "NGINX_LOADBALANCER_HTTPS_PORT":
			configuration.LoadBalancer.Port.Https, _ = strconv.ParseUint(value, 10, 16)
		case "NGINX_LOADBALANCER_HTTPS_TARGETPORT":
			configuration.LoadBalancer.Port.HttpsTarget, _ = strconv.ParseUint(value, 10, 16)
		case "KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID":
			configuration.CredentialIds.DefaultDockerRegistry = value
		case "MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID":
			configuration.CredentialIds.DefaultMavenRepository = value
		case "NPM_REPOSITORY_SECRETS_CREDENTIALS_ID":
			configuration.CredentialIds.DefaultNpmRepository = value
		case "VCS_REPOSITORY_SECRETS_CREDENTIALS_ID":
			configuration.CredentialIds.DefaultVcsRepository = value
		case "K8S_MGMT_VERSION_CHECK":
			configuration.K8sManagement.VersionCheck, _ = strconv.ParseBool(value)
		case "K8S_MGMT_ALTERNATIVE_CONFIG_FILE":
			configuration.AlternativeConfigFile = value
		case "K8S_MGMT_BASE_PATH":
			configuration.BasePath = value
		case "K8S_MGMT_DRY_RUN_DEBUG":
			configuration.K8sManagement.DryRunOnly, _ = strconv.ParseBool(value)
		}
	}
}

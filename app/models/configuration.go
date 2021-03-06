package models

import (
	"strconv"
	"strings"
)

var configuration Configuration

type contextServerCertificate struct {
	Context     string
	Certificate string
}

// Configuration represents the configuration files
type Configuration struct {
	BasePath string
	// Log Level
	LogLevel string
	// Use CLI only
	CliOnly bool
	// secrets file
	GlobalSecretsFile string
	// Alternative ConfigFile
	AlternativeConfigFile string
	// IP config
	IPConfig struct {
		IPConfigFile            string
		IPConfigFileDummyPrefix string
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
			ConfigurationURL string
		}
		// JobDSL relevant data
		JobDSL struct {
			BaseURL             string
			RepoValidatePattern string
			SeedJobScriptURL    string
		}
		// Jenkins Helm Chart relevant data
		Helm struct {
			Master struct {
				AdminPassword                       string
				AdminPasswordEncrypted              string
				DefaultProjectUserPasswordEncrypted string
				Label                               string
				DenyAnonymousReadAccess             string
				DefaultURIPrefix                    string
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
			HTTP        uint64
			HTTPTarget  uint64
			HTTPS       uint64
			HTTPSTarget uint64
		}
		Annotations struct {
			Enabled bool
			ExtDNS  struct {
				Ttl      uint64
				Hostname string
			}
		}
	}
	// Kubernetes relevant data
	Kubernetes struct {
		ServerCertificate         string
		ContextServerCertificates []contextServerCertificate
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
		Logging      struct {
			LogFile             string
			LogEncoding         string
			LogOverwriteOnStart bool
		}
	}
}

// GetConfiguration returns the current configuration
func GetConfiguration() Configuration {
	return configuration
}

// AssignDryRun is a helper method for the dry-run flag
func AssignDryRun(dryRun bool) {
	configuration.K8sManagement.DryRunOnly = dryRun
}

// AssignCliOnlyMode is a helper method to assign the CLI only (cli flag) mode to the configuration
func AssignCliOnlyMode(cliOnly bool) {
	configuration.CliOnly = cliOnly
}

// NOSONAR
// AssignToConfiguration assigns a key / value pair to the configuration object
func AssignToConfiguration(key string, value string) {
	if key != "" && value != "" {
		success := addKubernetesCertificate(key, value)
		if success {
			return
		}
		success = addJenkinsJCasCJobConfig(key, value)
		if success {
			return
		}
		success = addJenkinsMasterConfig(key, value)
		if success {
			return
		}
		success = addNginxConfig(key, value)
		if success {
			return
		}
		success = addLoadBalancerConfig(key, value)
		if success {
			return
		}
		success = addCredentialsConfig(key, value)
		if success {
			return
		}
		success = addCommonConfig(key, value)
		if success {
			return
		}
		success = addK8sManagementConfig(key, value)
		if success {
			return
		}
	}
}

func addKubernetesCertificate(key string, value string) (success bool) {
	if key == "KUBERNETES_SERVER_CERTIFICATE" {
		configuration.Kubernetes.ServerCertificate = value
		success = true
	} else if strings.HasPrefix(key, "KUBERNETES_SERVER_CERTIFICATE_") {
		context := strings.TrimPrefix(key, "KUBERNETES_SERVER_CERTIFICATE_")
		ctxCertificate := contextServerCertificate{
			Context:     context,
			Certificate: value,
		}
		configuration.Kubernetes.ContextServerCertificates = append(configuration.Kubernetes.ContextServerCertificates, ctxCertificate)
	}
	return success
}

func addJenkinsJCasCJobConfig(key string, value string) (success bool) {
	switch key {
	case "JENKINS_JCASC_CONFIGURATION_URL":
		configuration.Jenkins.JCasC.ConfigurationURL = value
		break
	case "JENKINS_JOBDSL_BASE_URL":
		configuration.Jenkins.JobDSL.BaseURL = value
		break
	case "JENKINS_JOBDSL_REPO_VALIDATE_PATTERN":
		configuration.Jenkins.JobDSL.RepoValidatePattern = value
		break
	case "JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL":
		configuration.Jenkins.JobDSL.SeedJobScriptURL = value
		break
	}
	return success
}

func addJenkinsMasterConfig(key string, value string) (success bool) {
	switch key {
	case "JENKINS_MASTER_ADMIN_PASSWORD":
		configuration.Jenkins.Helm.Master.AdminPassword = value
		success = true
		break
	case "JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED":
		configuration.Jenkins.Helm.Master.AdminPasswordEncrypted = replaceUnneededChars(value)
		success = true
		break
	case "JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED":
		configuration.Jenkins.Helm.Master.DefaultProjectUserPasswordEncrypted = replaceUnneededChars(value)
		success = true
		break
	case "JENKINS_MASTER_DEFAULT_LABEL":
		configuration.Jenkins.Helm.Master.Label = value
		success = true
		break
	case "JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS":
		configuration.Jenkins.Helm.Master.DenyAnonymousReadAccess = value
		success = true
		break
	case "JENKINS_MASTER_DEFAULT_URI_PREFIX":
		configuration.Jenkins.Helm.Master.DefaultURIPrefix = value
		success = true
		break
	case "JENKINS_MASTER_DEPLOYMENT_NAME":
		configuration.Jenkins.Helm.Master.DeploymentName = value
		success = true
		break
	case "JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS":
		configuration.Jenkins.Helm.Master.Persistence.StorageClass = value
		success = true
		break
	case "JENKINS_MASTER_PERSISTENCE_ACCESS_MODE":
		configuration.Jenkins.Helm.Master.Persistence.AccessMode = value
		success = true
		break
	case "JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE":
		configuration.Jenkins.Helm.Master.Persistence.Size = value
		success = true
		break
	case "JENKINS_MASTER_CONTAINER_IMAGE":
		configuration.Jenkins.Helm.Master.Container.Image = value
		success = true
		break
	case "JENKINS_MASTER_CONTAINER_IMAGE_TAG":
		configuration.Jenkins.Helm.Master.Container.ImageTag = value
		success = true
		break
	case "JENKINS_MASTER_CONTAINER_PULL_POLICY":
		configuration.Jenkins.Helm.Master.Container.PullPolicy = value
		success = true
		break
	case "JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME":
		configuration.Jenkins.Helm.Master.Container.PullSecretName = value
		success = true
		break
	}
	return success
}

func addNginxConfig(key string, value string) (success bool) {
	switch key {
	case "NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE":
		configuration.Nginx.Ingress.Controller.Container.Name = value
		success = true
		break
	case "NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS":
		configuration.Nginx.Ingress.Controller.Container.PullSecret = value
		success = true
		break
	case "NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE":
		configuration.Nginx.Ingress.Controller.Container.Namespace, _ = strconv.ParseBool(value)
		success = true
		break
	case "NGINX_INGRESS_DEPLOYMENT_NAME":
		configuration.Nginx.Ingress.Controller.DeploymentName = value
		success = true
		break
	case "NGINX_INGRESS_ANNOTATION_CLASS":
		configuration.Nginx.Ingress.AnnotationClass = value
		success = true
		break
	}
	return success
}

func addLoadBalancerConfig(key string, value string) (success bool) {
	switch key {
	case "NGINX_LOADBALANCER_ENABLED":
		configuration.LoadBalancer.Enabled, _ = strconv.ParseBool(value)
		success = true
		break
	case "NGINX_LOADBALANCER_HTTP_PORT":
		configuration.LoadBalancer.Port.HTTP, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	case "NGINX_LOADBALANCER_HTTP_TARGETPORT":
		configuration.LoadBalancer.Port.HTTPTarget, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	case "NGINX_LOADBALANCER_HTTPS_PORT":
		configuration.LoadBalancer.Port.HTTPS, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	case "NGINX_LOADBALANCER_HTTPS_TARGETPORT":
		configuration.LoadBalancer.Port.HTTPSTarget, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	case "NGINX_LOADBALANCER_ANNOTATIONS_ENABLED":
		configuration.LoadBalancer.Annotations.Enabled, _ = strconv.ParseBool(value)
		success = true
		break
	case "NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME":
		configuration.LoadBalancer.Annotations.ExtDNS.Hostname = value
		success = true
		break
	case "NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_TTL":
		configuration.LoadBalancer.Annotations.ExtDNS.Ttl, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	}
	return success
}

func addCredentialsConfig(key string, value string) (success bool) {
	switch key {
	case "KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID":
		configuration.CredentialIds.DefaultDockerRegistry = value
		success = true
		break
	case "MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID":
		configuration.CredentialIds.DefaultMavenRepository = value
		success = true
		break
	case "NPM_REPOSITORY_SECRETS_CREDENTIALS_ID":
		configuration.CredentialIds.DefaultNpmRepository = value
		success = true
		break
	case "VCS_REPOSITORY_SECRETS_CREDENTIALS_ID":
		configuration.CredentialIds.DefaultVcsRepository = value
		success = true
		break
	}
	return success
}

func addCommonConfig(key string, value string) (success bool) {
	switch key {
	case "LOG_LEVEL":
		configuration.LogLevel = value
		success = true
		break
	case "GLOBAL_SECRETS_FILE":
		configuration.GlobalSecretsFile = value
		success = true
		break
	case "IP_CONFIG_FILE":
		configuration.IPConfig.IPConfigFile = value
		success = true
		break
	case "IP_CONFIG_FILE_DUMMY_PREFIX":
		configuration.IPConfig.IPConfigFileDummyPrefix = value
		success = true
		break
	case "PROJECTS_BASE_DIRECTORY":
		configuration.Directories.ProjectsBaseDirectory = value
		success = true
		break
	case "TEMPLATES_BASE_DIRECTORY":
		configuration.Directories.TemplatesBaseDirectory = value
		success = true
		break
	case "K8S_MGMT_ALTERNATIVE_CONFIG_FILE":
		configuration.AlternativeConfigFile = value
		success = true
		break
	case "K8S_MGMT_BASE_PATH":
		configuration.BasePath = value
		success = true
		break
	}
	return success
}

func addK8sManagementConfig(key string, value string) (success bool) {
	switch key {
	case "K8S_MGMT_LOGGING_LOGFILE":
		configuration.K8sManagement.Logging.LogFile = value
		success = true
		break
	case "K8S_MGMT_LOGGING_ENCODING":
		configuration.K8sManagement.Logging.LogEncoding = value
		success = true
		break
	case "K8S_MGMT_LOGGING_OVERWRITE_ON_START":
		configuration.K8sManagement.Logging.LogOverwriteOnStart, _ = strconv.ParseBool(value)
		success = true
		break
	case "K8S_MGMT_VERSION_CHECK":
		configuration.K8sManagement.VersionCheck, _ = strconv.ParseBool(value)
		success = true
		break
	}

	return success
}

func replaceUnneededChars(value string) string {
	if strings.Contains(value, "'") {
		value = strings.Replace(value, "'", "", -1)
	}
	if strings.Contains(value, "\"") {
		value = strings.Replace(value, "\"", "", -1)
	}

	return value
}

package migration

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"strconv"
	"strings"
)

var yamlCfg = configuration.EmptyConfiguration()

// MigrateConfigurationV3 starts configuration migration
func MigrateConfigurationV3() string {
	readConfiguration(configuration.GetConfiguration().K8SManagement.BasePath)
	var status, successful = migrateFromCnfToYaml()

	if successful {
		status = MigrateDeploymentIPConfigurationV3()
	}

	return fmt.Sprintf("Status config migration: %v", status)
}

// readConfiguration reads the configuration of k8s-management
func readConfiguration(basePath string) {
	var configFile = configuration.GetConfiguration().CustomConfig.K8SManagement.ConfigFile
	if configFile != "" && strings.Contains(configFile, ".yaml") {
		configFile = strings.Replace(configFile, ".yaml", ".cnf", -1)
	}
	if files.FileOrDirectoryExists(files.AppendPath(basePath, configFile)) {
		readConfigurationFromFile(files.AppendPath(basePath, configFile))
	}
}

// Read configuration from k8s-management config file
func readConfigurationFromFile(configFile string) {
	log := logger.Log()
	// read configuration file. Replace unneeded double quotes if needed.
	data, err := os.Open(configFile)
	// check for error
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	} else {
		// everything seems to be ok. Read data with line scanner
		scanner := bufio.NewScanner(data)
		scanner.Split(bufio.ScanLines)

		// iterate over every line
		for scanner.Scan() {
			// trim the line to avoid problems
			line := strings.TrimSpace(scanner.Text())
			processLine(line)
		}
	}
	_ = data.Close()
}

// function to check if line is empty / a comment or a valid configuration.
// if line looks valid, try to append it to internal configuration structure
func processLine(line string) {
	// if line is not a comment (marker: "#") parse the configuration and assign it to the config
	if line != "" && !strings.HasPrefix(line, "#") {
		key, value := parseConfigurationLine(line)
		assignToConfiguration(key, value)
	}
}

// parse line of configuration and split it into key/value
func parseConfigurationLine(line string) (key string, value string) {
	// split line on "="
	lineArray := strings.Split(line, "=")
	if len(lineArray) == 2 {
		// assign to variables
		key = lineArray[0]
		value = lineArray[1]
		// if value contains double quotes, replace them with empty string
		if strings.Contains(value, "\"") {
			value = strings.Replace(value, "\"", "", -1)
		}
		if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
			value = strings.Replace(value, "'", "", -1)
		}
		return strings.TrimSpace(key), strings.TrimSpace(value)
	}
	return "", ""
}

//NOSONAR
// assignToConfiguration assigns a key / value pair to the configuration object
func assignToConfiguration(key string, value string) {
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
		_ = addK8sManagementConfig(key, value)
	}
}

func addKubernetesCertificate(key string, value string) (success bool) {
	if key == "KUBERNETES_SERVER_CERTIFICATE" {
		yamlCfg.Kubernetes.Certificates.Default = value
		success = true
	} else if strings.HasPrefix(key, "KUBERNETES_SERVER_CERTIFICATE_") {
		context := strings.TrimPrefix(key, "KUBERNETES_SERVER_CERTIFICATE_")
		// assign to old config
		if yamlCfg.Kubernetes.Certificates.Contexts == nil {
			yamlCfg.Kubernetes.Certificates.Contexts = make(map[string]string)
		}
		yamlCfg.Kubernetes.Certificates.Contexts[context] = value
	}
	return success
}

func addJenkinsJCasCJobConfig(key string, value string) (success bool) {
	switch key {
	case "JENKINS_JCASC_CONFIGURATION_URL":
		yamlCfg.Jenkins.Jcasc.ConfigurationURL = value
		break
	case "JENKINS_JOBDSL_BASE_URL":
		yamlCfg.Jenkins.JobDSL.BaseURL = value
		break
	case "JENKINS_JOBDSL_REPO_VALIDATE_PATTERN":
		yamlCfg.Jenkins.JobDSL.RepoValidatePattern = value
		break
	case "JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL":
		yamlCfg.Jenkins.Jcasc.SeedJobURL = value
		break
	}
	return success
}

func addJenkinsMasterConfig(key string, value string) (success bool) {
	switch key {
	case "JENKINS_MASTER_ADMIN_PASSWORD":
		yamlCfg.Jenkins.Controller.Passwords.AdminUser = value
		success = true
		break
	case "JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED":
		yamlCfg.Jenkins.Controller.Passwords.AdminUserEncrypted = replaceUnneededChars(value)
		success = true
		break
	case "JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED":
		yamlCfg.Jenkins.Controller.Passwords.DefaultUserEncrypted = replaceUnneededChars(value)
		success = true
		break
	case "JENKINS_MASTER_DEFAULT_LABEL":
		yamlCfg.Jenkins.Controller.CustomJenkinsLabel = value
		success = true
		break
	case "JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS":
		yamlCfg.Jenkins.Jcasc.AuthorizationStrategy.AllowAnonymousRead, _ = strconv.ParseBool(value)
		success = true
		break
	case "JENKINS_MASTER_DEFAULT_URI_PREFIX":
		yamlCfg.Jenkins.Controller.DefaultURIPrefix = value
		success = true
		break
	case "JENKINS_MASTER_DEPLOYMENT_NAME":
		yamlCfg.Jenkins.Controller.DeploymentName = value
		success = true
		break
	case "JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS":
		yamlCfg.Jenkins.Persistence.StorageClass = value
		success = true
		break
	case "JENKINS_MASTER_PERSISTENCE_ACCESS_MODE":
		yamlCfg.Jenkins.Persistence.AccessMode = value
		success = true
		break
	case "JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE":
		yamlCfg.Jenkins.Persistence.StorageSize = value
		success = true
		break
	case "JENKINS_MASTER_CONTAINER_IMAGE":
		yamlCfg.Jenkins.Container.Image = value
		success = true
		break
	case "JENKINS_MASTER_CONTAINER_IMAGE_TAG":
		yamlCfg.Jenkins.Container.Tag = value
		success = true
		break
	case "JENKINS_MASTER_CONTAINER_PULL_POLICY":
		yamlCfg.Jenkins.Container.PullPolicy = value
		success = true
		break
	case "JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME":
		yamlCfg.Jenkins.Container.PullSecret = value
		success = true
		break
	}
	return success
}

func addNginxConfig(key string, value string) (success bool) {
	switch key {
	case "NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE":
		yamlCfg.Nginx.Ingress.Container.Image = value
		success = true
		break
	case "NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS":
		yamlCfg.Nginx.Ingress.Container.PullSecret = value
		success = true
		break
	case "NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE":
		yamlCfg.Nginx.Ingress.Deployment.ForEachNamespace, _ = strconv.ParseBool(value)
		success = true
		break
	case "NGINX_INGRESS_DEPLOYMENT_NAME":
		yamlCfg.Nginx.Ingress.Deployment.DeploymentName = value
		success = true
		break
	case "NGINX_INGRESS_ANNOTATION_CLASS":
		yamlCfg.Nginx.Ingress.Annotationclass = value
		success = true
		break
	}
	return success
}

func addLoadBalancerConfig(key string, value string) (success bool) {
	switch key {
	case "NGINX_LOADBALANCER_ENABLED":
		yamlCfg.Nginx.Loadbalancer.Enabled, _ = strconv.ParseBool(value)
		success = true
		break
	case "NGINX_LOADBALANCER_HTTP_PORT":
		yamlCfg.Nginx.Loadbalancer.Ports.HTTP, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	case "NGINX_LOADBALANCER_HTTP_TARGETPORT":
		yamlCfg.Nginx.Loadbalancer.Ports.HTTPTarget, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	case "NGINX_LOADBALANCER_HTTPS_PORT":
		yamlCfg.Nginx.Loadbalancer.Ports.HTTPS, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	case "NGINX_LOADBALANCER_HTTPS_TARGETPORT":
		yamlCfg.Nginx.Loadbalancer.Ports.HTTPSTarget, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	case "NGINX_LOADBALANCER_ANNOTATIONS_ENABLED":
		yamlCfg.Nginx.Loadbalancer.Annotations.Enabled, _ = strconv.ParseBool(value)
		success = true
		break
	case "NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME":
		yamlCfg.Nginx.Loadbalancer.ExternalDNS.HostName = value
		success = true
		break
	case "NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_TTL":
		yamlCfg.Nginx.Loadbalancer.ExternalDNS.TTL, _ = strconv.ParseUint(value, 10, 16)
		success = true
		break
	}
	return success
}

func addCredentialsConfig(key string, value string) (success bool) {
	switch key {
	case "KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID":
		yamlCfg.Jenkins.Jcasc.CredentialIDs.Docker = value
		success = true
		break
	case "MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID":
		yamlCfg.Jenkins.Jcasc.CredentialIDs.Maven = value
		success = true
		break
	case "NPM_REPOSITORY_SECRETS_CREDENTIALS_ID":
		yamlCfg.Jenkins.Jcasc.CredentialIDs.Npm = value
		success = true
		break
	case "VCS_REPOSITORY_SECRETS_CREDENTIALS_ID":
		yamlCfg.Jenkins.Jcasc.CredentialIDs.Vcs = value
		success = true
		break
	}
	return success
}

func addCommonConfig(key string, value string) (success bool) {
	switch key {
	case "LOG_LEVEL":
		yamlCfg.K8SManagement.Log.Level = value
		success = true
		break
	case "GLOBAL_SECRETS_FILE":
		yamlCfg.K8SManagement.Project.SecretFiles = value
		success = true
		break
	case "IP_CONFIG_FILE":
		yamlCfg.K8SManagement.IPConfig.File = value
		success = true
		break
	case "IP_CONFIG_FILE_DUMMY_PREFIX":
		yamlCfg.K8SManagement.IPConfig.DummyPrefix = value
		success = true
		break
	case "PROJECTS_BASE_DIRECTORY":
		yamlCfg.K8SManagement.Project.BaseDirectory = value
		success = true
		break
	case "TEMPLATES_BASE_DIRECTORY":
		yamlCfg.K8SManagement.Project.TemplateDirectory = value
		success = true
		break
	}
	return success
}

func addK8sManagementConfig(key string, value string) (success bool) {
	switch key {
	case "K8S_MGMT_LOGGING_LOGFILE":
		yamlCfg.K8SManagement.Log.File = value
		success = true
		break
	case "K8S_MGMT_LOGGING_ENCODING":
		yamlCfg.K8SManagement.Log.Encoding = value
		success = true
		break
	case "K8S_MGMT_LOGGING_OVERWRITE_ON_START":
		yamlCfg.K8SManagement.Log.OverwriteOnRestart, _ = strconv.ParseBool(value)
		success = true
		break
	case "K8S_MGMT_VERSION_CHECK":
		yamlCfg.K8SManagement.VersionCheck, _ = strconv.ParseBool(value)
		success = true
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

func migrateFromCnfToYaml() (string, bool) {
	var yamlOutByte []byte
	yamlOutByte, err := yaml.Marshal(yamlCfg)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("Unable to create YAML file", err.Error())
		return "FAILED - Unable to create YAML file", false
	}
	var yamlConfig = string(yamlOutByte)
	loggingstate.AddInfoEntryAndDetails("New custom configuration", yamlConfig)

	var basePath = configuration.GetConfiguration().GetProjectBaseDirectory()
	var k8sConfigFile = configuration.GetConfiguration().CustomConfig.K8SManagement.ConfigFile
	var configFile = files.AppendPath(basePath, k8sConfigFile)
	if files.FileOrDirectoryExists(configFile) {
		return "FAILED - Config already exists", false
	}
	err = ioutil.WriteFile(configFile, yamlOutByte, 0644)
	if err != nil {
		return "FAILED - Unable to write new IP config", false
	}
	return "SUCCESS", true
}

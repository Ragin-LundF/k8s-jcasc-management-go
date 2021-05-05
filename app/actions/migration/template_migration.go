package migration

import (
	"fmt"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
)

// MigrateTemplatesToV3 : Migrate the templates to V3 with golang template system
func MigrateTemplatesToV3() string {
	var errors []string
	var placeholdersToMigrate = createMigrationPlaceholderMap()
	var cloudTemplateDirectory = files.AppendPath(configuration.GetConfiguration().GetProjectTemplateDirectory(), constants.DirProjectTemplateCloudTemplates)

	for oldPlaceholder, newPlaceholder := range placeholdersToMigrate {
		// replace in main templates
		err := replacePlaceholderInTemplates(configuration.GetConfiguration().GetProjectTemplateDirectory(), oldPlaceholder, newPlaceholder)
		if err != nil {
			errors = append(errors, err.Error())
		}
		// replace in cloud.kubernetes sub templates
		err = replacePlaceholderInTemplates(cloudTemplateDirectory, oldPlaceholder, newPlaceholder)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	var status string
	var errorMessages = ""
	if len(errors) == 0 {
		status = "SUCCESS"
	} else {
		status = "FAILED"
		for _, errorMessage := range errors {
			errorMessages = fmt.Sprintf("%v\n\n%v", errorMessages, errorMessage)
		}
	}

	return fmt.Sprintf("Status template migration: %v\n\n%v", status, errorMessages)
}

// createMigrationPlaceholderMap returns the map with the old and new placeholders
func createMigrationPlaceholderMap() map[string]string {
	return map[string]string{
		"##K8S_MGMT_JENKINS_SYSTEM_MESSAGE##":                "{{ .JCasc.SystemMessage }}",
		"##JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED##":        "{{ .JCasc.SecurityRealm.LocalUsers.AdminPassword }}",
		"##JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED##": "{{ .JCasc.SecurityRealm.LocalUsers.UserPassword }}",
		"##KUBERNETES_SERVER_CERTIFICATE##":                  "{{ .JCasc.Clouds.Kubernetes.ServerCertificate }}",
		"##NAMESPACE##":                                      "{{ .Base.Namespace }}",
		"##JENKINS_MASTER_DEPLOYMENT_NAME##":                 "{{ .Base.DeploymentName }}",
		"##K8S_MGMT_JENKINS_CLOUD_TEMPLATES##":               "{{ .JCasc.Clouds.Kubernetes.Templates.AdditionalCloudTemplates }}",
		"##KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID##":      "{{ .JCasc.CredentialIDs.DockerRegistryCredentialsID }}",
		"##JENKINS_MASTER_DEFAULT_URI_PREFIX##":              "{{ .Base.JenkinsUriPrefix }}",
		"##JENKINS_URL##":                                    "{{ .Base.JenkinsURL }}",
		"##JENKINS_MASTER_DEFAULT_LABEL##":                   "{{ .JenkinsHelmValues.Controller.CustomJenkinsLabels }}",
		"##JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL##":             "{{ .JCasc.JobsConfig.JobsSeedRepository}}",
		"##PROJECT_JENKINS_JOB_DEFINITION_REPOSITORY##":      "{{ .JCasc.JobsConfig.JobsDefinitionRepository }}",

		"##VCS_REPOSITORY_SECRETS_CREDENTIALS_ID##":   "{{ .JCasc.CredentialIDs.VcsRepositoryCredentialsID }}",
		"##NPM_REPOSITORY_SECRETS_CREDENTIALS_ID##":   "{{ .JCasc.CredentialIDs.NpmRepositorySecretsCredentialsID }}",
		"##MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID##": "{{ .JCasc.CredentialIDs.MavenRepositorySecretsCredentialsID }}",

		"##JENKINS_MASTER_CONTAINER_IMAGE##":                  "{{ .JenkinsHelmValues.Controller.Image }}",
		"##JENKINS_MASTER_CONTAINER_IMAGE_TAG##":              "{{ .JenkinsHelmValues.Controller.Tag }}",
		"##JENKINS_MASTER_CONTAINER_PULL_POLICY##":            "{{ .JenkinsHelmValues.Controller.ImagePullPolicy }}",
		"##JENKINS_MASTER_CONTAINER_IMAGE_PULL_SECRET_NAME##": "{{ .JenkinsHelmValues.Controller.ImagePullSecretName }}",
		"##JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS##":       "{{ .JenkinsHelmValues.Controller.AuthorizationStrategyAllowAnonymousRead }}",
		"##JENKINS_MASTER_ADMIN_PASSWORD##":                   "{{ .JenkinsHelmValues.Controller.AdminPassword }}",
		"##JENKINS_JCASC_CONFIGURATION_URL##":                 "{{ .JenkinsHelmValues.Controller.SidecarsConfigAutoReloadFolder }}",
		"##JENKINS_MASTER_PERSISTENCE_EXISTING_CLAIM##":       "{{ .Base.ExistingVolumeClaim }}",
		"##JENKINS_MASTER_PERSISTENCE_STORAGE_CLASS##":        "{{ .JenkinsHelmValues.Persistence.StorageClass }}",
		"##JENKINS_MASTER_PERSISTENCE_ACCESS_MODE##":          "{{ .JenkinsHelmValues.Persistence.AccessMode }}",
		"##JENKINS_MASTER_PERSISTENCE_STORAGE_SIZE##":         "{{ .JenkinsHelmValues.Persistence.Size }}",

		"##NGINX_INGRESS_ANNOTATION_CLASS##":                  "{{ .Nginx.Ingress.AnnotationIngressClass }}",
		"##PUBLIC_IP_ADDRESS##":                               "{{ .Base.IPAddress }}",
		"##NGINX_INGRESS_DEPLOYMENT_NAME##":                   "{{ .Nginx.Ingress.DeploymentName }}",
		"##NGINX_INGRESS_CONTROLLER_CONTAINER_IMAGE##":        "{{ .Nginx.Ingress.ContainerImage }}",
		"##NGINX_INGRESS_CONTROLLER_CONTAINER_PULL_SECRETS##": "{{ .Nginx.Ingress.ImagePullSecrets }}",
		"##NGINX_INGRESS_CONTROLLER_FOR_NAMESPACE##":          "{{ .Nginx.Ingress.EnableControllerForNamespace }}",
		"##NGINX_LOADBALANCER_ENABLED##":                      "{{ .Nginx.LoadBalancer.Enabled }}",
		"##NGINX_LOADBALANCER_HTTP_PORT##":                    "{{ .Nginx.LoadBalancer.Ports.HTTP.Port }}",
		"##NGINX_LOADBALANCER_HTTP_TARGETPORT##":              "{{ .Nginx.LoadBalancer.Ports.HTTP.TargetPort }}",
		"##NGINX_LOADBALANCER_HTTPS_PORT##":                   "{{ .Nginx.LoadBalancer.Ports.HTTPS.Port }}",
		"##NGINX_LOADBALANCER_HTTPS_TARGETPORT##":             "{{ .Nginx.LoadBalancer.Ports.HTTPS.TargetPort }}",
		"##NGINX_LOADBALANCER_ANNOTATIONS_ENABLED##":          "{{ .Nginx.LoadBalancer.Annotations.Enabled }}",
		"##NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_HOSTNAME##": "{{ .Nginx.LoadBalancer.Annotations.ExternalDnsHostname }}",
		"##NGINX_LOADBALANCER_ANNOTATIONS_EXT_DNS_TTL##":      "{{ .Nginx.LoadBalancer.Annotations.ExternalDnsTtl }}",

		"##K8S_MGMT_PERSISTENCE_VOLUME_CLAIM_NAME##": "{{ .Base.ExistingVolumeClaim }}",
		"##PROJECT_DIRECTORY##":                      "{{ .Base.Namespace }}",
	}
}

// replacePlaceholderInTemplates : Replace placeholder with value in all project files
func replacePlaceholderInTemplates(directory string, placeholder string, newValue string) (err error) {
	var templateFiles, _ = files.LoadTemplateFilesOfDirectory(directory)
	for _, templateFile := range templateFiles {
		err = replacePlaceholderInTemplate(templateFile, placeholder, newValue)
		if err != nil {
			return err
		}
	}
	return nil
}

// replacePlaceholderInTemplate : Replace a placeholder in a file with a value
func replacePlaceholderInTemplate(filename string, placeholder string, newValue string) (err error) {
	if files.FileOrDirectoryExists(filename) {
		successful, err := files.ReplaceStringInFile(filename, placeholder, newValue)
		if !successful || err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to replace [%s] in file [%s]", placeholder, filename), err.Error())
			return err
		}
	}

	return nil
}

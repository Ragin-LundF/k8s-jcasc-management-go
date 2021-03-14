# Jenkins Configuration as Code (JcasC) Helm Values Placeholder

Jenkins Configuration as Code Helm Values are part of the JcasC Helm Chart.

The tool uses the `jcasc_config` template to create the required configuration in Kubernetes.
It is also possible to use the placeholders in other templates.

## Placeholder variables for Jenkins configuration as Code Helm Values

| Placeholder | Description | Source | old config |
| --- | --- | --- | --- |
| `{{ .JCasc.Clouds.Kubernetes.Templates.AdditionalCloudTemplates }}` | Placeholder for `clouds.kubernetes.templates` for Jenkins Helm Values | user sub cloud template selection | n/a |
| `{{ .JCasc.CredentialIDs.DockerRegistryCredentialsID }}` | Placeholder for common Docker Jenkins credentialIDs | configuration `jenkins.jcasc.credentialIDs.docker` | `KUBERNETES_DOCKER_REGISTRY_CREDENTIALS_ID` |
| `{{ .JCasc.CredentialIDs.MavenRepositorySecretsCredentialsID }}` | Placeholder for common Maven Jenkins credentialIDs | configuration `jenkins.jcasc.credentialIDs.maven` | `MAVEN_REPOSITORY_SECRETS_CREDENTIALS_ID` |
| `{{ .JCasc.CredentialIDs.NpmRepositorySecretsCredentialsID }}` | Placeholder for common NPM Jenkins credentialIDs | configuration `jenkins.jcasc.credentialIDs.npm` | `NPM_REPOSITORY_SECRETS_CREDENTIALS_ID` |
| `{{ .JCasc.CredentialIDs.VcsRepositoryCredentialsID }}` | Placeholder for common VCS Jenkins credentialIDs | configuration `jenkins.jcasc.credentialIDs.vcs` | `VCS_REPOSITORY_SECRETS_CREDENTIALS_ID` |
| `{{ .JCasc.JobsConfig.JobsAvailable }}` | Can be used to check if jobs are available with `{{ if .JCasc.JobsConfig.JobsAvailable }}` | calculated; true if seed and job repositories are not empty | n/a |
| `{{ .JCasc.JobsConfig.JobsSeedRepository }}` | Placeholder for `jobs` configuration to define the seed job repository | configuration `jenkins.jcasc.seedJobURL` | `JENKINS_JOBDSL_SEED_JOB_SCRIPT_URL` |
| `{{ .JCasc.JobsConfig.JobsDefinitionRepository }}` | Placeholder for `jobs` configuration to define the job definition repository | user input | n/a |
| `{{ .JCasc.SecurityRealm.LocalUsers.AdminPassword }}` | Placeholder for `securityRealm.local.users` encrypted admin password | configuration `jenkins.controller.passwords.adminUserEncrypted` | `JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED` |
| `{{ .JCasc.SecurityRealm.LocalUsers.UserPassword }}` | Placeholder for `securityRealm.local.users` encrypted user password | configuration `jenkins.controller.passwords.defaultUserEncrypted` | `JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED` |
| `{{ .JCasc.SystemMessage }}` | Placeholder for `systemMessage` which is the Jenkins system welcome message | user input | n/a |

## More placeholder
| Description | Link |
| --- | --- |
| Common base placeholder | [TemplatePlaceholder.md](TemplatePlaceholder.md) |
| Jenkins deployment `jenkins_helm_values.yaml` placeholder | [JenkinsHelmValuesPlaceholder.md](JenkinsHelmValuesPlaceholder.md) |
| Nginx Ingress Controller `nginx_ingress_helm_values.yaml` placeholder | [NginxIngressControllerPlaceholder.md](NginxIngressControllerPlaceholder.md) |
| Persistent Volume Claim `pvc_claim.yaml` placeholder | [PersistentVolumeClaimPlaceholder.md](PersistentVolumeClaimPlaceholder.md) |

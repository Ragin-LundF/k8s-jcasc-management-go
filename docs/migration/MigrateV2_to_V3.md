# Migration from v2.x.x to v3.x.x

Version 3 introduces some changes on configuration and templating level.
The old configuration inherited from the Bash version is completely switched to YAML in this release.
This allows extended configurations for e.g. the IP addresses by storing additional attributes like domain.
It is also the prerequisite for future features such as deployments from the templates and saved configuration, instead of previously generated helmet configurations.

The templates have also been changed from the old `##PLACEHOLDER##` syntax to Golang templates, such as those used in Helm Charts for example.

To make this integration as good as possible, the tool supports some migration steps with small tools to keep the manual work as low as necessary.


## First steps
The first step should be to change the `k8s_jcasc_custom.cnf` to the new `k8s_jcasc_custom.yaml`.
This file now supports only the base configuration to find the actual project.
All further configurations must be stored in a project configuration (formerly the `K8S_MGMT_ALTERNATIVE_CONFIG_FILE` attribute). 

This file should look like that:

```yaml
k8sManagement:
  # Configuration file for project specific overrides. This file must be relative to the `basePath`.
  configFile: "./config/k8s_jcasc_mgmt_custom.yaml"
  # Base path for all projects. The path can be specified absolutely.
  basePath: "/deployments/k8s-jcasc-manaagement"
```

_Please note that the configuration file should be set to the current custom configuration file `.cnf`, but with the file extension `.yaml`.
The migration tool will make sure that the file is read with the same name but with the extension `.cnf`.
This simplifies the process to not have to adjust the file multiple times._

After this base custom configuration is available, the configuration and the templates can be transformed to YAML with the new structure.

### Transforming the configuration
The transformation of the configuration can be done via the tool as described below.
After the transformation, the configuration should be compared again before deleting the old `.cnf` file.

#### CLI mode
The CLI mode supports the transformation with the following command:

```bash
go run k8s-jcasc-mgmt.go -cli -migrate-configs-v2
```

After this is completed, there should be a new `ip_config.yaml` file.
Please compare it to the old one. If everything is ok, the old ip_config.cnf file can be deleted.

#### GUI mode
In GUI mode, after changing the basic configuration, you can simply start the tool.
It now offers the option `Tools` -> `Migrate config v2 -> v3` in the main menu.

After this was clicked, it shows the current status.

### Transforming the templates
_This part is only valid if custom templates are maintained (which is mostly true)._

As with the configuration, the templates for the value files can also be transformed via the tool.
This also includes the `cloud-templates`.

Since templates do not create new files but modify existing ones, since it is primarily an exchange of placeholders, the result should always be checked before committing/pushing.

#### CLI mode
For the CLI the following command can be used (it can also be combined with the option above):

```bash
go run k8s-jcasc-mgmt.go -cli -migrate-templates-v2
```

#### GUI mode
In GUI mode, the procedure is very similar to the configuration transformation.

Simply start the App and then go to the main menu `Tools` -> `Migrate templates v2 -> v3`.

#### Postprocessing
##### jcasc_config.yaml
The only step, which has to be evaluated here manually is to check, that the encrypted passwords have no `'` character inside.

In this structure

```yaml
  securityRealm:
    local:
      allowsSignup: false
      users:
        - id: "admin"
          password: "#jbcrypt:{{ .JCasc.SecurityRealm.LocalUsers.AdminPassword }}"
        - id: "project-user"
          password: "#jbcrypt:{{ .JCasc.SecurityRealm.LocalUsers.UserPassword }}"
```

the old templates (and project files) have a `#jbcrypt:'<password hash>'` inside to be compatible with the old Bash version.
For the new configuration, those `'` characters must be removed.
Else Jenkins is not starting and throws an `IllegalArgumentException` with the message `this method should only be called with a pre-hashed password`.

##### jenkins_helm_values.yaml
After the transformation above was done, the file `jenkins_helm_values.yaml` has to be edited by hand.

**_All changes must also be made in the project templates._** 

_**It is best to compare the project-specific file with the original template.**_

First ensure, that the file starts with:

```yaml
controller:
  # Used for label app.kubernetes.io/component
  componentName: "{{ .Base.DeploymentName }}"
```

Please take care, that the `componentName` has to be set here too.

A bit below the following part has to be exchanged:

The old definition
```yaml
master:
  [...]
  authorizationStrategy: |-
      <authorizationStrategy class="hudson.security.FullControlOnceLoggedInAuthorizationStrategy">
        <denyAnonymousReadAccess>##JENKINS_MASTER_DENY_ANONYMOUS_READ_ACCESS##</denyAnonymousReadAccess>
      </authorizationStrategy>
  [...]
```

must be changed to

```yaml
controller:
  [...]
  JCasC:
    enabled: true
    # Define, if we want to allow anonymous read access
    authorizationStrategy: |-
      loggedInUsersCanDoAnything:
        allowAnonymousRead: {{ .JenkinsHelmValues.Controller.AuthorizationStrategyAllowAnonymousRead }}
  [...]
```

Next point is the sidecar folder configuration.
In the old templates it looks like this:

```yaml
master:
  [...]
  sidecars:
    configAutoReload:
      # folder in the pod that should hold the collected dashboards:
      folder: "https://domain.tld/jcasc_config.yaml"
  [...]
```

In the new template the configAutoReload needs to be disabled:

```yaml
controller:
  [...]
  sidecars:
    configAutoReload:
      # This should be disabled. If enabled Jenkins tries to mount the URL as a folder, which is not working.
      # If it is disabled, it sets the variable only to the CASC_JENKINS_CONFIG environment variable, what we want.
      enabled: false
      # folder in the pod that should hold the collected dashboards:
      folder: "https://domain.tld/jcasc_config.yaml"
```

## Migrating existing Helm value files

For already created projects, some manual changes are also needed.

The main change takes place in the `jenkins_helm_values.yaml` file.
The reason for this is the new Helm Charts from Jenkins, which have also been updated.

### Exchanging `master` with `controller`
As the simplest step, the previous root node in these project files concerned to

```yaml
master:
```

can be replaced by

```yaml
controller:
  # Used for label app.kubernetes.io/component
  componentName: "jenkins-controller"
```

The `controller` element is the replacement for the old `master`.
`componentName` was added here and contains the deployment name, which is configured normally under:

```yaml
jenkins:
  [...]
  controller:
    [...]
    deploymentName: "jenkins-controller"
    [...]
```

in your custom configuration.
The default is value is `jenkins-controller`.
In the previous versions, this was named `jenkins-master` and should be replaced.

The last change concerns the following place in the (old) configuration:

```yaml
master:
  [...]
  authorizationStrategy: |-
    <authorizationStrategy class="hudson.security.FullControlOnceLoggedInAuthorizationStrategy">
      <denyAnonymousReadAccess>false</denyAnonymousReadAccess>
    </authorizationStrategy>
```

In the old version, XML was still used here and anonymous access was denied.

The new Jenkins Helm Charts use the following syntax:

```yaml
controller:
  JCasC:
    enabled: true
    # Define if we want to allow anonymous read access
    authorizationStrategy: |-
      loggedInUsersCanDoAnything:
        allowAnonymousRead: true
```

Since the new syntax `allowAnonymousRead` is opposite to `denyAnonymousReadAccess`, care must be taken to negate the value as well (`true` becomes `false` and vice versa).



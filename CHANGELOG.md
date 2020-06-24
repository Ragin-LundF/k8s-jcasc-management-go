# 1.13.1
* Hotfix for configuration of encrypted users.
  * Fixes the issue, that `JENKINS_MASTER_ADMIN_PASSWORD_ENCRYPTED` and `JENKINS_MASTER_PROJECT_USER_PASSWORD_ENCRYPTED` need encrypted password surrounded with `'` characters, that the `bash` version will not interpret this configuration as arguments.

# 1.13.0
* Support for create namespace from k8s-jcasc-mgmt
* Default logging enabled.
  * `k8s_jcasc_mgmt.cnf` has now the following 3 parameters:
    * `K8S_MGMT_LOGGING_LOGFILE`: default is `output.log`. This defines the default logfile. This value can be overwritten with the `-logfile` argument or simply comment it to disable it.
    * `K8S_MGMT_LOGGING_ENCODING`: default is: `console`. Allows the default logging type (`console` or `json`)
    * `K8S_MGMT_LOGGING_OVERWRITE_ON_START`: default is: `true`. Defines if logfile should be re-created on start. In this case, the system moves the old log (if exists) to the defined logfile name with suffix `.1`. Example: `output.log.1`.

# 1.12.1
Features:

* create new projects for Jenkins administration
* manage secrets
    * encrypt/decrypt secrets for secure commit to a VCS (version control system)
    * apply secrets to Kubernetes
        * for each project while installation or as an update (`applySecrets`)
        * for all known namespaces, that are configured in the `ip_config.cnf` file (applySecretsToAll)
    * store secrets globally for easy administration
    * store secrets per project for more security
* manage the Jenkins instances for a namespace with the project configuration
    * install
        * create namespace if it does not exist
        * install Jenkins
        * install nginx-ingress-controller per namespace (if configured)
        * install load balancer and ingress for Jenkins
    * uninstall
        * uninstall Jenkins installation
        * uninstall nginx-ingress-controller per namespace (if configured)
        * uninstall load balancer and ingress for Jenkins (other ingress routes will not be changed)
    * upgrade
* internal log viewer
* logging in JSON or console format

-> This release is equal to the 1.12.1 version of the bash implementation (https://github.com/Ragin-LundF/k8s-jcasc-management).

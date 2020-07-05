# 2.0.1

## Introducing new UI
* K8S-Jcasc-Management supports now native UIs for the following platforms:
  * Windows
  * Linux and BSD
  * MacOS X

This was realized with the [fyne](https://fyne.io/) framework.

To use K8S-JcasC-Management on a CLI, you can start it with the `-cli` flag.

## Build hints
If you have trouble to compile the project, please visit the [fyne developer](https://developer.fyne.io/started/) site first to check the prerequisites.
On Windows, it is recommended to install [TDM-GCC - tdm-gcc.tdragon.net](https://tdm-gcc.tdragon.net), which works very easy. It can be required to set the PATH variable to the `TDM-GCC` directory, if the `go` compiler still wants a `GCC`.

### Using on server-systems only (without GUI)
If you have trouble or if you want to use it on a server without `X11`, you can also exchange the `!ignore` and `ignore` in first line comment at the `/app/app_cli.go` and `/app/app_gui.go` file.
Golang will then use the `app_cli.go` file to compile and ignores the GUI implementation completely.

## Screenshots

### Welcome (Windows)
![alt text](docs/images/screenshot_gui_welcome_win.png "K8S GUI Welcome")

### Deployment (Mac Dark Theme)
![alt text](docs/images/screenshot_gui_deployment.png "K8S GUI Deployment")

### Create Project (Mac Light Theme)
![alt text](docs/images/screenshot_gui_createprj_light.png "K8S GUI Project Create")

# 1.13.1
* Hotfix for the configuration of encrypted users.
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

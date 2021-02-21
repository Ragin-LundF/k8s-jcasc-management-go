package kubernetesactions

import (
	"k8s-management-go/app/utils/kubectl"
	"k8s-management-go/app/utils/loggingstate"
	"sort"
	"strings"
)

// KubernetesConfiguration defines the current configuration
type KubernetesConfiguration struct {
	contexts       []string
	currentContext string
}

var kubernetesConfig KubernetesConfiguration

// CurrentContext returns the currently selected kubernetes context
func (k8sConfig KubernetesConfiguration) CurrentContext() string {
	return k8sConfig.currentContext
}

// Contexts returns the available kubernetes contexts
func (k8sConfig KubernetesConfiguration) Contexts() []string {
	return k8sConfig.contexts
}

// hasNoContexts checks if contexts are available
func (k8sConfig KubernetesConfiguration) HasNoContexts() bool {
	if len(k8sConfig.contexts) == 0 {
		return true
	}
	return false
}

// readKubernetesContexts gets the currently configured contexts
func readKubernetesContexts() {
	kubectlCmdArgs := []string{
		"get-contexts",
		"-o",
		"name",
	}
	kubectlCmdOutput, err := kubectl.ExecutorKubectl("config", kubectlCmdArgs)

	if err != nil {
		loggingstate.AddErrorEntryAndDetails("Unable to get Kubernetes configuration", err.Error())
		return
	}

	if kubectlCmdOutput != "" {
		contextNames := strings.Fields(kubectlCmdOutput)
		kubernetesConfig.contexts = append(kubernetesConfig.contexts, contextNames...)
	} else {
		return
	}
}

// readCurrentKubernetesContext gets the current context
func readCurrentKubernetesContext() {
	kubectlCmdArgs := []string{
		"current-context",
	}
	kubectlCmdOutput, err := kubectl.ExecutorKubectl("config", kubectlCmdArgs)

	if err != nil {
		loggingstate.AddErrorEntryAndDetails("Unable to get current Kubernetes context", err.Error())
	}

	kubectlCmdOutput = strings.TrimPrefix(kubectlCmdOutput, " ")
	kubectlCmdOutput = strings.TrimSuffix(kubectlCmdOutput, " ")
	kubectlCmdOutput = strings.ReplaceAll(kubectlCmdOutput, "\n", "")
	kubectlCmdOutput = strings.ReplaceAll(kubectlCmdOutput, "\r", "")

	kubernetesConfig.currentContext = kubectlCmdOutput
}

// GetKubernetesConfig returns the current configuration with the available contexts
func GetKubernetesConfig() KubernetesConfiguration {
	if kubernetesConfig.currentContext == "" {
		readCurrentKubernetesContext()
	}
	if kubernetesConfig.HasNoContexts() {
		readKubernetesContexts()
	}
	return kubernetesConfig
}

// ReloadKubernetesConfig : reload the current kubernetes context
func ReloadKubernetesContext() {
	readCurrentKubernetesContext()
	if kubernetesConfig.HasNoContexts() {
		readKubernetesContexts()
	}
}

// SwitchKubernetesConfig switches the current context
func SwitchKubernetesConfig(context string) (err error) {
	kubectlCmdArgs := []string{
		"use-context",
		context,
	}
	_, err = kubectl.ExecutorKubectl("config", kubectlCmdArgs)

	if err != nil {
		loggingstate.AddErrorEntryAndDetails("Unable to switch Kubernetes configuration", err.Error())
	}
	readCurrentKubernetesContext()

	return err
}

// ActionReadK8SContextWithFilter is a kubernetes context loader and filter
func ActionReadK8SContextWithFilter(filter *string) (k8sContextsList []string) {
	if filter != nil && *filter != "" {
		for _, k8sContext := range GetKubernetesConfig().Contexts() {
			if strings.Contains(k8sContext, *filter) {
				k8sContextsList = append(k8sContextsList, k8sContext)
			}
		}
	} else {
		k8sContextsList = append(k8sContextsList, kubernetesConfig.Contexts()...)
	}
	sort.Strings(k8sContextsList)
	return k8sContextsList
}

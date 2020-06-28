package createproject

import "k8s-management-go/app/utils/loggingstate"

// delegation method for replacement of global configuration
func ActionReplaceGlobalConfigDelegation(projectDirectory string) (success bool, err error) {
	success, err = ActionReplaceGlobalConfigNginxIngressCtrlHelmValues(projectDirectory)
	if !success || err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to replace global nginx-ingress-controller Helm values...abort", err.Error())
		return false, err
	}
	success, err = ActionReplaceGlobalConfigJCasCValues(projectDirectory)
	if !success || err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to replace global JCasc values...abort", err.Error())
		return false, err
	}
	success, err = ActionReplaceGlobalConfigJenkinsHelmValues(projectDirectory)
	if !success || err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to replace global Jenkins Helm values...abort", err.Error())
		return false, err
	}
	success, err = ActionReplaceGlobalConfigPvcValues(projectDirectory)
	if !success || err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable to replace global PVC values...abort", err.Error())
		return false, err
	}
	return success, nil
}

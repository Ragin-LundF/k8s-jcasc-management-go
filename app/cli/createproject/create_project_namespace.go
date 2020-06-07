package createproject

import (
	"errors"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/cli/logoutput"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"regexp"
	"strings"
)

func ProjectWizardAskForNamespace() (namespace string, err error) {
	log := logger.Log()
	// Validator for namespace name
	validate := func(input string) error {
		// a namespace name cannot be longer than 63 characters
		if len(input) > 63 {
			return errors.New("Namespace name is too long! You can only use max. 63 characters. ")
		}
		// Regex to have DNS compatible string
		regex := regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
		if !regex.Match([]byte(input)) {
			return errors.New("Namespace is not valid! It must fit to DNS specification! ")
		}
		// check, that namespace was not already used
		for _, ipConfig := range models.GetIpConfiguration().Ips {
			if strings.ToLower(ipConfig.Namespace) == strings.ToLower(input) {
				return errors.New("Namespace already in use! ")
			}
		}
		return nil
	}

	// Prepare prompt
	dialogs.ClearScreen()
	namespace, err = dialogs.DialogPrompt("Enter namespace name", validate)
	// check if everything was ok
	if err != nil {
		logoutput.AddErrorEntryAndDetails("  -> Unable to get name of new namespace!", err.Error())
		log.Error("[ProjectWizardAskForNamespace] Unable to get name of new namespace. %v\n", err)
	}

	return namespace, err
}

// Replace Namespace in templates
func ProcessTemplateNamespace(projectDirectory string, namespace string) (success bool, err error) {
	log := logger.Log()

	templateFiles := []string{
		files.AppendPath(projectDirectory, constants.FilenameJenkinsConfigurationAsCode),
		files.AppendPath(projectDirectory, constants.FilenamePvcClaim),
		files.AppendPath(projectDirectory, constants.FilenameNginxIngressControllerHelmValues),
	}

	for _, templateFile := range templateFiles {
		if files.FileOrDirectoryExists(templateFile) {
			successful, err := files.ReplaceStringInFile(templateFile, constants.TemplateNamespace, namespace)
			if !successful || err != nil {
				logoutput.AddErrorEntryAndDetails("  -> Unable to replace namespace in file ["+templateFile+"]", err.Error())
				log.Error("[ProcessTemplateNamespace] Unable to replace namespace in file [%v], \n%v", templateFile, err)
				return false, err
			}
		}

	}
	return true, err
}

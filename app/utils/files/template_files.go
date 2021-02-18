package files

import (
	"fmt"
	"k8s-management-go/app/utils/loggingstate"
)

// ReplacePlaceholderInTemplates : Replace placeholder with value in all project files
func ReplacePlaceholderInTemplates(directory string, placeholder string, newValue string) (success bool, err error) {
	templateFiles, err := LoadTemplateFilesOfDirectory(directory)
	for _, templateFile := range templateFiles {
		success, err = ReplacePlaceholderInTemplate(templateFile, placeholder, newValue)
		if err != nil {
			return success, err
		}
	}
	return true, nil
}

// ReplacePlaceholderInTemplate : Replace a placeholder in a file with a value
func ReplacePlaceholderInTemplate(filename string, placeholder string, newValue string) (success bool, err error) {
	if FileOrDirectoryExists(filename) {
		successful, err := ReplaceStringInFile(filename, placeholder, newValue)
		if !successful || err != nil {
			loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to replace [%s] in file [%s]", placeholder, filename), err.Error())
			return false, err
		}
	}

	return true, nil
}

// LoadTemplateFilesOfDirectory : Load all template files of a directory
func LoadTemplateFilesOfDirectory(directory string) ([]string, error) {
	var configFiles = ".yaml"
	var fileFilter = FileFilter{
		Suffix: &configFiles,
	}
	filesInDirectory, err := ListFilesOfDirectoryWithFilter(directory, &fileFilter)
	if filesInDirectory == nil {
		loggingstate.AddErrorEntry("-> Could not find any yaml files in directory [" + directory + "].")
		return []string{}, err
	}

	var templateFiles []string
	for _, file := range *filesInDirectory {
		templateFiles = append(templateFiles, AppendPath(directory, file))
	}
	return templateFiles, nil
}

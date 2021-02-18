package project

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"text/template"
)

type Project struct {
	Namespace struct {
		Name string
	}
	PersistentVolumeClaim *PersistentVolumeClaim
}

// ProcessTemplates : Interface implementation to process templates with PVC placeholder
func (project *Project) ProcessTemplates(projectDirectory string) (err error) {
	templateFiles, err := files.LoadTemplateFilesOfDirectory(projectDirectory)
	if err != nil {
		return err
	}

	for _, templateFile := range templateFiles {
		err = processWithTemplateEngine(templateFile, *project)
		if err != nil {
			_ = os.Remove(projectDirectory)
			loggingstate.AddErrorEntryAndDetails(
				fmt.Sprintf("-> Unable to process file [%v] with template engine", templateFile),
				err.Error())
			return err
		}

		err = processWithPlaceholderReplace(templateFile, *project)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails(
				fmt.Sprintf("-> Unable to process file [%v] with placeholder replacement", templateFile),
				err.Error())
			return err
		}
	}

	return nil
}

// processWithTemplateEngine : Process files with template engine
func processWithTemplateEngine(filename string, project Project) (err error) {
	templates, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}

	var processedTemplate bytes.Buffer
	err = templates.Execute(&processedTemplate, project)
	if err != nil {
		return err
	}

	if processedTemplate.Len() > 0 {
		_ = ioutil.WriteFile(filename, processedTemplate.Bytes(), 0)
	}

	return nil
}

// processWithPlaceholderReplace : process files with old "##" templates
func processWithPlaceholderReplace(filename string, project Project) (err error) {
	return project.PersistentVolumeClaim.ProcessTemplates(filename)
}
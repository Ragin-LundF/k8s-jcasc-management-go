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

// Project : Structure for a complete project
type Project struct {
	Namespace             string
	JenkinsHelmValues     *jenkinsHelmValues
	PersistentVolumeClaim *persistentVolumeClaim
	Nginx                 *nginx
}

// NewProject : create a new Project
func NewProject(namespace string) Project {
	return Project{
		Namespace:             namespace,
		JenkinsHelmValues:     NewJenkinsHelmValues(),
		PersistentVolumeClaim: NewPersistentVolumeClaim(namespace),
		Nginx:                 NewNginx(namespace, nil, nil),
	}
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
	}

	return nil
}

// processWithTemplateEngine : Process files with template engine
func processWithTemplateEngine(filename string, project Project) (err error) {
	// replace JCasC URL
	var jcascUrl bytes.Buffer
	jcascUrlTemplate, err := template.New("JcasC-URL").Parse(project.JenkinsHelmValues.Master.SidecarsConfigAutoReloadFolder)
	if err != nil {
		return err
	}

	err = jcascUrlTemplate.Execute(&jcascUrl, project)
	if err != nil {
		return err
	}
	project.JenkinsHelmValues.Master.SidecarsConfigAutoReloadFolder = jcascUrl.String()

	// replace templates
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

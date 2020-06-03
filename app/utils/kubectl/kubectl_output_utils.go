package kubectl

import (
	"k8s-management-go/app/utils"
	"strings"
)

// function to check if a kubectl output contains a value for a field
func CheckIfKubectlOutputContainsValueForField(kubectlOutput string, fieldName string, searchedValue string) bool {
	found := false
	if len(kubectlOutput) > 1 && strings.Contains(kubectlOutput, "\n") {
		// first we split the output by lines
		kubectlOutArr := strings.Split(kubectlOutput, "\n")
		// if the output has more than 1 line (header + content) continue with processing
		if len(kubectlOutArr) > 1 {
			// split line in fields and search for index of "NAME"
			kubectlFieldsArr := strings.Fields(kubectlOutArr[0])
			kubectlNameIdx := utils.IndexOfArr(fieldName, kubectlFieldsArr)
			// iterate over kubectl output lines
			for idx, kubectlOutLine := range kubectlOutArr {
				// skip first line, because it was processed previously
				if idx > 0 {
					// split line into fields
					kubectlLineFieldsArr := strings.Fields(kubectlOutLine)
					// look if name field has entry with same name as pvc declaration
					if searchedValue == kubectlLineFieldsArr[kubectlNameIdx] {
						// name was found -> stop the processing; we have what we want
						found = true
						break
					}
				}
			}
		}
	}
	return found
}

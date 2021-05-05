package kubectl

import (
	"fmt"
	"k8s-management-go/app/utils/arrays"
	"k8s-management-go/app/utils/logger"
	"strings"
)

// CheckIfKubectlOutputContainsValueForField is a function to check if a kubectl output contains a value for a field
func CheckIfKubectlOutputContainsValueForField(kubectlOutput string, fieldName string, searchedValue string) bool {
	var found = false

	var fieldValues, err = FindFieldValuesInKubectlOutput(kubectlOutput, fieldName)
	if err != nil || len(fieldValues) == 0 {
		return false
	}

	// search for field
	for _, fieldValue := range fieldValues {
		if fieldValue == searchedValue {
			found = true
			break
		}
	}

	return found
}

// FindFieldValuesInKubectlOutput finds field values in kubectl output
func FindFieldValuesInKubectlOutput(kubectlOutput string, fieldName string) (fieldValues []string, err error) {
	// First find the field index for the field name
	lineIndex, fieldIndex, err := FindFieldIndexInKubectlOutput(kubectlOutput, fieldName)
	if err != nil {
		return fieldValues, err
	}

	if fieldIndex >= 0 {
		var kubectlOutputLineArr = strings.Split(kubectlOutput, "\n")
		// analyze output and try to find name
		if len(kubectlOutputLineArr) > 0 {
			for lIdx, kubectlOutputLine := range kubectlOutputLineArr {
				if lIdx > lineIndex {
					var kubectlLineFields = strings.Fields(kubectlOutputLine)
					if len(kubectlLineFields) > fieldIndex {
						fieldValues = append(fieldValues, kubectlLineFields[fieldIndex])
					}
				}
			}
		}
	}

	return fieldValues, err
}

// FindFieldIndexInKubectlOutput finds the field index in output.
// it returns also the line index where the field was found.
func FindFieldIndexInKubectlOutput(kubectlOutput string, fieldName string) (lineIndex int, fieldIndex int, err error) {
	var log = logger.Log()
	// set default
	var fieldIdx = -1

	if !strings.Contains(kubectlOutput, fieldName) {
		log.Errorf("[FindFieldIndexInKubectlOutput] Kubectl output does not contain field name [%s]", fieldName)
	} else {
		// split output in array with lines
		var kubectlOutputLineArr = strings.Split(kubectlOutput, "\n")
		// analyze output and try to find name
		if len(kubectlOutputLineArr) > 0 {
			for lIdx, kubectlOutputLine := range kubectlOutputLineArr {
				if strings.Contains(kubectlOutputLine, fieldName) {
					lineIndex = lIdx
					var kubectlLineFields = strings.Fields(kubectlOutputLine)
					fieldIdx = arrays.IndexOfArr(fieldName, kubectlLineFields)
					// found index, abort loop
					if fieldIdx > -1 {
						break
					}
				}
			}
		}
	}

	if fieldIdx == -1 {
		err = fmt.Errorf("[FindFieldIndexInKubectlOutput] Cannot find index for fieldName [%s]", fieldName)
		log.Errorf(err.Error())
	}

	return lineIndex, fieldIdx, err
}

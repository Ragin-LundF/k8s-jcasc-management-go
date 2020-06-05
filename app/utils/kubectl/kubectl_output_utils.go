package kubectl

import (
	"errors"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/arrays"
	"k8s-management-go/app/utils/logger"
	"strings"
)

// function to check if a kubectl output contains a value for a field
func CheckIfKubectlOutputContainsValueForField(kubectlOutput string, fieldName string, searchedValue string) bool {
	found := false

	fieldValues, err := FindFieldValuesInKubectlOutput(kubectlOutput, fieldName)
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

// find field values in kubectl output
func FindFieldValuesInKubectlOutput(kubectlOutput string, fieldName string) (fieldValues []string, err error) {
	// First find the field index for the field name
	lineIndex, fieldIndex, err := FindFieldIndexInKubectlOutput(kubectlOutput, fieldName)
	if err != nil {
		return fieldValues, err
	}

	if fieldIndex >= 0 {
		kubectlOutputLineArr := strings.Split(kubectlOutput, "\n")
		// analyze output and try to find name
		if len(kubectlOutputLineArr) > 0 {
			for lIdx, kubectlOutputLine := range kubectlOutputLineArr {
				if lIdx > lineIndex {
					kubectlLineFields := strings.Fields(kubectlOutputLine)
					if len(kubectlLineFields) > fieldIndex {
						fieldValues = append(fieldValues, kubectlLineFields[fieldIndex])
					}
				}
			}
		}
	}

	return fieldValues, err
}

// find the field index in output.
// it returns also the line index where the field was found.
func FindFieldIndexInKubectlOutput(kubectlOutput string, fieldName string) (lineIndex int, fieldIndex int, err error) {
	log := logger.Log()
	// set default
	fieldIdx := -1

	if !strings.Contains(kubectlOutput, fieldName) {
		log.Error("[FindFieldIndexInKubectlOutput] Kubectl output does not contain field name [" + fieldName + "]")
	} else {
		// split output in array with lines
		kubectlOutputLineArr := strings.Split(kubectlOutput, "\n")
		// analyze output and try to find name
		if len(kubectlOutputLineArr) > 0 {
			for lIdx, kubectlOutputLine := range kubectlOutputLineArr {
				if strings.Contains(kubectlOutputLine, fieldName) {
					lineIndex = lIdx
					kubectlLineFields := strings.Fields(kubectlOutputLine)
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
		err = errors.New(constants.NewLine + "[FindFieldIndexInKubectlOutput] Cannot find index for fieldName [" + fieldName + "]")
		log.Error(err.Error())
	}

	return lineIndex, fieldIdx, err
}

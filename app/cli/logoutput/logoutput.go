package logoutput

import (
	"github.com/manifoldco/promptui"
	"k8s-management-go/app/cli/dialogs"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/utils/logger"
	"strings"
)

type loggingState struct {
	Type    string
	Entry   string
	Details string
}

var loggingStateEntries []loggingState

func AddInfoEntry(message string) {
	loggingStateEntries = append(loggingStateEntries, loggingState{Type: "INFO", Entry: message})
}

func AddInfoEntryAndDetails(message string, details string) {
	loggingStateEntries = append(loggingStateEntries, loggingState{Type: "INFO", Entry: message, Details: details})
}

func AddErrorEntry(message string) {
	loggingStateEntries = append(loggingStateEntries, loggingState{Type: "ERROR", Entry: message})
}

func AddErrorEntryAndDetails(message string, details string) {
	loggingStateEntries = append(loggingStateEntries, loggingState{Type: "ERROR", Entry: message, Details: details})
}

func clearLoggingState() {
	loggingStateEntries = nil
	AddInfoEntry("Navigate or search for log entries.")
}

// Show info and error output as select prompt with search
func DialogShowLogging() {
	log := logger.Log()
	// clear screen
	dialogs.ClearScreen()

	// Template for displaying MenuitemModel
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A4 [{{ .Type | green }}] {{ .Entry | white }}",
		Inactive: "  [{{ .Type | cyan }}] {{ .Entry | red }}",
		Selected: "\U000027A4 [{{ .Type | red | cyan }}] {{ .Entry | red }}",
		Details: `
--------- Log Entry ----------
{{ "Type   :" | faint }}	{{ .Type }}
{{ "Message:" | faint }}	{{ .Entry }}
{{ "Details:" | faint }}
{{.Details}}`,
	}

	// searcher (with "/")
	searcher := func(input string, index int) bool {
		logItem := loggingStateEntries[index]
		logEntry := strings.Replace(strings.ToLower(logItem.Entry), " ", "", -1)
		logType := strings.Replace(strings.ToLower(logItem.Type), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(logEntry, input) || strings.Contains(logType, input)
	}

	prompt := promptui.Select{
		Label:     "Log Output. Press Enter to leave this view",
		Items:     loggingStateEntries,
		Templates: templates,
		Size:      20,
		Searcher:  searcher,
	}

	_, _, err := prompt.Run()

	clearLoggingState()
	if err != nil {
		log.Error(err)
	}
}

// Show info and error output as select prompt with search
func DialogShowInfoAndError(info string, err error) {
	log := logger.Log()
	// clear screen
	dialogs.ClearScreen()

	// Prepare log output for info
	var infoErrArray []loggingState
	infoErrArray = append(infoErrArray, loggingState{Type: "INFO", Entry: "Navigate or search for log entries."})
	infoErrArray = append(infoErrArray, prepareInfo(info)...)
	infoErrArray = append(infoErrArray, prepareErr(err)...)

	// Template for displaying MenuitemModel
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A4 [{{ .Type | green }}] {{ .Entry | white }}",
		Inactive: "  [{{ .Type | cyan }}] {{ .Entry | red }}",
		Selected: "\U000027A4 [{{ .Type | red | cyan }}] {{ .Entry | red }}",
		Details: `
--------- Log Entry ----------
{{ "Type   :" | faint }}	{{ .Type }}
{{ "Content:" | faint }}	{{ .Entry }}`,
	}

	// searcher (with "/")
	searcher := func(input string, index int) bool {
		logItem := infoErrArray[index]
		logEntry := strings.Replace(strings.ToLower(logItem.Entry), " ", "", -1)
		logType := strings.Replace(strings.ToLower(logItem.Type), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(logEntry, input) || strings.Contains(logType, input)
	}

	prompt := promptui.Select{
		Label:     "Log Output. Press Enter to leave this view",
		Items:     infoErrArray,
		Templates: templates,
		Size:      20,
		Searcher:  searcher,
	}

	_, _, err = prompt.Run()

	if err != nil {
		log.Error(err)
	}
}

func prepareInfo(info string) []loggingState {
	var logStruct []loggingState
	if len(info) > 0 {
		if strings.Contains(info, constants.NewLine) {
			infoLineArray := strings.Split(info, constants.NewLine)
			for _, infoLine := range infoLineArray {
				logStruct = append(logStruct, loggingState{Type: "INFO", Entry: infoLine})
			}
		} else {
			logStruct = append(logStruct, loggingState{Type: "INFO", Entry: info})
		}
	}
	return logStruct
}

func prepareErr(err error) []loggingState {
	var logStruct []loggingState
	if err != nil && len(err.Error()) > 0 {
		if strings.Contains(err.Error(), constants.NewLine) {
			errorLineArray := strings.Split(err.Error(), constants.NewLine)
			for _, errorLine := range errorLineArray {
				logStruct = append(logStruct, loggingState{Type: "ERROR", Entry: errorLine})
			}
		} else {
			logStruct = append(logStruct, loggingState{Type: "ERROR", Entry: err.Error()})
		}
	}
	return logStruct
}

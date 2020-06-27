package loggingstate

type LoggingState struct {
	Type    string
	Entry   string
	Details string
}

var loggingStateEntries []LoggingState

func AddInfoEntry(message string) {
	loggingStateEntries = append(loggingStateEntries, LoggingState{Type: "INFO", Entry: message})
}

func AddInfoEntryAndDetails(message string, details string) {
	loggingStateEntries = append(loggingStateEntries, LoggingState{Type: "INFO", Entry: message, Details: details})
}

func AddErrorEntry(message string) {
	loggingStateEntries = append(loggingStateEntries, LoggingState{Type: "ERROR", Entry: message})
}

func AddErrorEntryAndDetails(message string, details string) {
	loggingStateEntries = append(loggingStateEntries, LoggingState{Type: "ERROR", Entry: message, Details: details})
}

func ClearLoggingState() {
	loggingStateEntries = nil
	AddInfoEntry("Navigate or search for log entries.")
}

func GetLoggingStateEntries() []LoggingState {
	return loggingStateEntries
}

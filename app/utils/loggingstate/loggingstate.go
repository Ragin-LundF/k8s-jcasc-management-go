package loggingstate

import (
	"k8s-management-go/app/utils/logger"
)

// LoggingState is responsible for internal logging and defines a type (info/error), an entry (short message) and details
type LoggingState struct {
	Type    string
	Entry   string
	Details string
}

var loggingStateEntries []LoggingState

// AddInfoEntry adds info entries without details to the LoggingState array
func AddInfoEntry(message string) {
	loggingStateEntries = append(loggingStateEntries, LoggingState{Type: "INFO", Entry: message})
}

// AddInfoEntryAndDetails adds info entries with details to the LoggingState array
func AddInfoEntryAndDetails(message string, details string) {
	loggingStateEntries = append(loggingStateEntries, LoggingState{Type: "INFO", Entry: message, Details: details})
}

// AddErrorEntry adds error entries without details to the LoggingState array
func AddErrorEntry(message string) {
	loggingStateEntries = append(loggingStateEntries, LoggingState{Type: "ERROR", Entry: message})
}

// AddErrorEntryAndDetails adds error entries with details to the LoggingState array
func AddErrorEntryAndDetails(message string, details string) {
	loggingStateEntries = append(loggingStateEntries, LoggingState{Type: "ERROR", Entry: message, Details: details})
}

// ClearLoggingState clears the LoggingState array
func ClearLoggingState() {
	// first log everything
	LogLoggingStateEntries()
	loggingStateEntries = nil
}

// GetLoggingStateEntries returns the LoggingState array
func GetLoggingStateEntries() []LoggingState {
	return loggingStateEntries
}

// LogLoggingStateEntries logs the StateEntries to the logfile
func LogLoggingStateEntries() {
	log := logger.Log()
	if loggingStateEntries != nil && cap(loggingStateEntries) > 0 {
		log.Info("---- Output of internal Logging history start ----")

		for _, logEntry := range loggingStateEntries {
			log.Infof("[%s] %s", logEntry.Type, logEntry.Entry)
			if logEntry.Details != "" {
				log.Info("--- Details start ---")
				log.Info(logEntry.Details)
				log.Info("--- Details end ---")
			}
		}

		log.Info("---- Output of internal Logging history end ----")

		// cleanup log to avoid doubles and do it only if it is not empty to avoid loop
		if loggingStateEntries != nil && cap(loggingStateEntries) > 0 {
			loggingStateEntries = nil
		}
	}
}

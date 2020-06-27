package loggingstate

import (
	"k8s-management-go/app/utils/logger"
)

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
	// first log everything
	LogLoggingStateEntries()
	loggingStateEntries = nil
}

func GetLoggingStateEntries() []LoggingState {
	return loggingStateEntries
}

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

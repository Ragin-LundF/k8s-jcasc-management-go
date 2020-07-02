package logger

import (
	"go.uber.org/zap"
)

// LogFilePath is a global variable for the logfile path
var LogFilePath string

// LogEncoding is a global variable for the logfile encoding (json or console)
var LogEncoding string

// Log returns an instance of SugaredLogger to log into a logfile
func Log() *zap.SugaredLogger {
	logConfig := zap.NewProductionConfig()
	if LogFilePath != "" {
		logConfig.OutputPaths = []string{
			LogFilePath,
		}
		if LogEncoding != "" {
			logConfig.Encoding = LogEncoding
		}
	}

	logger, _ := logConfig.Build()
	defer logger.Sync()
	return logger.Sugar()
}

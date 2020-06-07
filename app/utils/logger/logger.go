package logger

import (
	"go.uber.org/zap"
)

var LogFilePath string
var LogEncoding string

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

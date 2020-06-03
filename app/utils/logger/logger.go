package logger

import (
	"go.uber.org/zap"
)

var LogFilePath string

func Log() *zap.SugaredLogger {
	logConfig := zap.NewProductionConfig()
	if LogFilePath != "" {
		logConfig.OutputPaths = []string{
			LogFilePath,
		}
	}

	logger, _ := logConfig.Build()
	defer logger.Sync()
	return logger.Sugar()
}

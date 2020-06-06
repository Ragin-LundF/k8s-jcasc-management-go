package logger

import (
	"go.uber.org/zap"
	"k8s-management-go/app/constants"
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

// small helper for massive logging
func InfoLog(infoLog string, message string) (info string) {
	log := Log()
	log.Info(message)
	info = infoLog + constants.NewLine + message
	return info
}

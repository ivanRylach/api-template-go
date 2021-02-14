package util

import "go.uber.org/zap"

func ConfigureLogging() {
	logger, _ := zap.NewDevelopment()

	zap.ReplaceGlobals(logger)

	defer logger.Sync() // flushes buffer, if any
}

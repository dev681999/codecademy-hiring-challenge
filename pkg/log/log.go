package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is default logger
var Logger *zap.Logger
var L *zap.Logger

func init() {
	New(false)
}

// New sets up logger
func New(debug bool) *zap.Logger {
	var config zap.Config

	if !debug {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ := config.Build()

	zap.ReplaceGlobals(logger)

	Logger = logger
	L = logger

	return logger
}

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// init Starts a new zap enviroment.
func init() {
	var err error
	// Creates a new production enviroment config
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

// Info : Log out a information message
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug : Log out a debug message
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error : Log out a error message
func Error(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func InitLogger(logPath, logLevel string) error {

	// 因为zap不支持日志归档，所以lumberjack的hook用来归档日志
	hook := lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1024,
		MaxAge:     7,
		MaxBackups: 3,
		Compress:   true,
	}

	w := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch logLevel {
	case "DEBUG":
		level = zap.DebugLevel
	case "INFO":
		level = zap.InfoLevel
	case "ERROR":
		level = zap.ErrorLevel
	default:
		level = zap.DebugLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level)

	logger = zap.New(core)

	return nil
}

func GetLogger() *zap.Logger {
	return logger
}
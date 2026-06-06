package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func InitLogger(env string, debug bool) error {
	var config zap.Config

	if debug {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
	}

	var err error
	Logger, err = config.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	SugarLogger = Logger.Sugar()
	
	defer Logger.Sync()

	return nil
}

func Info(msg string, fields ...interface{}) {
	if SugarLogger != nil {
		SugarLogger.Infow(msg, fields...)
	}
}

func Error(msg string, fields ...interface{}) {
	if SugarLogger != nil {
		SugarLogger.Errorw(msg, fields...)
	}
}

func Debug(msg string, fields ...interface{}) {
	if SugarLogger != nil {
		SugarLogger.Debugw(msg, fields...)
	}
}

func Warn(msg string, fields ...interface{}) {
	if SugarLogger != nil {
		SugarLogger.Warnw(msg, fields...)
	}
}

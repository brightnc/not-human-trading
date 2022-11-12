package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type HTTPLog struct {
	IP,
	Method,
	Path,
	RequestBody,
	Query,
	ResponseBody string
	Status int
}

var (
	zapLog *zap.SugaredLogger
)

// Init ...
// initial log
func Init() {
	var err error
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.StacktraceKey = ""
	defaultLog, err := loggerConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalln("Got error while building zap logger config.")
		return
	}
	zapLog = defaultLog.Sugar()
	return
}

// HTTP ...
// log http information
func HTTP(info HTTPLog) {
	zapLog.Infow(fmt.Sprintf("request method %s path %s", info.Method, info.Path),
		"IP", info.IP,
		"method", info.Method,
		"path", info.Path,
		"requestBody", info.RequestBody,
		"query", info.Query,
		"responseBody", info.ResponseBody,
	)
	defer zapLog.Sync()
}

func Infof(format string, args ...interface{}) {
	zapLog.Infof(format, args...)
	defer zapLog.Sync()
}

func Info(args ...interface{}) {
	zapLog.Info(args)
	defer zapLog.Sync()
}

func Fatalf(format string, args ...interface{}) {
	zapLog.Fatalf(format, args)
	defer zapLog.Sync()
}
func Error(args ...interface{}) {
	zapLog.Error(args)
	defer zapLog.Sync()
}

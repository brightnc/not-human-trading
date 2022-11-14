package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLog *zap.SugaredLogger
)

// Init ...
// initial log
func Init(component string) {
	var err error
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.StacktraceKey = ""
	loggerConfig.InitialFields = map[string]interface{}{
		"component": component,
	}
	defaultLog, err := loggerConfig.Build(zap.AddCallerSkip(1))

	if err != nil {
		log.Fatalln("Got error while building zap logger config.")
		return
	}
	zapLog = defaultLog.Sugar()
	return
}

// Infof ...
// Infof uses fmt.Sprintf to log a templated message.
func Infof(format string, args ...interface{}) {
	zapLog.Infof(format, args...)
	defer zapLog.Sync()
}

// Info ...
// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	zapLog.Info(args)
	defer zapLog.Sync()
}

// InfoWithCorrelationID ...
func InfoWithCorrelationID(msg, correlationID string) {
	zapLog.Infow(msg, "correlationID", correlationID)
	defer zapLog.Sync()
}

// ErrorWithCorrelationID ...
func ErrorWithCorrelationID(msg, correlationID string) {
	zapLog.Errorw(msg, "correlationID", correlationID)
	defer zapLog.Sync()
}

// Fatalf ...
// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(format string, args ...interface{}) {
	zapLog.Fatalf(format, args)
	defer zapLog.Sync()
}

// Error ...
// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	zapLog.Error(args)
	defer zapLog.Sync()
}

// Errorf ...
// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	zapLog.Errorf(template, args)
	defer zapLog.Sync()
}

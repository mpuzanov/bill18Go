package logger

import (
	"os"

	"github.com/mpuzanov/bill18Go/config"

	"github.com/sirupsen/logrus"
)

//Logger ...
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	//Logger.Formatter = &logrus.TextFormatter{DisableColors: true}
	Logger.Formatter = new(logrus.TextFormatter)
	Logger.Formatter.(*logrus.TextFormatter).TimestampFormat = "02-01-2006 15:04:05"
	Logger.Formatter.(*logrus.TextFormatter).FullTimestamp = true
}

//SetupLogger ...
func SetupLogger(cfg *config.Config) error {
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		Logger.Error(err)
	}
	Logger.SetLevel(level)

	if cfg.LogToFile {
		file, err := os.OpenFile(cfg.LogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			Logger.Printf("logToFile: %s  http-server: %s", cfg.LogFileName, cfg.Listen)
			Logger.Out = file
		} else {
			Logger.Info("Failed to log to file, using default stderr")
		}
	}
	Logger.Println("log.Level:", Logger.Level)
	return nil
}

// Debug logs a debug message
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// Traceln logs a debug message
func Traceln(args ...interface{}) {
	Logger.Traceln(args...)
}

// Debugf logs a formatted debug messsage
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Info logs an informational message
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Infof logs a formatted informational message
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Error logs an error message
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Errorf logs a formatted error message
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Warn logs a warning message
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Warnf logs a formatted warning message
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Fatal logs a fatal error message
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Fatalf logs a formatted fatal error message
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

// WithFields returns a new log enty with the provided fields
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Logger.WithFields(fields)
}

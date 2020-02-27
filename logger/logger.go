package logger

import (
	"github.com/sirupsen/logrus"
)

// The Logger interface
type Interface interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// The Logger structure
type Logger struct {
	logger *logrus.Logger
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Returns a new logger object
func NewLogger(isDebug bool) *Logger {
	// set debug level and the output file
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	if isDebug {
		log.SetLevel(logrus.DebugLevel)
	}
	return &Logger{logger: log}
}

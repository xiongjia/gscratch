package logger

import (
	logrus "github.com/Sirupsen/logrus"
	"path/filepath"
	"runtime"
)

type LogLevel uint8

const (
	ErrorLevel LogLevel = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

var (
	curLevel LogLevel = InfoLevel
)

func SetLevel(level LogLevel) {
	curLevel = level
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
}

type Logger struct {
	level LogLevel
	entry *logrus.Entry
}

func New(tag string) *Logger {
	_, srcFile, _, _ := runtime.Caller(0)
	srcFileBase := filepath.Base(srcFile)
	return &Logger{
		entry: logrus.WithFields(logrus.Fields{"tag": tag, "src": srcFileBase}),
		level: InfoLevel,
	}
}

func (logger *Logger) Errorf(fmt string, args ...interface{}) {
	logger.entry.Errorf(fmt, args...)
}

func (logger *Logger) Warnf(fmt string, args ...interface{}) {
	if curLevel >= WarnLevel {
		logger.entry.Warnf(fmt, args...)
	}
}

func (logger *Logger) Infof(fmt string, args ...interface{}) {
	if curLevel >= InfoLevel {
		logger.entry.Infof(fmt, args...)
	}
}

func (logger *Logger) Debugf(fmt string, args ...interface{}) {
	if curLevel >= DebugLevel {
		logger.entry.Debugf(fmt, args...)
	}
}

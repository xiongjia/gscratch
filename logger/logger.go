package logger

import (
	log "github.com/Sirupsen/logrus"
)

type LogLevel uint8

const (
	ErrorLevel LogLevel = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

type Logger struct {
	level LogLevel
	entry *log.Entry
}

func New(section string) *Logger {
	return &Logger{
		entry: log.WithFields(log.Fields{"section": section}),
		level: InfoLevel,
	}
}

func (logger *Logger) Errorf(fmt string, args ...interface{}) {
	logger.entry.Errorf(fmt, args...)
}

func (logger *Logger) Warnf(fmt string, args ...interface{}) {
	logger.entry.Warnf(fmt, args...)
}

func (logger *Logger) Infof(fmt string, args ...interface{}) {
	logger.entry.Infof(fmt, args...)
}

func (logger *Logger) Debugf(fmt string, args ...interface{}) {
	logger.entry.Debugf(fmt, args...)
}

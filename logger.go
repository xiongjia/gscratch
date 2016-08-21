package gscratch

import (
	log "github.com/Sirupsen/logrus"
)

type Logger struct {
}

func LogCreate() *Logger {
	return &Logger{}
}

func (logger *Logger) Infof(fmt string, args ...interface{}) {
	log.Infof(fmt, args...)
}

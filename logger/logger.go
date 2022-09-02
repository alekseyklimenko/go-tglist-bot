package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Get() *logrus.Logger {
	return log
}

func NewEntry() *logrus.Entry {
	return logrus.NewEntry(log)
}

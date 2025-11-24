package logger

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetLevel(logrus.InfoLevel)
	Log.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	Log.Info("Logger initialized!")
}

func Init() {
	Log.Info("Logger online.")
}

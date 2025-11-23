package logger

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// init викликається автоматично при імпорті пакету
func init() {
	Log = logrus.New()
	Log.SetLevel(logrus.InfoLevel)
	Log.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	Log.Info("Logger initialized!")
}

// Init дозволяє повторно ініціалізувати логер, якщо потрібно
func Init() {
	Log.Info("Logger online.")
}

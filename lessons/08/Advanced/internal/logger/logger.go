package logger

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// Функція init ініціалізовує логер
func init() {
	Log = logrus.New()
	Log.SetLevel(logrus.InfoLevel)
	Log.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	Log.Info("Logger initialized!")
}

// Функція Init запускається лише один раз для модулю
// Повідомляє, що модуль є активним
func Init() {
	Log.Info("Logger online.")
}

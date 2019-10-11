
package log

import (
	"github.com/sirupsen/logrus"
	"github.com/mattn/go-colorable"
)

var levels = []string{
	"debug",
	"info",
	"warn",
	"error",
	"fatal",
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetOutput(colorable.NewColorableStdout())
}

func Warn(msg ...interface{}) {
	logrus.Warn(msg)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args)
}
func Debug(msg ...interface{}) {
	logrus.Debug(msg)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(msg ...interface{}) {
	logrus.Info(msg)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Error(msg ...interface{}) {
	logrus.Error(msg)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatal(msg ...interface{}) {
	logrus.Fatal(msg)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func SetLevel(level string) {
	logrusLevel, err := logrus.ParseLevel(level)

	if err != nil {
		logrus.Fatalf("\"%s\" is not a valid logging level", level)
	}

	logrus.SetLevel(logrusLevel)
}

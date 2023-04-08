package log

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()

func init() {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true

	// Set specific colors for prefix and timestamp
	formatter.SetColorScheme(&prefixed.ColorScheme{
		PrefixStyle:    "blue+b",
		TimestampStyle: "white+h",
	})

	log.Formatter = formatter
	log.Level = logrus.DebugLevel
}

func NewLogger(instance string, removeNewLine bool) *logrus.Entry {
	newLogger := log
	entry := newLogger.WithFields(logrus.Fields{
		"prefix": instance,
	})

	if removeNewLine {
		if len(entry.Message) > 0 && entry.Message[len(entry.Message)-1] == '\n' {
			b := []byte(entry.Message[0 : len(entry.Message)-1])
			entry.Message = string(b)
		}
		newLogger.Formatter.Format(entry)
	}

	return entry
}

func Event(instance, msg string) {
	if instance != "" {
		log.WithFields(logrus.Fields{
			"prefix": instance,
		}).Info(msg)
	} else {
		log.Info(msg)
	}
}

func Fatal(instance, msg string, err error) {
	log.WithFields(logrus.Fields{
		"prefix": instance,
	}).Fatal(msg, ": ", err.Error())
}

func Error(instance, msg string, err error) {
	log.WithFields(logrus.Fields{
		"prefix": instance,
	}).Errorln(msg, ": ", err.Error())
}

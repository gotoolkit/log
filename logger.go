package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	std = New()
)

const (
	defaultLevel = DebugLevel
)

type Logger struct {
	ILogger
}

func New(opts ...Option) *Logger {
	options := options{
		formatter: &logrus.JSONFormatter{},
		level:     defaultLevel,
		out:       os.Stderr,
	}

	for _, o := range opts {
		o.apply(&options)
	}
	l := logrus.New()
	l.SetFormatter(options.formatter)
	l.SetLevel(logrus.Level(options.level))
	l.SetOutput(options.out)
	return &Logger{
		ILogger: l,
	}
}

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
	*logrus.Logger
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
	if options.formatter != nil {
		l.SetFormatter(options.formatter)
	}

	l.SetOutput(options.out)
	l.SetLevel(logrus.Level(options.level))
	return &Logger{
		Logger: l,
	}
}

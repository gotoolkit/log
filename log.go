package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var ()

const (
	defaultLevel = DebugLevel
)

func Setup(opts ...Option) error {

	options := options{
		formatter: &logrus.JSONFormatter{},
		level:     defaultLevel,
		out:       os.Stderr,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	if options.formatter != nil {
		logrus.SetFormatter(options.formatter)
	}

	logrus.SetOutput(options.out)
	logrus.SetLevel(logrus.Level(options.level))

	return nil
}

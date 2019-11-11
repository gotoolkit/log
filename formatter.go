package log

import (
	"github.com/sirupsen/logrus"
)

type Formatter interface {
	logrus.Formatter
}

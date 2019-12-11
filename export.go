package log

import (
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetFormatter(formatter Formatter) {
	std.SetFormatter(formatter)
}

func SetLevel(level Level) {
	std.SetLevel(logrus.Level(level))
}

func SetOutput(output io.Writer) {
	std.SetOutput(output)
}
func SetOutputFile(filename string, maxSizeMB, maxBackups, maxDays int, compress bool) {
	std.SetOutput(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSizeMB, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxDays,  //days
		Compress:   compress, // disabled by default
	})
}

func AddHook(hook Hook) {
	std.AddHook(hook)
}

func WithContext(ctx context.Context) *logrus.Entry {
	return std.WithContext(ctx)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return std.WithField(key, value)
}

func WithFields(fields map[string]interface{}) *logrus.Entry {
	return std.WithFields(fields)
}

func WithTime(t time.Time) *logrus.Entry {
	return std.WithTime(t)
}

func Trace(args ...interface{}) {
	std.Trace(args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Print(args ...interface{}) {
	std.Print(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Warning(args ...interface{}) {
	std.Warning(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	std.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func Traceln(args ...interface{}) {
	std.Traceln(args...)
}

func Debugln(args ...interface{}) {
	std.Debugln(args...)
}

func Println(args ...interface{}) {
	std.Println(args...)
}

func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

func Warnln(args ...interface{}) {
	std.Warnln(args...)
}

func Warningln(args ...interface{}) {
	std.Warningln(args...)
}

func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}

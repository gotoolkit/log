package log

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

type options struct {
	level     uint32
	out       io.Writer
	formatter Formatter
	hooks     []Hook
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithOutputFile(filename string, maxSize, maxBackups, maxAge int, compress bool) Option {
	return optionFunc(func(o *options) {
		o.out = &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    maxSize, // megabytes
			MaxBackups: maxBackups,
			MaxAge:     maxAge,   //days
			Compress:   compress, // disabled by default
		}
	})
}

func WithOutput(out io.Writer) Option {
	return optionFunc(func(o *options) {
		o.out = out
	})
}

func WithFormatter(formatter Formatter) Option {
	return optionFunc(func(o *options) {
		o.formatter = formatter
	})
}

func WithLevel(level uint32) Option {
	return optionFunc(func(o *options) {
		o.level = level
	})
}

func WithHooks(hook Hook) Option {
	return optionFunc(func(o *options) {
		o.hooks = append(o.hooks, hook)
	})
}
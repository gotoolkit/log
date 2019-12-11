package log

type Level uint32

const (
	PanicLevel uint32 = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

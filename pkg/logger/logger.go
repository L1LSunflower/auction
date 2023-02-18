package logger

import (
	"github.com/L1LSunflower/auction/pkg/logger/drivers"
	"github.com/L1LSunflower/auction/pkg/logger/message"
)

const (
	DriverStdout  = "stdout"
	DriverGraylog = "graylog"
)

var Log Interface

// Interface is a main interface for all app loggers
type Interface interface {
	Debug(msg *message.LogMessage)
	Info(msg *message.LogMessage)
	Warn(msg *message.LogMessage)
	Error(msg *message.LogMessage)
	Fatal(msg *message.LogMessage)
	Panic(msg *message.LogMessage)
}

func New(level string, driver string, addr string) Interface {
	switch driver {
	case DriverStdout:
		Log = drivers.MakeStdoutLogger(level)
		return Log
	default:
		Log = drivers.MakeStdoutLogger(level)
		return Log
	}
}

package logger

import (
	"github.com/L1LSunflower/auction/pkg/logger/drivers"
)

const (
	DriverStdout = "stdout"
	DriverFile   = "file"
)

var Log drivers.LogInterface

func New(level string, driver string) drivers.LogInterface {
	switch driver {
	case DriverStdout:
		Log = drivers.MakeStdoutLogger(level)
		return Log
	case DriverFile:
		Log = drivers.MakeFileLogger(level)
		return Log
	default:
		Log = drivers.MakeStdoutLogger(level)
		return Log
	}
}

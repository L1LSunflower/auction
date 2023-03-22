package drivers

import "github.com/L1LSunflower/auction/pkg/logger/message"

// LogInterface is a main interface for all app loggers
type LogInterface interface {
	Debug(msg *message.LogMessage)
	Info(msg *message.LogMessage)
	Warn(msg *message.LogMessage)
	Error(msg *message.LogMessage)
	Fatal(msg *message.LogMessage)
	Panic(msg *message.LogMessage)
	Close()
}

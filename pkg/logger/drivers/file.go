package drivers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/L1LSunflower/auction/pkg/logger/message"
)

type FileDriver struct {
	log   *logrus.Logger
	level logrus.Level
	file  *os.File
}

func (f *FileDriver) Debug(msg *message.LogMessage) {
	f.write(logrus.DebugLevel, msg)
}

func (f *FileDriver) Info(msg *message.LogMessage) {
	f.write(logrus.InfoLevel, msg)
}

func (f *FileDriver) Warn(msg *message.LogMessage) {
	f.write(logrus.WarnLevel, msg)
}

func (f *FileDriver) Error(msg *message.LogMessage) {
	f.write(logrus.ErrorLevel, msg)
}

func (f *FileDriver) Fatal(msg *message.LogMessage) {
	f.write(logrus.FatalLevel, msg)
}

func (f *FileDriver) Panic(msg *message.LogMessage) {
	f.write(logrus.PanicLevel, msg)
}

func (f *FileDriver) Close() {
	f.file.Close()
}

func (f *FileDriver) write(level logrus.Level, msg *message.LogMessage) {
	j, err := json.Marshal(msg)

	if err != nil {
		return
	}

	f.log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: false})
	f.log.Log(level, string(j))
}

// MakeFileLogger creates logrus instance and sets log level
func MakeFileLogger(level string) LogInterface {
	var lev logrus.Level

	switch level {
	case message.TraceLevel:
		lev = logrus.TraceLevel
		break
	case message.DebugLevel:
		lev = logrus.DebugLevel
		break
	case message.InfoLevel:
		lev = logrus.InfoLevel
		break
	case message.WarnLevel:
		lev = logrus.WarnLevel
		break
	case message.ErrorLevel:
		lev = logrus.ErrorLevel
		break
	case message.FatalLevel:
		lev = logrus.FatalLevel
		break
	case message.PanicLevel:
		lev = logrus.PanicLevel
		break
	default:
		lev = logrus.WarnLevel
	}

	f, err := os.Create("./app.log")
	if err != nil {
		fmt.Printf("failed to create log file with error: %s\n", err.Error())
	}

	f, err = os.OpenFile("./app.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Failed to create logfile" + "app.log")
		panic(err)
	}

	l := logrus.New()
	l.SetLevel(lev)
	//l.Out = io.MultiWriter(f, os.Stdout)
	mw := io.MultiWriter(f)
	l.SetOutput(mw)

	return &FileDriver{log: l, level: lev, file: f}
}

package log

import (
	"os"
	"runtime"
)

type ilog interface {
	New()
	Info(string)
	Debug(string)
	Warning(string)
	Error(string)
	Fail(string)
}

var instace ilog

type LogType string

const (
	console LogType = "console"
	gcp     LogType = "gcp"
)

func New(t LogType) {
	instace = &Console{}
	switch t {
	case gcp:
		instace = &GCP{}
	case console:
		instace = &Console{}
	}
	instace.New()
}

func Info(msg string) {
	instace.Info(msg)
}

func Debug(msg string) {
	instace.Debug(msg)
}

func Warning(msg string) {
	instace.Warning(msg)
}

func Error(msg string) {
	stackSlice := make([]byte, 512)
	s := runtime.Stack(stackSlice, false)
	instace.Error(string(stackSlice[0:s]))
	instace.Error(msg)
}

func Fail(msg string) {
	stackSlice := make([]byte, 512)
	s := runtime.Stack(stackSlice, false)
	instace.Error(string(stackSlice[0:s]))
	instace.Fail(msg)
	os.Exit(0)
}

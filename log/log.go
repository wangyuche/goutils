package log

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
	switch t {
	case gcp:
		instace = &GCP{}
		instace.New()
		return
	case console:
		instace = &Console{}
		instace.New()
		return
	default:
		panic("Log Type Error")
	}
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
	instace.Error(msg)
}

func Fail(msg string) {
	instace.Fail(msg)
}

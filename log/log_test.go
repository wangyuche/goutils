package log

import (
	"os"
	"testing"
)

func TestConsole(t *testing.T) {
	//os.Setenv("LogShowTime", "True")
	os.Setenv("LogShowTime", "False")
	New(LogType(console))
	Info("msg Info")
	Debug("msg Debug")
	Warning("msg Warning")
	Error("msg Error")
	Fail("msg Fail")
}

func TestGCP(t *testing.T) {
	os.Setenv("GCP_PROJECTID", "op-lab")
	os.Setenv("HOSTNAME", "aries-test")
	New(LogType(gcp))
	Info("msg Info")
	Debug("msg Debug")
	Warning("msg Warning")
	Error("msg Error")
	Fail("msg Fail")
}

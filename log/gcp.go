package log

import (
	"context"
	"os"
	"strings"

	"cloud.google.com/go/logging"
)

type GCP struct {
	Logger *logging.Logger
}

func (this *GCP) New() {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, os.Getenv("GCP_PROJECTID"))
	if err != nil {
		panic("Failed to create client")
	}
	name := strings.Split(os.Getenv("HOSTNAME"), "-")
	logName := "golanglog"
	if len(name) > 1 {
		logName = name[0]
	}
	this.Logger = client.Logger(logName)
}

func (this *GCP) Info(msg string) {
	this.Logger.Log(logging.Entry{
		Payload:  msg,
		Severity: logging.Info,
	})
	this.Logger.Flush()
}

func (this *GCP) Debug(msg string) {
	this.Logger.Log(logging.Entry{
		Payload:  msg,
		Severity: logging.Debug,
	})
	this.Logger.Flush()
}

func (this *GCP) Warning(msg string) {
	this.Logger.Log(logging.Entry{
		Payload:  msg,
		Severity: logging.Warning,
	})
	this.Logger.Flush()
}

func (this *GCP) Error(msg string) {
	this.Logger.Log(logging.Entry{
		Payload:  msg,
		Severity: logging.Error,
	})
	this.Logger.Flush()
}

func (this *GCP) Fail(msg string) {
	this.Logger.Log(logging.Entry{
		Payload:  msg,
		Severity: logging.Alert,
	})
	this.Logger.Flush()
}

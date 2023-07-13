package log

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Console struct {
	ShowTime bool
}

func (this *Console) New() {
	b, _ := strconv.ParseBool(os.Getenv("LogShowTime"))
	this.ShowTime = b
}

func (this *Console) Info(msg string) {
	if this.ShowTime {
		fmt.Printf("%-9s %-19s %s\n", "[Info]", time.Now().Format("2006-01-02 15:04:05"), msg)
	} else {
		fmt.Printf("%-9s %s\n", "[Info]", msg)
	}

}

func (this *Console) Debug(msg string) {
	if this.ShowTime {
		fmt.Printf("%-9s %-19s %s\n", "[Debug]", time.Now().Format("2006-01-02 15:04:05"), msg)
	} else {
		fmt.Printf("%-9s %s\n", "[Debug]", msg)
	}
}

func (this *Console) Warning(msg string) {
	if this.ShowTime {
		fmt.Printf("%-9s %-19s %s\n", "[Warning]", time.Now().Format("2006-01-02 15:04:05"), msg)
	} else {
		fmt.Printf("%-9s %s\n", "[Warning]", msg)
	}
}

func (this *Console) Error(msg string) {
	if this.ShowTime {
		fmt.Printf("%-9s %-19s %s\n", "[Error]", time.Now().Format("2006-01-02 15:04:05"), msg)
	} else {
		fmt.Printf("%-9s %s\n", "[Error]", msg)
	}
}

func (this *Console) Fail(msg string) {
	if this.ShowTime {
		fmt.Printf("%-9s %-19s %s\n", "[Fail]", time.Now().Format("2006-01-02 15:04:05"), msg)
	} else {
		fmt.Printf("%-9s %s\n", "[Fail]", msg)
	}
	os.Exit(0)
}

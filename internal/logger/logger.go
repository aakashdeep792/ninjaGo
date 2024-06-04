package logger

import (
	"fmt"
	"sync"
)

var o sync.Once
var logLevel string

const (
	INFO  = "INFO"
	DEBUG = "DEBUG"
)

func UpdateLogLevel(level string) {
	o.Do(func() {
		logLevel = level
	})
}

func Log(a ...any) {
	fmt.Println(a...)
}

func Logf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func Debug(a ...any) {
	if logLevel == DEBUG {
		fmt.Println(a...)
	}
}

func Debugf(format string, a ...any) {
	if logLevel == DEBUG {
		fmt.Printf(format+"\n", a...)
	}
}

func Errorf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func Error(a ...any) {
	fmt.Println(a...)
}

func Warn(a ...any) {
	fmt.Println(a...)
}

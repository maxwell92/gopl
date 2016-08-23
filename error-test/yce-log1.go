package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel uintptr

const (
	FATAL LogLevel = 1
	ERROR LogLevel = 2
	WARN  LogLevel = 3
	INFO  LogLevel = 4
	DEBUG LogLevel = 5
	TRACE LogLevel = 6
)

var loglevels = [...]string{
	FATAL: "FATAL",
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
	TRACE: "TRACE",
}

type YceLogger struct {
	Logger *log.Logger
}

func New() *YceLogger {
	return &YceLogger{
		Logger: log.New(os.Stderr, "", log.Lshortfile),
	}
}

func (l LogLevel) Log() string {
	return loglevels[l]
}

func (y *YceLogger) Printf(level LogLevel, funcName, context string) {
	y.Logger.SetPrefix(fmt.Sprintf("%s [%s] ", time.Now().Format("2006-01-02 15:04:05.000"), level.Log()))
	y.Logger.Printf("%s: %s\n", funcName, context)
}

func (y *YceLogger) Fatalf(level LogLevel, funcName, context string) {
	y.Logger.SetPrefix(fmt.Sprintf("%s [%s] ", time.Now().Format("2006-01-02 15:04:05.000"), level.Log()))
	y.Logger.Fatalf("%s: %s\n", funcName, context)

}

func main() {
	y := New()
	y.Printf(DEBUG, "main()", "main debug log")
	y.Fatalf(FATAL, "main()", "main fatal error")
}

package main

import (
	"log"
	"os"
)

type mylogger struct {
	l *log.Logger
}

func New() *mylogger {
	l := log.New(os.Stderr, "", log.Lshortfile)
	return &mylogger{
		l: l,
	}
}

func main() {
	ml := New()
	ml.l.Println("Log")
}

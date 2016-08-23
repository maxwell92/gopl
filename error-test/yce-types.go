package main

import (
	"encoding/json"
	"error"
	"log"
)

// YceLogger

type YceLogger struct {
	*log.Logger
}

func (l *YceLogger) Printf(format string, args ...interface{}) {

}

func (l *YceLogger) Fatalf(format string, args ...interface{}) {

}

// YceError

type Errno uintptr

const (
	ETIMEOUT Errno = 1 /*Operation time out*/

)

var errors = [...]string{
	ETIMEOUT: "Operation time out",
}

type ResponseError struct {
	// need format convertion from uintptr to int
	Code    Errno  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type error interface {
	Error() string
	EncodeSelf() []byte
}

type YceError struct {
	errmsg string
}

//TODO: standard err convert to YceError
func New(text string) *YceError {
	return &YceError{errmsg: text}
}

func (e *YceError) Error() string {
	return e.errmsg
}

func (e *YceError) EncodeSelf() []byte {
	errJson, _ := json.Marshal(e.errmsg)
	return erJson
}

func (n Errno) Error() string {
	if 0 <= int(n) && int(n) < len(errors) {
		return errors[n]
	}
	//TODO: Errno out of defination range
	return "Error Never Seen"
}

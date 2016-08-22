package main

import (
	"fmt"
)

type myerror interface {
	Error() string
}

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	text string
}

func (e *errorString) Error() string {
	return e.text
}

func Errorf(format string, args ...interface{}) myerror {
	str := Sprintf(format, args...)
	return New(str)
}

func main() {
	msg := "this is an error"
	myerr := Errorf("%s\n", msg)
	fmt.Println(myerr.text)
}

package main

import (
	"log"
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

func main() {

	log.Printf("%v\n", New("this is error").Error())
	log.Println(New("this is error").Error())

	myerr := New("myerror is error")
	errmsg := myerr.Error()
	log.Println(errmsg)

	errmsg2 := New("myerror is error2").Error()
	log.Println(errmsg2)
}

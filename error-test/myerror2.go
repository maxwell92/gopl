package main

import (
	"fmt"
)

type error interface {
	Error() string
}

type myerror struct {
	text string
}

func New(text string) *myerror {
	return &myerror{text: text}
}

func (err *myerror) Error() string {
	return err.text
}

func Errorf(format string, arg ...interface{}) *myerror {
	return New(fmt.Sprintf(format, arg...))

}

func main() {
	err := New("abc error")
	fmt.Println(err.Error())
	fmt.Println(err.text)
	fmt.Println(Errorf("this is error").Error())
}

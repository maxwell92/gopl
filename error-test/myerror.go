package main

import (
	"fmt"
	//	"log"
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

// different method with interface doesn't work
/*
func (e errorString) Errornormal() string {
	return e.text
}

func (e *errorString) Errorpointer() string {
	return e.text
}
*/
func main() {
	/*
		log.Printf("%v\n", New("this is error").Error())
		log.Println(New("this is error").Error())

		myerr := New("myerror is error")
		errmsg := myerr.Error()
		log.Println(errmsg)

		errmsg2 := New("myerror is error2").Error()
		log.Println(errmsg2)

		fmt.Println(New("EOF") == New("EOF"))
	*/

	// no matter it's pointer receiver for normal receiver,
	// no same error text address
	/*
		fmt.Println("pointer")
		for i := 0; i < 5; i++ {
			msg := New("EOF").Error()
			fmt.Printf("%p\n", &msg)
		}

		fmt.Println("normal")
		for i := 0; i < 5; i++ {
			msg := New("EOF").Error()
			fmt.Printf("%p\n", &msg)
		}
	*/
}

package main

import (
	"fmt"
	"syscall"
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

func (e syscall.Errno) Error() string {
	if 0 <= int(e) && int(e) < len(syscall.errors) {
		return syscall.errors[e]
	}
	return fmt.Sprintf("errno %d", e)
}

func main() {
	var err myerror = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}

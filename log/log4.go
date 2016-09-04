package main

import (
	"fmt"
	"log"
)

type MyWriter struct {
	message string
}

func (m MyWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return 0, nil
}

func main() {
	w := &MyWriter{
		message: "Hello",
	}
	log.SetOutput(w)
	log.Println("Hello World from log pkg")
}

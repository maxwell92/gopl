package main

import "log"

type A struct {
	value string
}

func (a *A) Test() string {
	log.Print("Test Called")
	return a.value
}

func getA() *A {
	return nil
}

func main() {
	a := getA()
	a.Test()
}

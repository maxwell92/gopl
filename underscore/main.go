package main

import (
	"fmt"
)

type Object interface {
	Get()
}

type Car struct {
	Name string
}

func (c *Car) Get() {

}

var _ Object = &Car{}

func main() {
	fmt.Println("Hello")
}

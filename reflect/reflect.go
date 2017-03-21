package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	Name  string
	Price string
}

func main() {
	a := 123
	b := "abc"
	c := []int32{1, 2, 3}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Printf("%v\n", reflect.TypeOf(c).String() == "[]int32")

	benz := &Car{
		Name:  "G65",
		Price: "$ 35,000",
	}

	fmt.Println(reflect.TypeOf(benz))
	fmt.Printf("%v\n", reflect.TypeOf(benz).String() == "Car")
	fmt.Printf("%v\n", reflect.TypeOf(benz).String() == "*Car")
	fmt.Printf("%v\n", reflect.TypeOf(benz).String() == "main.Car")
	fmt.Printf("%v\n", reflect.TypeOf(benz).String() == "*main.Car")
}

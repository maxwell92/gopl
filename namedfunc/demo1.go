package main

import (
	"fmt"
	"reflect"
)

type Arith func(a, b int) int

func main() {
	// 3, main.Arith
	var add Arith
	add = func(a, b int) int {
		return a + b
	}

	fmt.Printf("%d, %s\n", add(1, 2), reflect.TypeOf(add).String())

	// 0, func(int, int) int
	divide := func(a, b int) int {
		return a / b
	}
	fmt.Printf("%d, %s\n", divide(1, 2), reflect.TypeOf(divide).String())
}

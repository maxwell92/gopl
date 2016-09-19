package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 123
	b := "abc"
	c := []int32{1, 2, 3}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Printf("%v\n", reflect.TypeOf(c).String() == "[]int32")
}

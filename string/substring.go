package main

import (
	"fmt"
	//"string"
)

func main() {
	str := "Hello World"
	l := len(str)
	fmt.Printf("str[%d]: %v\n", 0, str[0])
	fmt.Printf("str[%d]: %v\n", l-1, str[l-1])

	fmt.Printf("str[%d:%d]: %s\n", 0, 3, str[0:3])
}

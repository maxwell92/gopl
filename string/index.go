package main

import (
	"fmt"
)

func main() {
	str := "healthz"
	fmt.Println(string(str[0]))

	/*
		str = string(append([]byte(str), byte(0)))
		str[0] = byte('/')
		fmt.Println(str)
	*/

	l := len(str)
	str = str + "/" + str
	fmt.Println(str)
	str = str[l:]
	fmt.Println(str)
}

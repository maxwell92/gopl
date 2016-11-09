package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	fmt.Printf("%v\n", str == "")

	fmt.Println(strings.EqualFold(str, "Hello World"))
}

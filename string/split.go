package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "#hello"
	str2 := " #world"
	str3 := "hello # world"

	fmt.Printf("str1: %d, %s\n", len(strings.Split(str1, "#")), strings.Split(str1, "#")[0])
	fmt.Printf("str2: %d, %s\n", len(strings.Split(str2, "#")), strings.Split(str2, "#")[0])
	fmt.Printf("str3: %d, %s\n", len(strings.Split(str3, "#")), strings.Split(str3, "#")[0])
}

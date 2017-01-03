package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "#hello"
	str2 := " #world"
	str3 := "hello # world"
	str4 := " "
	str5 := "    "
	str6 := " a"
	str7 := "a "

	fmt.Printf("str1: %d, %s\n", len(strings.Split(str1, "#")), strings.Split(str1, "#")[0])
	fmt.Printf("str2: %d, %s\n", len(strings.Split(str2, "#")), strings.Split(str2, "#")[0])
	fmt.Printf("str3: %d, %s\n", len(strings.Split(str3, "#")), strings.Split(str3, "#")[0])
	fmt.Printf("str4: %d, %s\n", len(strings.Split(str4, "#")), strings.Split(str4, "#")[0])
	fmt.Printf("str5: %d, %s\n", len(strings.Split(str5, "#")), strings.Split(str5, "#")[0])
	fmt.Printf("str6: %d, %s\n", len(strings.Split(str6, "#")), strings.Split(str6, "#")[0])
	fmt.Printf("str7: %d, %s\n", len(strings.Split(str7, "#")), strings.Split(str7, "#")[0])
}

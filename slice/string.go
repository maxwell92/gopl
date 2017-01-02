package main

import (
	"fmt"
	"strings"
)

func main() {
	str := make([]string, 0)
	fmt.Println(len(str))

	str = append(str, "")
	fmt.Println(len(str))

	str1 := strings.Join([]string{"Hello", "Mushroom"}, "Strawberry")
	fmt.Println(str1)
	fmt.Printf("%v, %d\n", str1, len(str1))
}

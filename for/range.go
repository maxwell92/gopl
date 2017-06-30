package main

import (
	"fmt"
)

func main() {
	strs := []string{
		"aaa",
		"bbb",
		"ccc",
		"ddd",
	}

	for _, s := range strs {
		fmt.Println(s)
	}
}

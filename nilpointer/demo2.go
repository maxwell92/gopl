package main

import (
	"fmt"
)

func main() {
	var dic map[string]string
	// dic := make(map[string]string)
	fmt.Printf("%p\n", dic)
	fmt.Printf("%s\n", dic["dev"])

	// segment fault
	dic["dev"] = "yce"
	fmt.Printf("%s\n", dic["dev"])
}

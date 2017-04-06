package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
	fmt.Printf("%s\n", strings.Trim(" abc\n", "\n"))
}

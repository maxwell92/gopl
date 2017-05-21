package main

import (
	"fmt"
)

const (
	read   byte = 1 << iota // 0001
	write                   // 0010
	exec                    // 0100
	freeze                  // 1000
)

func main() {
	fmt.Printf("%05b, %05b %05b, %05b\n", read, write, exec, freeze)
	a := read | write | freeze // 0001 | 0100 | 1000 = 1011
	b := read | freeze | exec  // 0001 | 1000 | 0100 = 1101
	c := a &^ b                // 1011 &^ 1101 = 0010

	fmt.Printf("%04b &^ %04b = %04b\n", a, b, c)
}

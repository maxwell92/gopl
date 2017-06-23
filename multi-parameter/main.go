package main

import (
	"fmt"
)

func print(str ...string) {
	fmt.Println(str) // works well
	// fmt.Printf(str...)
}

func main() {
	hello := "hello"
	space := " "
	world := "world"
	print(hello, space, world)
}

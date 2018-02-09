package main

import (
	"fmt"
)

func main() {
	switch x := 1; x {
	case 1:
		fmt.Println("this is case1")
		fallthrough
	case 2:
		fmt.Println("this is case2")
		fallthrough
	default:
		fmt.Println("this is default")
	}
}

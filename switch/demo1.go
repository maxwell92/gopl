package main

import (
	"fmt"
)

func main() {
	switch x := 1; x {
	case 1:
		x += 1
		fmt.Println(x) // 2
		fallthrough
	case 3:
		fmt.Println(x) // 2
	}
}

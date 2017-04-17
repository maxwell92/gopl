package main

import (
	"fmt"
)

func main() {
	for i, j := 0, 10; i < 10; i++, j-- {
		fmt.Println(i)
		fmt.Println(j)
	}
}

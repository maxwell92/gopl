package main

import (
	"fmt"
)

func main() {
	list := make([]int, 2)
	list[0] = 0
	list[1] = 1
	fmt.Println(list)
	list = append(list, 2)
	fmt.Println(list)
}

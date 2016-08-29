package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	for _, v := range a {
		fmt.Println(v)
		a = append(a, v)
	}
	fmt.Println(a)
}

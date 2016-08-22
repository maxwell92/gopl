package main

import "fmt"

func main() {
	a := []int{1, 2, 1, 3}
	for _, v := range a {
		switch v {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		default:
			fmt.Println("default")
		}
	}
}

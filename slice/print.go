package main

import "fmt"

func main() {
	b := []int{1, 2, 3}
	fmt.Println(b)

	l := b[1:3]
	fmt.Println(l)

	l = b[1:]
	fmt.Println(l)

	k := b[:3]
	fmt.Println(k)
}

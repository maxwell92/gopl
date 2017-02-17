package main

import "fmt"

func main() {
	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b)

	page := 1
	start := page - 1
	offset := 5

	a := b[start : start+offset]
	fmt.Println(a)

	l := b[0:3]
	fmt.Println(l)

	l = b[1:]
	fmt.Println(l)

	k := b[:3]
	fmt.Println(k)
}

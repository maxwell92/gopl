package main

import "fmt"

const (
	ID0 = 1 << iota //1
	ID1             //2
	ID2             //4
	ID4             //8

)

func main() {
	fmt.Println(ID0)
	fmt.Println(ID1)
	fmt.Println(ID2)
	fmt.Println(ID4)
}

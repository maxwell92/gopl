package main

import "fmt"

const (
	ID0 = iota //0
	ID1        //1
	ID2        //2

	// iota value is 3 now
	ID3 = iota + 3 //6
	ID4            //7
	ID5            //8
)

func main() {
	fmt.Println(ID0)
	fmt.Println(ID1)
	fmt.Println(ID2)

	fmt.Println(ID3)
	fmt.Println(ID4)
	fmt.Println(ID5)
}

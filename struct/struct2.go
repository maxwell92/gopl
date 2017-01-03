package main

import (
	"fmt"
)

type Car struct {
	Name string
}

func main() {
	mycar := &Car{}

	if mycar == nil {
		fmt.Println("it's nil/null")
	}
	fmt.Printf("%v\n", mycar)
	fmt.Printf("%p\n", mycar)

	garage := make([]*Car, 5) // allocate 5 elements of type *Car
	// garage := make([]*Car, 0)

	garage = append(garage, mycar)

	fmt.Printf("len: %d\n", len(garage)) // 6
}

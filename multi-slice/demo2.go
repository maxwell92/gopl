package main

import (
	"fmt"
)

func main() {
	table := [][]int32{
		{
			1, 2, 3, 4, 5,
		},
		{
			6, 7, 8, 9, 10,
		},
	}

	fmt.Printf("%v\n", table)
	fmt.Printf("%v\n", table[0])
	fmt.Printf("%v\n", table[][0])
}

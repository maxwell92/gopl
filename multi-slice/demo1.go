package main

import (
	"fmt"
)

func main() {
	single := []int{
		1,
		2,
		3,
		4,
	}

	fmt.Printf("%v\n", single)

	multi := [][]int{
		[]int{
			1,
			2,
			3,
			4,
		},
		[]int{
			2,
			3,
			4,
			5,
		},
	}

	fmt.Printf("%v\n", multi)

	multi1 := make([][]int, 4)
	multi1[0] = make([]int, 4) 
	multi1[1] = make([]int, 4) 
	multi1[2] = make([]int, 4) 
	multi1[3] = make([]int, 4) 

	multi1[0][0] = 0
	multi1[0][1] = 1
	multi1[0][2] = 2
	multi1[0][3] = 3
	multi1[1][0] = 4
	multi1[1][1] = 5
	multi1[1][2] = 6
	multi1[1][3] = 7

	fmt.Printf("%v\n", multi1)

}


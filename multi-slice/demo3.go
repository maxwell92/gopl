package main

import (
	"fmt"
)

func main() {
	table := make([][]int, 0)	
	fmt.Printf("%v\n", table)

	row0 := make([]int, 5)
	table = append(table, row0)
	fmt.Printf("%v\n", table)
	
	colx := make([]int, 1)
	for i:=0; i < len(table); i++ {
		table[i] = append(table[i], colx...)
	}
	fmt.Printf("%v\n", table)
}

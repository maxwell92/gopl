package main

import (
	"fmt"
)

func main() {
	var k1, k2, k3, k4 string

	fmt.Println("k1:")
	fmt.Scanf("%s", &k1)
	fmt.Println("k2:")
	fmt.Scanf("%s", &k2)
	fmt.Println("k3:")
	fmt.Scanf("%s", &k3)
	fmt.Println("k4:")
	fmt.Scanf("%s", &k4)

	fmt.Printf("%s#\n", k1)
	fmt.Printf("%s#\n", k2)
	fmt.Printf("%s#\n", k3)
	fmt.Printf("%s#\n", k4)
}

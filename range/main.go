package main

import (
	"fmt"
)

type Car struct {
	Name string
}

func main() {

	garage := make([]*Car, 3)

	benz := &Car{Name: "Benz"}
	bmw := &Car{Name: "BMW"}
	lexus := &Car{Name: "Lexus"}

	garage[0] = benz
	garage[1] = bmw
	garage[2] = lexus

	for i, c := range garage {
		fmt.Printf("%d, %s\n", i, c.Name)
	}
}

package main

import (
	"fmt"
)

type Brand struct {
	Name   string
	CarPtr *Car
}

type Car struct {
	Name  string
	Brand *Brand
}

func main() {
	c := &Car{
		Name: "ML350",
		Brand: &Brand{
			Name: "Benz",
		},
	}

	c.Brand.CarPtr = c

	fmt.Println(c.Name)
	fmt.Println(c.Brand.Name)
	fmt.Println(c.Brand.CarPtr.Name)
}

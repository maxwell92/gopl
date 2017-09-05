package main

import (
	"fmt"
)

type Car struct{
	Name string
}

func main() {
	var benz *Car
	fmt.Printf("%p\n", benz)
	fmt.Printf("%s\n", benz.Name)
}

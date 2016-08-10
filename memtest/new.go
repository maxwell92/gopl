package main

import "fmt"

type Car struct {
	Name string
	Type []string
}

func main() {
	benz := new(Car)
	benz.Name = "Benz"
	benz.Type = make([]string, 2)
	benz.Type[0] = "ML350" // index out of range
	benz.Type[1] = "G55"

	fmt.Println(benz)

	var landrover Car
	landrover.Name = "Land Rover"
	landrover.Type = make([]string, 2)
	landrover.Type[0] = "Discovery" // index out of range
	landrover.Type[0] = "RangeRover"

	fmt.Println(landrover)
}

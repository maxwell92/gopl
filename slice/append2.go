package main

import "fmt"

func main() {
	hello := make([]string, 0)
	hello = append(hello, "hello ")
	fmt.Println(hello)
	world := make([]string, 0)
	world = append(world, " world")
	fmt.Println(world)

	//13: cannot use world[:len(world)] (type []string) as type string in append
	//hello = append(hello[:len(hello)], world[:len(world)])

	//14: cannot use copy(hello[:len(hello)], world[:len(world)]) (type int) as type []string in assignment
	//hello = copy(hello[:len(hello)], world[:len(world)])
	hello = append(hello, world...)
	fmt.Println(hello)
}

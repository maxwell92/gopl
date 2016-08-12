package main

import "fmt"

type Car struct {
	Name string
	Type []string
}

// receiver parameter: pointer
func (mycar *Car) Print() {
	fmt.Println("This is Print")
}

// receiver parameter: named type
func (myCar Car) Display() {
	fmt.Println("This is Display")
}

// cross call from named type
func (myCar Car) CrossCall1() {
	fmt.Println("cross call from named type:")
	myCar.Print()
	myCar.Display()
}

// cross call from pointer
func (myCar *Car) CrossCall2() {
	fmt.Println("cross call from pointer:")
	myCar.Print()
	myCar.Display()
}

func main() {
	// receiver argument: pointer
	Benz := new(Car)
	Benz = &Car{
		Name: "Benz",
		Type: []string{"ML350", "G55"},
	}

	// receiver argument: named type
	var BMW Car
	BMW = Car{
		Name: "BMW",
		Type: []string{"M3", "M6"},
	}

	// recevier argument | recevier parameter
	//   pointer         |     pointer
	Benz.Print()
	//   pointer         |     named type
	Benz.Display()
	//   named type      |     pointer
	BMW.Print()
	//   named type      |     named type
	BMW.Display()
	// all pass

	Benz.CrossCall2()
	BMW.CrossCall1()
	// all pass
}

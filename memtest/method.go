package main

import "fmt"

type Car struct {
	Name string
	Type []string
}

// Pointer mycar is invisiable inside the Print()
// It is called the Pointer Receiver.
// It is used for copying all the variable if it is too large.
func (mycar *Car) Print() {
	//fmt.Println(myCar.Name) //myCar undefined.
	//fmt.Println(Name) // Name undefined.
	//name := myCar.Name
	//fmt.Println(name)

	mycar = &Car{
		Name: "Land Rover",
		Type: []string{"Range Rover", "Discovery"},
	}
	fmt.Println(*mycar)
}

// Variable mycar is visiable inside the Print()
// It is called the Named Type.
func (myCar Car) Display() {
	fmt.Println(myCar)
}

func main() {
	Benz := new(Car)
	Benz = &Car{
		Name: "Benz",
		Type: []string{"ML350", "G55"},
	}
	// No matter pointer or normal variable it is, method could be called by it.
	Benz.Print()
	Benz.Display()

	var BMW Car
	BMW = Car{
		Name: "BMW",
		Type: []string{"M3", "M6"},
	}

	BMW.Print()
	BMW.Display()

}

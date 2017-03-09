package main

import "fmt"

type ICar interface{
	Beep(msg string)
	Run(start, dest string) bool
}

type Person struct {
	Name string	
	Car ICar
}

type Benz struct {
	Name string
}

func (b *Benz) Beep(msg string) {
	fmt.Println(msg)
}

func (b *Benz) Run(s, d string) bool {
	return true
}

func (b *Benz) Speed() {}


func main() {

	p := &Person{
		Name: "mushroom",
		Car: &Benz{
			Name: "benz",	
		},
	}

	p.Car.Beep("this")
	p.Car.Run("Beijing", "San Francisco")

/*
	var a interface{}
	var b = 1 

	a = b
	fmt.Println(a)
*/
}

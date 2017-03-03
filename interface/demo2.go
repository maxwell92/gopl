package main

import "fmt"

type Animal interface {
	Hi()
}

type Cat struct{}

func (c Cat) Hi() { fmt.Println("Cat!") }

type Dog struct{}

func (d Dog) Hi() { fmt.Println("Dog!") }

func main() {
	Zoo := make([]Animal, 0)
	c := new(Cat)
	d := new(Dog)
	Zoo = append(Zoo, c)
	Zoo = append(Zoo, d)

	for _, a := range Zoo {
		a.Hi()
	}
}

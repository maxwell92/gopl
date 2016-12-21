package main

import "fmt"

type Thing interface {
	Print()
}

type Id int
type Name string

func (i *Id) Print() {
	fmt.Println("This is Number")
}

func (n *Name) Print() {
	fmt.Println("This is String")
}

func main() {
	str := new(Name)
	str.Print()

	id := new(Id)
	id.Print()
}

package main

import "fmt"

type IPrint interface {
	PrintSelf()
}

func Print(ip IPrint) {
	ip.PrintSelf()
}

type Mush int

func (m Mush) PrintSelf() {
	fmt.Println(m)
}

type Room string

func (r Room) PrintSelf() {
	fmt.Println(r)
}

func main() {
	mu := new(Mush)
	ro := new(Room)

	Print(mu)
	Print(ro)
}

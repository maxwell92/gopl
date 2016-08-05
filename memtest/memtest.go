package main

import (
	"fmt"
)

type user struct {
	Name   string
	Age    float32
	Gender int
}

func main() {
	me := new(user)
	me.Name = "amx"
	me.Age = 23.4
	me.Gender = 1

	fmt.Printf("%p\n", &me)
	fmt.Printf("%#x %p %p\n", &me.Name, &me.Age, &me.Gender)

	/*
		var i int
		var p *int
		i = 1
		p = &i
		fmt.Printf("%d %p %d\n", i, p, *p)
	*/
}

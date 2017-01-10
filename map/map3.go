package main

import (
	"fmt"
)

type Page struct {
	P       int
	Content string
}

func main() {
	dic := make(map[string]*Page, 5)
	dic["1"] = &Page{
		P:       1,
		Content: "Hello World",
	}

	if p, c := dic["1"]; c {
		// p is the value, and c is bool
		fmt.Printf("p: %v, c: %v\n", p, c)
	}
}

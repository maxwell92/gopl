package main

import "fmt"

func Say() {
	fmt.Println("Hello World")
}

func main() {
	a := 1
	if a == 1 {
		Say()
	}
}

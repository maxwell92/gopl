package main

import "fmt"

func doList(args ...interface{}) {
	fmt.Printf("doList: %d, %v\n", len(args), args)
}

func List(args ...interface{}) {
	fmt.Printf("List: %d, %v\n", len(args), args)
	doList(args)
}

func main() {
	List(1, 2, 3)
}

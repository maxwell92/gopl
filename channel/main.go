package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string, 5)
	ch2 := make(chan string, 5)
	ch3 := make(chan string, 5) // if ch3 := make(chan string) it will cause deadlock
	str := "Hello world"
	ch1 <- str
	ch2 <- str
	ch3 <- str

	fmt.Printf("ch1: %v, %s\n", ch1, <-ch1)
	fmt.Printf("ch2: %v, %s\n", ch2, <-ch2)
	x, ok := <-ch3 //ok will be bool
	fmt.Printf("ch3: %v, %v\n", x, ok)
}

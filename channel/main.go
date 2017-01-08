package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string, 5)
	ch2 := make(chan string, 5)
	str := "Hello world"
	ch1 <- str
	ch2 <- str

	fmt.Printf("ch1: %v, %s\n", ch1, <-ch1)
	fmt.Printf("ch2: %v, %s\n", ch2, <-ch2)
}

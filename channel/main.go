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
	// below will cause "too many arguement in range error"
	/*
		for k, v := range ch3 {
			fmt.Println(v)
		}
	*/

	ch4 := make(chan string, 5)
	ch4 <- "1"
	ch4 <- "2"
	ch4 <- "3"
	ch4 <- "4"

	fmt.Printf("total: %d\n", len(ch4))
	// range will get element from the channel!
	for s := range ch4 {
		fmt.Printf("%s: %d\n", s, len(ch4))
	}

	// it's blocking channel
	ch5 := make(chan string)
	// it's non-blocking channel
	ch6 := make(chan string, 1)
}

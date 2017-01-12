package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string, 5)
	ch1 <- "1"
	ch1 <- "2"
	ch1 <- "3"
	ch1 <- "4"
	ch1 <- "5"
	// Channel didn't support indexing
	/*
		for i := 0; i < 5; i++ {
			fmt.Printf("%s\n", ch1[i])
		}
	*/

}

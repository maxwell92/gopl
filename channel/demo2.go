package main

import (
	"fmt"
	"time"
)

var ch1 chan string
var ch2 chan string

func sender(info string, interval int, ch chan string) {
	for {
		ch <- info
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func main() {
	ch1 = make(chan string)
	ch2 = make(chan string, 10)

	go sender("A", 1, ch1)
	go sender("B", 10, ch2)
	for {
		select {
		case s1 := <-ch1:
			fmt.Println(s1)
		case s2 := <-ch2:
			fmt.Println(s2)
		}
	}
}

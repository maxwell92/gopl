package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a: ", i)
			if i == 1 {
				runtime.Gosched()
			}
		}
	}()

	<-exit
}

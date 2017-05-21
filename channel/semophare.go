package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{}, 2) //control how many routines could run at a time

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
	}

	wg.Wait()
}

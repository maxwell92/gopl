package main

import (
	"fmt"
	"sync"
)

var common = 1
var alock sync.Mutex
var block sync.Mutex

func read(wg *sync.WaitGroup) {
	alock.Lock()
	block.Lock()
	common = 3
	block.Unlock()
	alock.Unlock()
	fmt.Println(common)
	wg.Done()
}

func write(wg *sync.WaitGroup) {
	block.Lock()
	alock.Lock()
	common = 2
	alock.Unlock()
	block.Unlock()
	fmt.Println(common)
	wg.Done()
}

func main() {
	cnt := 2
	wg := new(sync.WaitGroup)
	wg.Add(cnt)

	for i := 0; i < cnt; i++ {
		if i/2 == 0 {
			go write(wg)
		} else {
			go read(wg)
		}
	}
	wg.Wait()
}

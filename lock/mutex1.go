package main

import (
	"fmt"
	"sync"
)

var common = 1
var alock sync.Mutex
var block sync.Mutex

func read() {
	block.Lock()
	common = 3
	block.Unlock()
	fmt.Println(common)
}

func write() {
	alock.Lock()
	read()
	common = 2
	alock.Unlock()
	fmt.Println(common)

}

func access(wg *sync.WaitGroup) {
	write()
	wg.Done()
}

func main() {
	cnt := 2
	wg := new(sync.WaitGroup)
	wg.Add(cnt)

	for i := 0; i < cnt; i++ {
		go access(wg)
	}
	wg.Wait()
}

package main

import (
//	"sync"
	"fmt"
	"time"
	"sync"
)

var common = 1
var rlock sync.RWMutex
/*
func read(wg *sync.WaitGroup) {
	fmt.Printf("Read: %d\n", common)
	wg.Done()
}

func write(wg *sync.WaitGroup) {
	common = 2
	fmt.Printf("Write: %d\n", common)
	wg.Done()
}

func main() {
	cnt := 5
	wg := new(sync.WaitGroup)
	wg.Add(cnt)

	for i := 1; i < cnt; i++ {
		go read(wg)
	}

	go write(wg)
	wg.Wait()
}
*/


func read() {
	rlock.Lock()
	for i := 0; i < 5; i++ {
		fmt.Printf("Read: %d\n", common)

	}
	rlock.Unlock()
}

func write() {
	common = 2
	fmt.Printf("Write: %d\n", common)
}

func main() {

	go read()

	for i := 0; i < 2; i++ {
		go write()
	}
	time.Sleep(time.Duration(2) * time.Second)
}

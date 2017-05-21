package main

import (
	"fmt"
	"time"
)

func test() {
	defer println("defer")
	println("hi")
	time.Sleep(time.Second * 5)
}

func main() {
	test()
	fmt.Println("main")
}

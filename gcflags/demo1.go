package main

import ()

func test() *int {
	a := 0x100
	return &a
}

func main() {
	var a *int = test()
	println(a, *a)
}

// go build -gcflags "-m" demo1.go
// go tool objdump -s "main½.main" demo1

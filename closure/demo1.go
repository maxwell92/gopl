package main

import ()

func test(x int) func() {
	return func() {
		println(x)
	}
}

func main() {
	f := test(123)
	f()
}

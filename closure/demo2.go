package main

func test(x int) func() {
	println(&x)

	return func() {
		println(&x, x)
	}
}

func main() {
	f := test(0x100)
	f()
}

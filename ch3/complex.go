package main

import (
	"fmt"
)

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	a := 1 + 2i
	b := 3i + 4
	fmt.Println(a)
	fmt.Println(b)
}

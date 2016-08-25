package main

import (
	"fmt"
)

func main() {
	cnt := 0
	for x := 2; x < 100; x++ {
		for i := 2; i < x; i++ {
			if x%i == 0 {
				if i <= x/i {
					fmt.Printf("%d prime number\n", x)
					cnt++
					break
				}
			}
		}
	}

	fmt.Printf("%d total\n", cnt)
}

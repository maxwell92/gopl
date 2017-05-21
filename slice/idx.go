package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}

	for idx, a := range array {
		fmt.Printf("%d: %d\n", idx, a)
		if a == 3 {
			fmt.Println("found")
			break
		}

		if idx == (len(array) - 1) {
			fmt.Println("not found")
			return
		}

	}
}

package main

import (
	"fmt"
)

func Average(array []int) int {
	var sum int
	for _, data := range array {
		sum += data
	}

	return sum / len(array)
}

func main() {
	array := []int{1, 2, 3, 4}

	average := Average(array)
	fmt.Println(average)
}

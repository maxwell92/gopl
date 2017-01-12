package main

import "fmt"

func main() {
	str := make([]string, 5)

	str[0] = "0"
	str[1] = "1"
	str[2] = "2"
	str[3] = "3"
	str[4] = "4"

	fmt.Printf("[1:]: %v\n", str[1:])
}

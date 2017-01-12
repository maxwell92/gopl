package main

import "fmt"

func main() {
	str := make([]string, 2)
	str[0] = "0"
	str[1] = "1"
	fmt.Println(str)

	// it's index out of range error
	// str = str[1:2]
	// str[2] = "2"
	// fmt.Println(str)
}

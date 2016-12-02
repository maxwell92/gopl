package main

import "fmt"

func main() {
	str := make([]string, 0)
	fmt.Println(len(str))

	str = append(str, "")
	fmt.Println(len(str))
}

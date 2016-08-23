package main

import "fmt"

func main() {
	str := "hello world"
	str = fmt.Sprintf("%s", str)
	str = fmt.Sprintf("%s", str)

	fmt.Println(str)
	fmt.Println(str)
	fmt.Println(str)
}

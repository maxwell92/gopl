package main

import "fmt"

func Print(content string, print func(...interface{}) (int, error)) {
	print(content)
}

func main() {
	Print("hello, world", fmt.Println)
}

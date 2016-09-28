package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("hi.txt")
	if err != nil {
		fmt.Println("read file error")
	} else {
		fmt.Println(string(bytes))
	}
}

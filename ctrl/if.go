package main

import (
	"fmt"
)

func Test() (bool, string) {
	if 2 > 1 {
		return true, ""
	} else {
		return false, "wrong"
	}
}

func main() {
	if ok, err := Test(); ok {
		fmt.Println(ok)
	} else {
		fmt.Println(err)
	}
}

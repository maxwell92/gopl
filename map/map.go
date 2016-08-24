package main

import (
	"fmt"
)

const (
	ab int = 1
	bc int = 2
	cd int = 3
)

var errmap = map[int]string{
	ab: "ab",
	bc: "bc",
	cd: "cd",
}

func main() {
	fmt.Println(errmap[bc])
}

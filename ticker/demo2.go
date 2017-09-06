package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Duration(2) * time.Second)
	for i := range t.C {
		fmt.Println("Hi ", i)
	} 
}

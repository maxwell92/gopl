package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Duration(2) * time.Second)
	for i := range t.C {
		fmt.Println("Hi ", i)
	} 
}

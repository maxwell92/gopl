package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Duration(2) * time.Second)
	for {
		select {
		case <-t.C:
			{
				fmt.Printf("timeout @%s\n", time.Now())
				t.Reset(time.Duration(1) * time.Second)
			}
		}
	}
}

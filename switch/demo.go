package main

import (
	"fmt"
)

func main() {
	str := "start"
start:
	for {
		switch str {
		case "run":
			fmt.Println("run")
			return
		case "start":
			fmt.Println("start")
			break start
		}
	}
}

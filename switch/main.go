package main

import (
	"fmt"
)

func main() {
	str := "1"

	switch str {
	case "1":
		str = "one"
		//		break
	case "2":
		str = "not one"
		//		break
	default:
		str = "unknown"
	}
	fmt.Println(str)
}

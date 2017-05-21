package main

import "fmt"

func main() {
	dic := make(map[string]string)
	dic["sheep"] = "magic"
	fmt.Println(dic["sheep"])
	fmt.Println(dic["magic"])
}

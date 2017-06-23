package main

import (
	"fmt"
)

func Sprintf(a, b, c string) string {
	str := fmt.Sprintf("%s%s%s", a, b, c)
	return str
}

func Add(a, b, c string) string {
	str := a + b + c
	return str
}

func main() {
	hello := "hello"
	space := " "
	world := "world"

	fmt.Println(Sprintf(hello, space, world))
	fmt.Println(Add(hello, space, world))

}

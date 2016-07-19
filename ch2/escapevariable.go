package main
import "fmt"

var global *int

func main() {
	fmt.Println(global) // nil
	//fmt.Println(*global)
	f()
	fmt.Println(global)
	//fmt.Println(*global)
}

func f() {
	var i int
	i = 1
	global = &i
}

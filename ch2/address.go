package main
import "fmt"

func main() {
	fmt.Println(f())
	fmt.Println(f())
}

func f() *int {
	v := 1
	return &v
}

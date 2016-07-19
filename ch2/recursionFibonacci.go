package main
import (
	"os"
	"fmt"
	"strconv"
)

var i int

func main() {
	n, _ err := strconv.Atoi(os.Args[1])
	i = 0
	fmt.Println(fibo(n))
}

func fibo(n int) int {	
	if i < n {
		return fibo(a) + fibo(b)	
	}	
}

package main
import (
	"os"
	"fmt"
	"strconv"
)

var i int

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	i = 0
	fmt.Println(fibo(n))
}

func fibo(n int) int {	
	if n == 1 || n == 2{
		return 1	
	} else {
		return fibo(n - 1) + fibo(n - 2)	
	}
}

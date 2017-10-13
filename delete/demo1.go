package main
import (
	"fmt"
)

func main() {
	// list := []int{1, 2, 3} 
	mp := map[int]int{
		0: 0,
		1: 1,
		2: 2,
	}
	fmt.Println(mp)
	delete(mp, 2)
	fmt.Println(mp)
}

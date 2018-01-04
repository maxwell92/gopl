package main
import(
	"fmt"
)

func main() {
	arr1 := make([]int, 0)
	arr2 := []int{}
	var arr3 []int
	
	println(&arr1)
	println(&arr2)
	println(&arr3)

	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n", arr2)
	fmt.Printf("%v\n", arr3)
}

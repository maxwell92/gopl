package main
import (
	"fmt"
	"unsafe"
)

func main() {
	var number = 30 
	var str = "hello"
	
	ptrNumber := &number
	ptrStr := &str

	fmt.Printf("%d, %d\n", number, *ptrNumber)
	fmt.Printf("%s, %s\n", str, *ptrStr)
/*
 * Wrong use of unsafe.Pointer()
	ptrAny := unsafe.Pointer(&ptrNumber)
	fmt.Printf("%d\n", ptrAny)
*/
}

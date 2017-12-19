package main
import(
	"fmt"
	"reflect"
	"unsafe"
)

func pp(ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf("%x\n", *h)

}

func main() {
	s := "abcdefg"	
	s1 := s[:3]
	
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))

	ptr := &s
	pp(ptr)
}

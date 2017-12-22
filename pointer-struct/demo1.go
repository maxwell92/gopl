package main
import(
	"fmt"
	"unsafe"
)
type data struct {
	x int
	Y int
}

func NewData() *data {
	return new(data)
}

func main() {
	d := NewData()	
	d.Y = 200
	p :=  (* struct{x int})(unsafe.Pointer(d))
	p.x = 100
	
	fmt.Printf("%v\n", *d)
}

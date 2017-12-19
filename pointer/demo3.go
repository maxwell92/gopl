package main
import(
	"unsafe"
)

func main() {
	a := 0
	p := &a
	println(unsafe.Sizeof(p))
}

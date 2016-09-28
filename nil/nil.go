package main

import (
	"fmt"
	"unsafe"
)

type test struct {
	msg string
}

func nilget() *test {
	return nil
}

func checkNotNil(value interface{}) bool {
	v := unsafe.Pointer(&value)
	if v == nil {
		fmt.Printf("nil: %v\n", v)
		return false
	} else {
		fmt.Printf("not nil: %v\n", v)
		return true
	}

}

func main() {
	value := nilget()
	fmt.Printf("%v\n", checkNotNil(value))
}

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type test struct {
	msg string
}

func nilget() *test {
	return nil
}

// better not pass pointer nil to interface{}
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

func checkValidate(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Ptr && value != nil {
		return true
	} else {
		return false
	}
}

func main() {
	value := nilget()
	//fmt.Printf("%v\n", checkNotNil(value))
	fmt.Printf("%v\n", checkValidate(value))
}

// go version 1.7
package main

import (
	"fmt"
)

type Object interface {
	Get()
}

type Unstructured struct {
	Name string
}

// 如果实现了Object接口，一切正常
func (u *Unstructured) Get() {}

// 如果只使用了下面的方法，而没有实现Object接口，会有如下报错：
// ./main.go:19: cannot use Unstructured literal (type *Unstructured) as type Object in assignment:
// *Unstructured does not implement Object (missing Get method)

// func (u *Unstructured) Got() {}

var _ Object = &Unstructured{}

func main() {
	fmt.Println("Hello")
}

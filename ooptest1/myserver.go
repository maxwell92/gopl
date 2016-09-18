package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type showController struct {
	baseController
	msgController
}

type baseController struct {
	*iris.Context
}

type msgController struct {
	msg string
}

func (s showController) Get() {
	//fmt.Println(s.msg) //TODO: why s.msg is "" ???
	s.msg = "Hello World"
	s.Write(s.msg)
}

func main() {
	myshow := new(showController)
	myshow.msg = "Hello World"

	fmt.Println(myshow.msg)
	iris.API("/msg", *myshow) //TODO: how to reflect ??
	iris.Listen(":9999")
}

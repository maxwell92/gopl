package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type Someone struct {
	Name string
}

type OperationAPI struct {
	*iris.Context
	path   string
	person *Someone
}

func (op *OperationAPI) sync() {
	fmt.Printf("sync, op: %v, %x\n", op, &op)
	fmt.Printf("sync, filepath: %s, %x\n", op.path, &op.path)
	fmt.Printf("sync, Name: %s, %x\n", op.person.Name, &op.person.Name)
}

// the op here is not the same with the one in iris / Init() / main
func (op OperationAPI) Get() {
	op.person = new(Someone)
	op.person.Name = "Moss"

	fmt.Printf("get, op: %v, %x\n", op, &op)
	fmt.Printf("get, filepath: %s, %x\n", op.path, &op.path)
	fmt.Printf("get, Name: %s, %x\n", op.person.Name, &op.person.Name)
	op.sync()
}

func Init() *OperationAPI {
	op := new(OperationAPI)
	op.path = "path"
	op.person = new(Someone)
	op.person.Name = "Roy"

	fmt.Printf("init, op: %v, %x\n", op, &op)
	fmt.Printf("init, filepath: %s, %x\n", op.path, &op.path)
	fmt.Printf("init, Name: %s, %x\n", op.person.Name, &op.person.Name)

	return op
}

func main() {
	op := Init()

	fmt.Printf("main, op: %v, %x\n", op, &op)
	iris.API("/sync", *op)
	iris.Listen(":8081")
}

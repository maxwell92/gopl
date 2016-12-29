package main

import (
	"fmt"
	"github.com/kataras/iris"
	"os/exec"
)

type OperationAPI struct {
	*iris.Context
}

func (o OperationAPI) Get() {
	fmt.Println("Sync Start")
	/*
		cmd := exec.Command("/bin/bash", "-C", "./helloworld.sh")
		result, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(result))
	*/

	cmd := exec.Command("/bin/bash", "-C", "./helloworld.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		o.Write("Sync Error\n", err)
	}
	fmt.Println("Sync over")
	/*
		o.Write(string(result))
	*/
	o.Write("Sync Over\n")
}

func main() {
	op := new(OperationAPI)
	iris.API("/sync", *op)
	iris.Listen(":8081")
}

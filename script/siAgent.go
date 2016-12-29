package main

import (
	"fmt"
	"github.com/kataras/iris"
	"os"
	"os/exec"
)

type OperationAPI struct {
	*iris.Context
	scriptPath string
}

var usage = `Usage: siAgent <filepath>
`

// op here is hot the same with the oen in Init()
func (op OperationAPI) Get() {
	fmt.Println("Sync Start")
	/*
		cmd := exec.Command("/bin/bash", "-C", "./helloworld.sh")
		result, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(result))
	*/

	fmt.Printf("filepath: %s", op.scriptPath)
	//cmd := exec.Command("/bin/bash", "-C", op.scriptPath)
	cmd := exec.Command("/bin/bash", "-C", "./helloworld.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		op.Write("Sync Error\n", err)
	}
	fmt.Println("Sync over")
	/*
		o.Write(string(result))
	*/
	op.Write("Sync Over\n")
}

func UsageAndExit() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

}

func Init() {

	op := new(OperationAPI)
	op.scriptPath = os.Args[1]

	f, err := os.Open(op.scriptPath)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	iris.API("/sync", *op)
	iris.Listen(":8081")
}

func main() {
	UsageAndExit()
	Init()
}

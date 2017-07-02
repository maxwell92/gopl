package main

import (
	"fmt"
	"os/exec"
)

func main() {
	/*
		command1 := "/bin/echo"
		command2 := "Hello World"
		cmd := exec.Command(command1, command2)
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}

		fmt.Println(string(output))
	*/

	command1 := "echo"
	args1 := []string{
		"Hello World",
	}
	cmd1 := exec.Command(command1, args1...)
	cmdOutput, err := cmd1.StdoutPipe()
	if err != nil {
		panic(err)
	}

	cmd1.Start()

	command2 := "grep"
	args2 := []string{
		"Hello a",
	}
	cmd2 := exec.Command(command2, args2...)
	cmd2.Stdin = cmdOutput

	output, err := cmd2.Output()
	if err != nil {
		// exit code 1
		panic(err)
	}

	fmt.Println(string(output))

}

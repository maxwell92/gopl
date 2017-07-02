package main

import (
	"fmt"
	"os/exec"
)

func main() {
	command1 := "/bin/echo"
	command2 := "Hello World"
	cmd := exec.Command(command1, command2)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}

package main

import (
	"fmt"
	"os"
	"tools/agent"
)

func UsageAndExit() {
	if len(os.Args) != 1 {
		fmt.Fprint(os.Stderr, agent.USAGE)
		os.Exit(1)
	}

}

func main() {
	UsageAndExit()
	agent.Run()
}

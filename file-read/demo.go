package main

import (
	"bufio"
	"fmt"
	// "io"
	"os"
)

//define function to read file line by line
func ReadfilebyLine(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fp := bufio.NewReader(f)
/*
	for {
		oneline, err1 := fp.ReadString('\n')
		if err1 != nil && io.EOF == err1 {
			break
		}
		fmt.Print(oneline)
	}
*/
	for {
		line, err := fp.ReadBytes('\n')
		if err != nil {
			fmt.Printf("%s\n", err)	
			break
		}
		// fmt.Println(string(line))
		print(string(line))
	}
}

func main() {
	ReadfilebyLine("./makefile.sh")
}

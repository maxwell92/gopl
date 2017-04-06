package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name .... \n")

	input, _ := consoleReader.ReadString('\n')

	fmt.Println("Your name is : ", input)

}

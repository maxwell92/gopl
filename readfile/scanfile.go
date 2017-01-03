package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./helloworld")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		str := scanner.Text()
		if strings.Split(str, "#")[0] == "" || len(strings.Split(str, " ")) == 1 {
			fmt.Printf("\"%s\" is a comment or uncomplete pair. \n", str)
		} else {
			// content := str[:len(str)]
			content := strings.Split(str, " ")
			fmt.Printf("Pair: %s\n", content)

			// key := strings.Split(content, " ")[0]
			// value := strings.Split(content, " ")[1]

			fmt.Printf("key: %s\n", content[0])
			fmt.Printf("Value: %s\n", content[1])

		}

	}
}

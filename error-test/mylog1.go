package main

import (
	"log"
	"os"
)

func main() {
	mylog := log.New(os.Stdout, "This is mylog", log.Lshortfile)
	mylog.Println("ssss")
}

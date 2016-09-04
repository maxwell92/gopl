package main

import "log"

func main() {
	//log.SetFlags(log.Lshortfile)
	log.SetPrefix("[Custom Prefix]")
	log.Println("Hello World from log pkg")
}

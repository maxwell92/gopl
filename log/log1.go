package main

import "log"

func main() {
	message := "Fatal"
	//log.Fatalf("It's %s\n", message)
	log.Panicf("It's %s\n", message)
}

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Ldate: %d\n", log.Ldate)
	fmt.Printf("Ltime: %d\n", log.Ltime)
	fmt.Printf("Lmicroseconds: %d\n", log.Lmicroseconds)
	fmt.Printf("Llongfile: %d\n", log.Llongfile)
	fmt.Printf("Lshortfile: %d\n", log.Lshortfile)
	fmt.Printf("LUTC: %d\n", log.LUTC)
	fmt.Printf("LstdFlags: %d\n", log.LstdFlags)
}

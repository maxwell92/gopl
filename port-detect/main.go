package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	DEV1     = "172.21.1.11:8080"
	DEV2     = "172.17.106.52:8080"
	NOTEXIST = "172.21.1.12:8080"
	WRONG    = "172.21.1$80"
	ERROR    = "172.21:8080"
)

func test(host string) bool {
	log.Printf("Dialing %s\n", host)
	dialer := &net.Dialer{
		Timeout: time.Duration(10) * time.Second,
	}
	/*
		conn, err := net.Dial("tcp", host)
	*/

	conn, err := dialer.Dial("tcp", host)
	if err != nil {
		log.Printf("Dial %s Error: err=%s", host, err)
		return false
	}
	defer conn.Close()

	log.Printf("Connected to %s\n", host)

	return true
}

func main() {
	hosts := []string{DEV1, DEV2, NOTEXIST, WRONG, ERROR}

	for _, h := range hosts {
		fmt.Println(test(h))
	}
}

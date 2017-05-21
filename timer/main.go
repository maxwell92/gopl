package main

import (
	"log"
	"time"
)

var eventCh chan int

func send() {
	for {
		time.Sleep(time.Duration(1) * time.Second)
		eventCh <- 1
	}
}

func main() {
	eventCh = make(chan int)
	t := time.NewTimer(time.Duration(10) * time.Second)
	timeoutCh := t.C
	for {
		select {
		case s := <-timeoutCh:
			{
				log.Println("over", s)
				return
			}
		case e := <-eventCh:
			{
				log.Println("run", e)
			}
		}
	}
}

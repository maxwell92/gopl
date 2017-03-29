package main

import (
	"log"
	"time"
)

func sender(e int, interval int, ch chan int) {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		ch <- e
	}
}

func main() {
	eventChan := make(chan int)
	stopCh := make(chan int)

	go sender(1, 1, eventChan)
	go sender(1, 10, stopCh)

	for {
		select {
		case e := <-eventChan:
			{
				log.Println(e)
			}
		case t := <-stopCh:
			{
				log.Println(t)
				return
			}
		}
	}
}

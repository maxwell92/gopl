package main

import (
	"fmt"
	"log"
	"time"
)

func emit(eventChan chan string) {
	count := 0
	for {
		eventChan <- ""
		count++
		log.Println(count)
		if count == 3 {
			time.Sleep(time.Duration(5) * time.Second)
			count = 0
			continue
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func main() {
	timeout := time.NewTimer(time.Duration(3) * time.Second)
	eventChan := make(chan string)
	stopChan := make(chan bool)

	go emit(eventChan)
	for {
		select {
		case <-eventChan:
			{
				timeout.Reset(time.Duration(3) * time.Second)
			}
		case <-timeout.C:
			{
				fmt.Println(time.Now().String())
				// Deadlock
				// stopChan <- true
			}
		case <-stopChan:
			{
				fmt.Println("Stop!")
			}
		}
	}

}

/*
func main() {
	timeout := time.NewTicker(time.Duration(3) * time.Second)
	count := 0

	for {
		select {
		case <-timeout.C:
			{
				count++
				if count == 3 {
					continue
				}
				fmt.Println(count)
			}
		}
	}
}
*/

/*
func main() {
	timeout := time.NewTimer(time.Duration(3) * time.Second)
	count := 0

	for {
		select {
		case <-timeout.C:
			{
				count++
				fmt.Println(count)
				timeout.Reset(time.Duration(3) * time.Second)
				if count == 3 {
					timeout.Stop()
					return
				}
			}
		}
	}
}
*/

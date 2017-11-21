package main
import (
	"time"
	"fmt"
)

func main() {
	timeCh := time.NewTimer(5 * time.Second).C
	tickCh := time.NewTicker(2 * time.Second).C
	go func() {
		for {
			select {
				case _ = <- timeCh : {
					fmt.Printf("time: %s\n", time.Now().String())
				}
				case _ = <- tickCh : {
					fmt.Printf("tick: %s\n", time.Now().String())
				}	
			}	
		}
	}()
	time.Sleep(15 * time.Second)
}

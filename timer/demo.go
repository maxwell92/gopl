package main
import (
	"fmt"
	"time"
) 

func main() {
	select {
		case <-time.After(time.Second * time.Duration(1)):
			fmt.Println(time.Now().String())
	}	
}

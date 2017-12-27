package main
import(
	"sync"
	"time"
)

import "C"

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			C.sleep(1)
			wg.Done()
		}()
	}

	wg.Wait()
	println("done!")
	time.Sleep(time.Second * 5)
}

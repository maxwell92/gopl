package main
import(
	"fmt"
	"math"
	"runtime"
	"sync"
)

func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += i
	}
	println(x)
}

func test(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	numCPU := runtime.NumCPU()
	n := runtime.GOMAXPROCS(0)
	fmt.Printf("numCPU=%d, maxProcs=%d\n", numCPU, n)
	// test(n)
	test2(n)
}
	

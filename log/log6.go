package main

import (
	"fmt"
	"runtime"
	"sync"
)

type mylog struct {
	mu *sync.Mutex
}

func (l *mylog) Caller1(calldepth, flag, i int) {
	l.Caller2(calldepth, flag, i)
}

func (l *mylog) Caller2(calldepth, flag, i int) {
	l.Caller3(calldepth, flag, i)
}
func (l *mylog) Caller3(calldepth, flag, i int) {
	l.Caller4(calldepth, flag, i)
}
func (l *mylog) Caller4(calldepth, flag, i int) {
	l.Caller5(calldepth, flag, i)
}
func (l *mylog) Caller5(calldepth, flag, i int) {
	l.Caller6(calldepth, flag, i)
}
func (l *mylog) Caller6(calldepth, flag, i int) {
	l.PrintCaller(calldepth, flag, i)
}

func (l *mylog) PrintCaller(calldepth, flag, i int) {
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	if flag != 0 {
		l.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		l.mu.Lock()
	}
	fmt.Printf("[%d] %s:%d\n", i, file, line)
}

func (l *mylog) Output(calldepth int, flag int, i int, wg *sync.WaitGroup) {
	l.Caller1(calldepth, flag, i)
	wg.Done()
}

func main() {
	l := &mylog{
		mu: new(sync.Mutex),
	}

	wg := new(sync.WaitGroup)
	wg.Add(50000)
	for i := 0; i < 50000; i++ {
		go l.Output(1, 1, i, wg)
	}

	wg.Wait()
	//	time.Sleep(time.Duration(2) * time.Second)

}

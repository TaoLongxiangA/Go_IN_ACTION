package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32
	wg      sync.WaitGroup
)

func incCounter(id int) {
	//fmt.Println(counter)
	defer wg.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt32(&counter, 1)
	}
	runtime.Gosched()
}

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println(counter)
}

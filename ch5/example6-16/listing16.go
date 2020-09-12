package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mx      sync.Mutex
)

func incCounter(id int) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mx.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mx.Unlock()
	}
}

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println(counter)
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func incCounter(id int) {
	//fmt.Println(counter)
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println(counter)
}

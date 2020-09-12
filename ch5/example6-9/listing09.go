package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func incCounter(id int) {
	defer wg.Done()
	value := counter
	//runtime.Gosched()
	value++
	counter = value
}

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println(counter)
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	count int32
	wg    sync.WaitGroup
)

func doWork(id string) {
	defer wg.Done()
	for {
		fmt.Printf("%s is running\n", id)
		time.Sleep(10 * time.Millisecond)
		if atomic.LoadInt32(&count) == 1 {
			break
		}
	}
}

func main() {
	//runtime.GOMAXPROCS(1)
	wg.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)
	atomic.StoreInt32(&count, 1)
	wg.Wait()
}

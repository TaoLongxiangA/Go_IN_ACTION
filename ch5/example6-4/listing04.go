package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func PrintPrime(prefix string) {
	defer wg.Done()
next:
	for i := 2; i < 5000; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(1)
	go PrintPrime("A")
	go PrintPrime("B")
	wg.Wait()
}

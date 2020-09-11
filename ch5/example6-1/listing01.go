package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for j := 'a'; j < 'a'+26; j++ {
				fmt.Printf("%c ", j)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for j := 'A'; j < 'A'+26; j++ {
				fmt.Printf("%c ", j)
			}
		}
	}()

	wg.Wait()
}

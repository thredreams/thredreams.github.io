package main

import (
	"fmt"
	"sync"
	"time"
)

var j = 0

var k = 0

func main() {
	var wg sync.WaitGroup
	var rm sync.RWMutex
	wg.Add(2)
	go func() {
		time.Sleep(20 * time.Millisecond)
		for i := 0; i < 1000; i++ {
			rm.Lock()
			j++
			fmt.Printf("Write lock %d\n", j)
			rm.Unlock()
		}
		wg.Done()
	}()

	go func() {
		time.Sleep(18 * time.Millisecond)
		for i := 0; i < 1000; i++ {
			rm.RLock()
			k++
			fmt.Printf("Read lock %d\n", j)
			fmt.Printf("Read lock %d\n", k)
			rm.RUnlock()
		}
		wg.Done()
	}()

	wg.Wait()
}

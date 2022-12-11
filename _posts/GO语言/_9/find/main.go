package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var data []int = []int{1000000: 0}

func find0() {
	start := time.Now().UnixNano()
	for i := 0; i < len(data); i++ {
		if data[i] == 101 {
			fmt.Printf("find0 find the ans:%v, cost time: %v\n", i, time.Now().UnixNano()-start)
			return
		}
	}
}

func gofind(start, end int, starttime int64, wg *sync.WaitGroup) {
	for i := start; i < end; i++ {
		if data[i] == 101 {
			fmt.Printf("find1 find the ans:%v, cost time: %v\n", i, time.Now().UnixNano()-starttime)
			wg.Done()
			return
		}
	}
	wg.Done()
}

func find1() {
	var wg sync.WaitGroup
	starttime := time.Now().UnixNano()
	numOfCPU := runtime.NumCPU()
	batchSize := int(math.Floor(float64(1000000 / numOfCPU)))
	for i := 0; i < numOfCPU-1; i++ {
		wg.Add(1)
		go gofind(i*batchSize, (i+1)*batchSize, starttime, &wg)
	}
	go gofind(7*batchSize, len(data), starttime, &wg)
	wg.Wait()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)

	}
	ans := rand.Intn(1000000)
	data[ans] = 101
	find0()
	find1()
}

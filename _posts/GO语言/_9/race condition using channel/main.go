package main

import (
	"fmt"
	"sync"
)

var x = 0 //全局变量

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1 //访问全局变量
	<-ch
	wg.Done()
	fmt.Println(x)
}
func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x) //结果不确定
}

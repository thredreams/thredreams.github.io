package main

import (
	"fmt"
	"sync"
)

var x = 0 //全局变量

func increment(wg *sync.WaitGroup) {
	x = x + 1 //访问全局变量
	wg.Done()
	fmt.Println(x)
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
		if i == 999 {
			fmt.Println("i=999")
		}
	}
	w.Wait()
	fmt.Println("final value of x", x) //结果不确定
}

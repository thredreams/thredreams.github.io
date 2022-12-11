package main

import (
	"fmt"
	"sync"
)

var x = 0 //全局变量

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1 //锁定后访问全局变量
	m.Unlock()
	wg.Done()
	//fmt.Println(x)
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex //声明一个互斥锁
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x) //结果确定
}

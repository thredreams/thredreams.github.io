package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func main() {
	for i := 0; i < 5; i++ {
		go func(x int) {
			cond.L.Lock()         // wait 前，必须要先加锁
			defer cond.L.Unlock() //保障释放资源
			cond.Wait()
			fmt.Println(x)
			time.Sleep(time.Second * 1)
		}(i)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Signal....")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	fmt.Println("Signal....")
	cond.Signal() // 3 秒之后，下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	fmt.Println("Broadcast...")
	cond.Broadcast() // 3 秒之后，下发通知给所有已经获取锁的goroutine

	time.Sleep(time.Second * 3)
}

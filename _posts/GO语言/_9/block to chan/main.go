package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan bool) {
	fmt.Println("slowFunc began")
	time.Sleep(time.Second * 2)
	c <- true
	fmt.Println("sleeper() finished")
}

func main() {
	fmt.Println("main began")
	c := make(chan bool) //默认只能存储一条消息
	//c <- false
	start := time.Now().UnixNano() //记录当前时间的纳秒数
	//c <- true
	//slowFunc(c)
	go slowFunc(c)
	//time.Sleep(time.Second * 3) //避免main函数太快结束，导致 slowFunc()未运行完毕
	ready := <-c //阻塞，等待读取完成情况
	if ready {
		fmt.Printf("all time was: %v\n", (time.Now().UnixNano()-start)/1e9)
	}
}

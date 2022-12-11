package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}

func main() {
	start := time.Now().UnixNano() //记录当前时间的纳秒数
	//slowFunc()
	go slowFunc()
	//fmt.Println("I am not shown until slowFunc() completes")
	fmt.Println("I can show before  slowFunc() completes")
	time.Sleep(time.Second * 5) //避免main函数太快结束，导致 slowFunc()未运行完毕
	fmt.Printf("all time was: %v\n", (time.Now().UnixNano()-start)/1e9)
}

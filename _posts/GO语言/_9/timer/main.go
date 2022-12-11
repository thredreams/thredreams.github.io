package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("Timer 1 begin ", time.Now())

	//等待计时结束信号,信号就是当前时间戳
	//<-timer1.C
	//fmt.Println("Timer 1 end  ", time.Now())

	fmt.Println("Timer 1 end  ", <-timer1.C)
	timer2 := time.NewTimer(time.Second)
	fmt.Println("Timer 2 begin ", time.Now())
	go func() {

		//如强制超时，则timer2.C无法等到计时结束信号
		<-timer2.C

		fmt.Println("Timer 2 end  ", time.Now())
	}()

	//强制超时
	//stop2 := timer2.Stop()
	//if stop2 {
	//	fmt.Println("Timer 2 stopped", time.Now())
	//}

	time.Sleep(2 * time.Second)

}

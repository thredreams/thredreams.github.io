package main

import (
	"fmt"
	"time"
)

//ticker是一个定时触发的计时器，
//它会以一个间隔往channel发送信号，
//而channel的接收者可以以固定的时间间隔从channel中读取信号

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	i := 0
	for {

		//<-ticker.C
		fmt.Println("ticker time：", <-ticker.C)
		i++
		fmt.Println("i=", i)
		if i == 5 {
			ticker.Stop()
			break
		}
	}

}

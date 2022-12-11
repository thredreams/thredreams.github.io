package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(2 * time.Second) //可取消
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(2 * time.Second) //可取消
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	var reply string
	select {
	case reply = <-output1:
		fmt.Println(reply)
	case reply = <-output2:
		fmt.Println(reply)
	}
}

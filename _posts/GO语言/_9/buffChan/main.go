package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 5)
	go pump(ch1) // pump hangs
	go pull(ch1)
	time.Sleep(time.Second * 3)
}
func pump(ch chan int) {
	for i := 0; ; i++ {
		//for i := 0; i < 7; i++ {
		ch <- i
		fmt.Println("pump i", i)
	}
}
func pull(ch chan int) {
	for i := 0; ; i++ {
		//for i := 0; i < 7; i++ {
		fmt.Println(<-ch)
	}
}

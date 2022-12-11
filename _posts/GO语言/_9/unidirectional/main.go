package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
	//fmt.Println(<-sendch)
}

func main() {
	sendch := make(chan int)

	go sendData(sendch)
	//fmt.Println(sendch)
	fmt.Println(<-sendch)
}

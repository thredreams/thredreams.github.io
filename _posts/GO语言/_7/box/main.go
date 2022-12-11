package main

import "fmt"

func main() {
	var n int = 61
	var ei interface{} = n
	var m int = n
	fmt.Println("n  address:", &n)
	fmt.Println("m  address:", &m)
	fmt.Println("ei address:", &ei)

	var l int = 51
	ei = &l
	//m = &l
	p := ei.(*int)
	fmt.Println("l  address:", &l)
	fmt.Println("ei address:", &ei)
	fmt.Println("p  value  :", p)
	fmt.Println("p  address:", &p)
}

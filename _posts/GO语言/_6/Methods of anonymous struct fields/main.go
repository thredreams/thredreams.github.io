package main

import (
	"fmt"
)

type address struct {
	city  string
	state string
}

type person struct {
	firstName string
	lastName  string
	address
}

func (a *address) fullAddress() { //接收者是子结构指针
	fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}
func main() {
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}
	p.address.fullAddress() //完整调用方法
	p.fullAddress()         //父结构对象直接调用子结构方法

}

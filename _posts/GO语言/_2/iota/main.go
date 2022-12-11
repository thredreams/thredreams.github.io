package main

import "fmt"

//用iota来生成枚举值
const (
	a0 = iota
	a1 = iota
	a2 = iota
)

//用iota来快速生成枚举值
const (
	b0 = iota
	b3
	b1
	b2
)

//用iota来快速生成枚举值
const (
	c0 = iota
	c1
	c2
	c3 = 5
	c4 = iota - 1
	c5
)

// $GOROOT/src/time/time.go
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Printf("%v,%v,%v \n", a0, a1, a2)
	fmt.Printf("%v,%v,%v \n", b0, b1, b2)
	fmt.Printf("%v,%v,%v,%v,%v,%v  \n", c0, c1, c2, c3, c4, c5)
	fmt.Printf("%v,%v,%v \n", Sunday, Tuesday, Thursday)
}

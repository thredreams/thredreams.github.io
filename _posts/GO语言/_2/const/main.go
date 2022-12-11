package main

import "fmt"

const Pi float32 = 3.1415926
const e = 2.71828 //忽略类型变为无类型常量，更实用

//多个常量同时定义，常用于枚举
const (
	Red    = 0
	Yellow = 1
	Blue   = 2
)
const (
	Apple, Banana     = 11, 22
	Strawberry, Grape //隐式重复前一个非空表达式
	Pear, Watermelon
)

func main() {

	fmt.Printf("%v,%v \n%v,%v \n%v,%v  \n", Apple, Banana, Strawberry, Grape, Pear, Watermelon)
	var tmp float64 = 1
	fmt.Println(tmp + e)
	fmt.Println(tmp + 1.23456)
	fmt.Println(Pi + e)
	fmt.Print(tmp + Pi) //有类型常量不好用

}

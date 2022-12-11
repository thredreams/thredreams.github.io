package main

import "fmt"

var a int
var b, c int //多个同类变量同时声明
var h int    //全局变量声明了但不使用，OK

func main() {

	var b bool //遮盖父代码块变量

	var ( //多个不同类变量同时声明
		e bool
		f string
	)
	//var g int //函数内部变量声明了不使用，要报错

	fmt.Printf("a address: %v  value: %v \n", &a, a) //注意此处为printf函数
	fmt.Printf("b address: %v  value: %v \n", &b, b)
	fmt.Printf("c address: %v  value: %v \n", &c, c)
	fmt.Printf("e address: %v  value: %v \n", &e, e) //bool 零值为false
	fmt.Printf("f address: %v  value: %q \n", &f, f) //string 零值为""
}

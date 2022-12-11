package main

import (
	"fmt"
)

//var t = "Hello"

func appendStr() func(string) string {

	t := "Hello" //t被闭包引用，会导致相关内存不被释放
	a := " World"
	c := func(b string) string {
		t = t + a + " " + b
		return t
	}

	return c
}

func main() {
	a := appendStr()

	fmt.Println(a("China"))
	fmt.Println(a("Ningbo"))

	b := appendStr()

	fmt.Println(b("Everyone"))
	fmt.Println(b("!"))
}

package main

import (
	"fmt"
	"reflect"
)

type Money float64

func main() {
	var x Money = 58.9
	fmt.Println(9.0 / int(2))
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())               //查看底层类型
	fmt.Println("settability of v:", v.CanSet()) //能否被修改 x的地址不能修改
	p := reflect.ValueOf(&x)
	fmt.Println("kind:", p.Kind())                      //查看底层类型
	fmt.Println("settability of v:", p.CanSet())        //能否被修改，指向x的指针的地址不能修改
	fmt.Println("settability of v:", p.Elem().CanSet()) //能否被修改 *p 可以修改
}

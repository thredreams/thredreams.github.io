package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

//方法接收者是值
func (e Employee) changeName(newName string) {
	e.name = newName
}

//方法接收者是指针
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}
func main() {
	e := Employee{
		name: "Mark",
		age:  50,
	}

	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	//(&e).changeAge(51) //实际调用方法对象为指针
	e.changeAge(51) //GO 提供语法糖支持用变量来代替变量指针
	fmt.Printf("\nEmployee age after change: %d", e.age)
}

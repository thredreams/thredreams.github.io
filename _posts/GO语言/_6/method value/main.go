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
	fmt.Printf("\nEmployee name in func is: %s", e.name)
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
	//方法值赋给函数变量
	//f1 := e.changeName //对象方法值
	//f2 := Employee.changeName //类型方法值
	//f3 := e.changeAge //对象方法值
	//f4 := (*Employee).changeAge //类型方法值

	//fmt.Printf("Employee name before change: %s", e.name)
	//e.changeName("Michael") //方法调用
	//f1("Michael") //普通函数调用
	//f2(e, "Jack") //方法表达式调用
	//Employee.changeName(e, "Michael") //类型调用

	//fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	//e.changeAge(51) //方法调用
	//f3(51) //普通函数调用
	(*Employee).changeAge(&e, 52) //类型调用

	//f4(&e, 53)
	fmt.Printf("\nEmployee age after change: %d", e.age)
}

package main

import (
	"fmt"
)

type Printer interface { //接口类型命名通常以er为后缀
	Print()
}

type S1 struct{}

func (s S1) Print() { //实现Printer接口
	fmt.Println("S1 print")
}

type S2 struct{}

func (s *S2) Print() { //实现Printer接口，但接收者是指针
	fmt.Println("S2 print")
}

func main() {
	var i Printer
	fmt.Println(i)
	//i.Print() // error

	//必须初始化
	i = S1{}
	fmt.Printf("(%v, %T)\n", i, i) //查看底层类型的值和具体类型
	i.Print()

	//i = S2{} //报错，方法的接收者才能给接口变量赋值
	i = &S2{}
	fmt.Printf("(%v, %T)\n", i, i) //查看底层类型的值和具体类型
	i.Print()

}

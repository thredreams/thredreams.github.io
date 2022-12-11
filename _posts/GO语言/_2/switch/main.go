package main

import "fmt"

func Expr(n int) int {
	fmt.Println(n)
	return n
}

func main() {
	switch Expr(2) { //首先求值
	case Expr(1), Expr(2), Expr(3): //Expr(3)将被忽略。
		fmt.Println("enter into case1")
		fallthrough //fallthrough将执行权直接转移到下一个case执行语句
	case Expr(4): //fallthrough导致case表达式Expr(4)的求值略过。
		fmt.Println("enter into case2")
	}
}

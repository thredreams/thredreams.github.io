package main //定义当前代码所属的包，main是特殊包名，表示当前是一个可执行程序，不是库

import "fmt" //导入标准库的fmt（format）包

func main() { //程序执行入口
	for x := 100; x < 1000; x++ {
		i := x / 100
		j := (x % 100) / 10
		k := x % 10
		if i*100+j*10+k == i*i*i+j*j*j+k*k*k {
			fmt.Printf("%-4d", x)
		}
	}
	fmt.Print("hello, world") //默认不需要; 加;也没错
}

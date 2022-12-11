package main

import "runtime"

func main() {

	//获取当前 GOMAXPROCS
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	println("NumCPU=", runtime.NumCPU())
	//设置 GOMAXPROCS 的值为i
	runtime.GOMAXPROCS(2)
	//获取当前 GOMAXPROCS
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))

}

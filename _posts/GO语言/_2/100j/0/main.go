package main

import "fmt"

func main() {

	//如下代码存在什么问题
	for cock := 1; cock <= 100; cock++ {
		for hen := 1; hen <= 100; hen++ {
			for chick := 1; chick <= 100; chick++ {
				if 7*cock+5*hen+chick/3-100 != 0 {
					continue
				}
				if cock+hen+chick-100 != 0 {
					continue
				}
				if chick%3 != 0 {
					continue
				}
				fmt.Printf("Cock:%v\n", cock)
				fmt.Printf("Hens:%v\n", hen)
				fmt.Printf("Chicks:%v\n", chick)
			}
		}
	}

}

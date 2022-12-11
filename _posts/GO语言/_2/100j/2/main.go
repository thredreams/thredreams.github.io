package main

import "fmt"

func main() {

	//如下代码存在什么问题
	for cock := 1; cock <= 13; cock++ { //7*13=91，公鸡最多13只
		for hen := 1; hen <= 18; hen++ { //5*18=90，母鸡最多18只
			if 7*cock+5*hen+(100-cock-hen)/3-100 != 0 {
				continue
			}
			if (100-cock-hen)%3 != 0 {
				continue
			}
			fmt.Printf("Cock:%v\n", cock)
			fmt.Printf("Hens:%v\n", hen)
			fmt.Printf("Chicks:%v\n", (100 - cock - hen))
		}
	}

}

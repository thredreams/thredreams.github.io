package main

import "fmt"

func main() {

	for cock := 1; cock <= 13; cock++ { //7*13=91，公鸡最多13只
		for hen := 1; hen <= 18; hen++ { //5*18=90，母鸡最多18只
			chick := 100 - cock - hen
			if (chick%3 == 0) && (7*cock+5*hen+chick/3-100 == 0) {
				fmt.Printf("Cock:%v\n", cock)
				fmt.Printf("Hens:%v\n", hen)
				fmt.Printf("Chicks:%v\n", chick)
			}
		}
	}

}

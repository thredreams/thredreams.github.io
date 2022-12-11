package main

import (
	"fmt"
)

func main() {
	/*
		A去则B也去
		D、E两人中必有人去
		B、C两人必有人去，但只去一人
		C、D 两人要么都去，要么都不去
		若E去，则A也去
	*/
	for I := 0; I < 32; I++ {
		if ((I&16)>>4 == 1) && ((I&8)>>3 == 0) {
			continue
		}
		if ((I&2)>>1 == 0) && (I&1 == 0) {
			continue
		}
		if ((I & 8) >> 3) == ((I & 4) >> 2) {
			continue
		}
		if ((I & 4) >> 2) != ((I & 2) >> 1) {
			continue
		}

		if (I&1 == 1) && ((I&16)>>4 == 0) {
			continue
		}
		//输出结果
		fmt.Printf("A:%v\n", (I&16>>4 == 1))
		fmt.Printf("B:%v\n", (I&8>>3 == 1))
		fmt.Printf("C:%v\n", (I&4>>2 == 1))
		fmt.Printf("D:%v\n", (I&2>>1 == 1))
		fmt.Printf("E:%v\n\n", (I&1 == 1))
	}

}

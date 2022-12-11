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
		A, B, C, D, E := (I&16)>>4, (I&8)>>3, (I&4)>>2, (I&2)>>1, I&1
		if (A == 1) && (B == 0) {
			continue
		}
		if (D == 0) && (E == 0) {
			continue
		}
		if C != D {
			continue
		}
		if B == C {
			continue
		}
		if (E == 1) && (A == 0) {
			continue
		}
		//输出结果
		fmt.Printf("A:%v\n", (A == 1))
		fmt.Printf("B:%v\n", (B == 1))
		fmt.Printf("C:%v\n", (C == 1))
		fmt.Printf("D:%v\n", (D == 1))
		fmt.Printf("E:%v\n\n", (E == 1))
	}

}

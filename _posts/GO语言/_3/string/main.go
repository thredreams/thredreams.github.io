package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// original string
	var s string = "hello"
	fmt.Println("original string:", s)

	// 通过切片修改
	sl := []byte(s)
	sl[0] = 't'
	fmt.Println("slice:", string(sl))
	fmt.Println("after reslice, the original string is:", string(s))

	// 通过指针修改
	p := (*uintptr)(unsafe.Pointer(&s))
	var array *[5]byte = (*[5]byte)(unsafe.Pointer(*p))
	var len *int = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof((*uintptr)(nil))))
	for i := 0; i < (*len); i++ {
		fmt.Printf("%p => %c\n", &((*array)[i]), (*array)[i])
		p1 := &((*array)[i])
		v := (*p1)
		(*p1) = v + 1 //try to change the character
	}
}

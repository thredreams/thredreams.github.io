package main

import (
	"fmt"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from fullName", r)
	}
}
func recoverMain() {
	if r := recover(); r != nil {
		fmt.Println("recovered from main", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer recoverMain()
	defer fmt.Println("deferred call in main")
	//firstName := "Elon"
	lastName := "Potter"
	//fullName(&firstName, &lastName)
	fullName(nil, &lastName)
	fmt.Println("returned normally from main")
}
